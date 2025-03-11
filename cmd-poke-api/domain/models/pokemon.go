package model

type Type = string

type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Types  []Type `json:"types"`
}

func (p *Pokemon) String() string {
	// json string
	types := ""
	for i, t := range p.Types {
		if i == 0 {
			types += `"` + t + `"`
		} else {
			types += `,"` + t + `"`
		}
	}

	return `{"name":"` + p.Name + `","height":` + string(p.Height) + `,"weight":` + string(p.Weight) + `,"types":` + types + `}`
}
