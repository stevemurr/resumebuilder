package personal

// Personal --
type Personal struct {
	Title        string
	FirstName    string
	LastName     string
	PhoneNumber  string
	Github       string
	Email        string
	Address      string
	CityAndState string
}

// Name --
func (p *Personal) Name() string {
	return "Personal"
}

// GetTemplate --
func (p *Personal) GetTemplate() string {
	return `
\firstname{\textcolor{dark_gray}{<.FirstName>}} % Your first name
\familyname{\textcolor{dark_gray}{<.LastName>}} % Your last name
\title{<.Title>}
\address{<.Address>}{<.CityAndState>}
\phone{<.PhoneNumber>}
\email{<.Email>}
\homepage{https://github.com/<.Github>}
`
}
