package models

type Page struct{
	Title string
	Warnings []string
	Alerts []string
	Messages []string
	Model interface{}
}

func (p *Page) AddWarning(m string){
	p.Warnings = append(p.Warnings, m)
}
func (p *Page) AddAlert(m string){
	p.Alerts = append(p.Alerts, m)
}
func (p *Page) AddMessage(m string){
	p.Messages = append(p.Messages, m)
}
