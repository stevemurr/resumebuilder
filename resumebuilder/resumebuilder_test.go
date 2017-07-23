package resumebuilder

import (
	"fmt"
	"os"
	"resumebuilder/model"
	"testing"
)

var (
	expected = `
\documentclass[11pt, a4paper, sans]{moderncv} % Font sizes: 10, 11, or 12; paper sizes: a4paper, letterpaper, a5paper, legalpaper, executivepaper or landscape; font families: sans or roman

\moderncvstyle{classic} % CV theme - options include: 'casual' (default), 'classic', 'oldstyle' and 'banking'
\moderncvcolor{bluej} % CV color - options include: 'blue' (default), 'orange', 'green', 'red', 'purple', 'grey' and 'black'

\usepackage{} % Used for inserting dummy 'Lorem ipsum' text into the template
\usepackage[scale=0.92]{geometry} % Reduce document margins
%\setlength{\hintscolumnwidth}{3cm} % Uncomment to change the width of the dates column
%\setlength{\makecvtitlenamewidth}{10cm} % For the 'classic' style, uncomment to adjust the width of the space allocated to your name

\addtolength{\oddsidemargin}{+.25in}
\addtolength{\textwidth}{-.5in}

\definecolor{light_gray}{HTML}{b9b9b9}
\definecolor{dark_gray}{HTML}{636363}
	
\firstname{\textcolor{dark_gray}{Steven}} % Your first name
\familyname{\textcolor{dark_gray}{Murr}} % Your last name
\title{Software}
\address{940 Scott Ct.}{Campbell, CA 95008}
\phone{408.4102513}
\email{stevemurr@gmail.com}
\homepage{https://github.com/stevemurr}

\begin{document}
\makecvtitle
\section{Skills}
\cvitem{Langauges}{\textbf{Proficient:} \hspace{17pt}  Go, Python \newline
\textbf{Familiar with:} \hspace{0.4pt}  Python, Go }
\cvitem{Frameworks}{ React, Go}
\hspace{-0.14cm}\cvitem{Tools/Libraries}{\hspace{0cm} Scikit Learn, React}
\cvitem{Devices}{ Raspberry PI, ESP8266}
\section{Experience}
\cventry{Jan 2017 \\-- Feb 2017}{Software}{Microsoft}{}{}{
	\textcolor{dark_gray}{\textit{ Go, Python}}
	\begin{itemize}
		\item this is point 1\item this is point 2
	\end{itemize}
}

\cventry{Jan 2017 \\-- Feb 2017}{Tester}{Namco Networks}{}{}{
	\textcolor{dark_gray}{\textit{ Go, Python}}
	\begin{itemize}
		\item this is point 3\item this is point 4
	\end{itemize}
}
\section{Education}
\cventry{Jan 2018 \\-- Feb 2019}{B.S. in Computer Science}{West Valley}{}{}{
\textbf{60\% Complete}
}\end{document}`
)

func TestMake(t *testing.T) {
	f, err := os.Open("/Users/murr/go/src/resumebuilder/test/config.toml")
	if err != nil {
		t.Error("cant open the test file")
	}
	var config model.Config
	if err := ReadTOML(f, &config); err != nil {
		t.Error("could not read the toml filej")
	}
	result := Make(config)
	fmt.Println(result)
}
