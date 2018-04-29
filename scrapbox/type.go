package main

type Project struct {
	Name        string
	DisplayName string
	Exported    uint
	Pages       []PagesType
}

type PagesType struct {
	Title   string
	Created uint
	Updated uint
	Lines   []LinesType
}

type LinesType struct {
	Text    string
	Created uint
	Updated uint
}
