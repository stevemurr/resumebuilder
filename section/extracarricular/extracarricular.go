package extracarricular

// Extracarricular --
type Extracarricular struct {
	From     string   `json:"from"`
	To       string   `json:"to"`
	Title    string   `json:"title"`
	Company  string   `json:"company"`
	Subtitle []string `json:"subtitle"`
	Points   []string `json:"points"`
}

// Name returns the section header
func (e *Extracarricular) Name() string {
	return "Experience"
}

// GetTemplate returns the template
func (e *Extracarricular) GetTemplate() string {
	return `
\cventry{<.From> \\-- <.To>}{<.Title>}{<.Company>}{}{}{
	\textcolor{dark_gray}{\textit{<range $index, $element := .Subtitle><if $index>,<end> <.><end>}}
	\begin{itemize}
		<range .Points>\item <.><end>
	\end{itemize}
}
`
}
