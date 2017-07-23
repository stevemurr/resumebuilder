package settings

// Settings --
type Settings struct {
	Scaling string
	CVStyle string
	CVColor string
}

// Name --
func (s *Settings) Name() string {
	return "Settings"
}

// GetTemplate --
func (s *Settings) GetTemplate() string {
	return `
\documentclass[11pt, a4paper, sans]{moderncv} % Font sizes: 10, 11, or 12; paper sizes: a4paper, letterpaper, a5paper, legalpaper, executivepaper or landscape; font families: sans or roman

\moderncvstyle{<.CVStyle>} % CV theme - options include: 'casual' (default), 'classic', 'oldstyle' and 'banking'
\moderncvcolor{<.CVColor>} % CV color - options include: 'blue' (default), 'orange', 'green', 'red', 'purple', 'grey' and 'black'

\usepackage{} % Used for inserting dummy 'Lorem ipsum' text into the template
\usepackage[scale=<.Scaling>]{geometry} % Reduce document margins
%\setlength{\hintscolumnwidth}{3cm} % Uncomment to change the width of the dates column
%\setlength{\makecvtitlenamewidth}{10cm} % For the 'classic' style, uncomment to adjust the width of the space allocated to your name

\addtolength{\oddsidemargin}{+.25in}
\addtolength{\textwidth}{-.5in}

\definecolor{light_gray}{HTML}{b9b9b9}
\definecolor{dark_gray}{HTML}{636363}
	`
}
