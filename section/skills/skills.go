package skills

/*
\section{Skills}
\cvitem{Langauges}{\textbf{Proficient:} \hspace{17pt} Python, Go, NodeJS, Bash \newline
				\textbf{Familiar with:} \hspace{0.4pt} Swift, Java, C }
\cvitem{Frameworks}{React, Angular, SciKit Learn}
\hspace{-0.14cm}\cvitem{Tools/Libraries}{\hspace{0.14cm}Git, Travis CI, GNU Make, Bootstrap, jQuery}
\cvitem{Devices}{Raspberry PI, Banana PI, ESP8266, Arduino}
*/

// Skills --
type Skills struct {
	Proficient []string
	Familiar   []string
	Frameworks []string
	Libraries  []string
	Devices    []string
}

// GetTemplate returns the template as string
// This can come from an external file or hardcorded string
// as seen here.
func (s Skills) GetTemplate() string {
	return `
\cvitem{Langauges}{\textbf{Proficient:} \hspace{17pt} <range $index, $element := .Proficient><if $index>,<end> <.><end> \newline
\textbf{Familiar with:} \hspace{0.4pt} <range $index, $element := .Familiar><if $index>,<end> <.><end> }
\cvitem{Frameworks}{<range $index, $element := .Frameworks><if $index>,<end> <.><end>}
\hspace{-0.14cm}\cvitem{Tools/Libraries}{\hspace{0cm}<range $index, $element := .Libraries><if $index>,<end> <.><end>}
\cvitem{Devices}{<range $index, $element := .Devices><if $index>,<end> <.><end>}
`
}

// Name returns the name of the section
func (s Skills) Name() string {
	return "Skills"
}

// Props --
func (s Skills) Props() interface{} {
	return s
}
