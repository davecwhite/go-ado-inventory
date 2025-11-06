package main

type UserResults struct {
	Count int     `json:"count"`
	Value []users `json:"value"`
}

type users struct {
	Descriptor    string
	DisplayName   string
	PrincipalName string
	MailAddress   string
	SubjectKind   string
	Domain        string
}

type ProjectResults struct {
	Count int       `json:"count"`
	Value []project `json:"value"`
}

type project struct {
	Id          string
	Name        string
	Description string
	State       string
	Visibility  string
	LastUpdate  string
	URL         string
}
