package file

import "io/ioutil"

type Page struct { // make sure to include .html
	Title string
	Body  []byte
}

func (p *Page) Save() error {
	filename := p.Title
	return ioutil.WriteFile(filename, p.Body, 0600)
}
func Load(title string) (*Page, error) {
	body, err := ioutil.ReadFile(title)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
