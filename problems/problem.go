package problems

import (
	"encoding/json"
	"modelos/table"
)

const (
	Min = iota
	Max
)

type P struct {
	T      *table.T
	D      *table.T
	Minmax uint
}

func New(t *table.T, d *table.T, minmax uint) *P {
	p := P{}
	p.T = t
	p.D = d
	p.Minmax = minmax
	return &p
}

func (p *P) Copy() *P {
	j, _ := json.Marshal(p)
	cpy := new(P)
	json.Unmarshal(j, &cpy)
	return cpy
}
