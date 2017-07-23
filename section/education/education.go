package education

// Education --
type Education struct {
	To          string
	From        string
	Institution string
	Degree      string
	Note        string
}

// Name --
func (e *Education) Name() string {
	return "Education"
}

// GetTemplate --
func (e *Education) GetTemplate() string {
	return `
\cventry{<.To> \\ <.From>}{<.Degree>}{<.Institution>}{}{}{
<.Note>
}`
}

// \textbf{<.BoldNote>}
