package resumebuilder

import (
	"bytes"
	"html/template"
	"io"
	"resumebuilder/model"
	"resumebuilder/section"

	"github.com/naoina/toml"
)

var (
	leftDelim  = "<"
	rightDelim = ">"
	footer     = `\end{document}`
	begin      = `
			\begin{document}
			\makecvtitle
	`
)

// Parse returns the template as a string
func Parse(s section.Section) string {
	var buf bytes.Buffer
	tmpl := template.Must(template.New("").Delims(leftDelim, rightDelim).Parse(s.GetTemplate()))
	if err := tmpl.Execute(&buf, s); err != nil {
		return ""
	}
	return buf.String()
}

// ReadTOML --
func ReadTOML(f io.Reader, val interface{}) error {
	if err := toml.NewDecoder(f).Decode(val); err != nil {
		return err
	}
	return nil
}

// Make --
func Make(config model.Config) string {
	result := Parse(&config.Settings)
	result += Parse(&config.Personal)
	result += begin
	result += `\section{Skills}`
	result += Parse(config.Skills)
	result += `\section{Experience}`
	for _, value := range config.Experience {
		result += Parse(&value)
	}
	result += `\section{Education}`
	for _, value := range config.Education {
		result += Parse(&value)
	}

	result += `\section{Extracarriculr}`
	for _, value := range config.Extracarricular {
		result += Parse(&value)
	}
	result += footer
	return result
}
