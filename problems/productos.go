package problems

import (
	"math"
	"math/big"
	"modelos/table"
	"strconv"
)

func (p *P) ComprarProductoReducirRestriccion(restriccion int, newR float64, costo float64) (bool, float64, string) {
	cpy := p.Copy()
	oldR := cpy.D.Cks[restriccion-1]
	s := "Para ver si comprar unidades de producto ya fabricado es conveniente, sera necesario analizar como se comporta la tabla del dual al modificar la restriccion asociada.\n"
	s += cpy.IterateDCambio(restriccion-1, oldR/math.Abs(oldR)*newR)
	ratZ := new(big.Rat).SetFloat64(cpy.D.Z).RatString()
	z := cpy.D.Z - costo
	s += "\n\nPor lo tanto, resulta que Z = " + ratZ + " - " + strconv.FormatFloat(costo, 'g', -1, 64) + " = " + strconv.FormatFloat(z, 'g', -1, 64) + "."
	return cpy.D.Z > p.D.Z, z, s
}

func (p *P) AgregarProducto(indices []float64, ck float64) (bool, float64, string) {
	inv := p.GetInversa()
	s := "Para agregar un producto, es necesario buscar la matriz inversa del directo, la cual vale:\n\n"
	s += table.ColsString(inv)
	c := table.NewCol(len(inv))
	for i := 0; i < len(inv[0].Values); i++ {
		for j := 0; j < len(inv); j++ {
			c.Values[i] += inv[j].Values[i] * indices[j]
		}
	}
	sindex := "["
	for _, i := range indices {
		sindex += new(big.Rat).SetFloat64(i).RatString() + ", "
	}
	sindex = sindex[0:len(sindex)-2] + "]"
	vaux := "["
	for _, i := range c.Values {
		vaux += new(big.Rat).SetFloat64(i).RatString() + ", "
	}
	vaux = vaux[0:len(vaux)-2] + "]"
	s += "\nLuego, debemos multiplicarla por los indices correspondientes al nuevo producto " + sindex + ", de donde se obtiene: " + vaux
	s += "\nEste vector sera agregado en forma de columna a la tabla optima del directo, quedando entonces:\n\n"

	cpy := p.Copy()
	cpy.T.Cols = append(cpy.T.Cols[0:cpy.T.N], c)
	for _, c := range p.T.Cols[cpy.T.N:] {
		cpy.T.Cols = append(cpy.T.Cols, c)
	}
	cpy.T.Cks = append(cpy.T.Cks[0:cpy.T.N], ck)
	for _, c := range p.T.Cks[cpy.T.N:] {
		cpy.T.Cks = append(cpy.T.Cks, c)
	}
	var newz float64
	for i, x := range cpy.T.Xks {
		newz += cpy.T.Cks[x-1] * c.Values[i]
	}
	newz -= ck
	cpy.T.Zks = append(cpy.T.Zks[0:cpy.T.N], newz)
	for _, z := range p.T.Zks[cpy.T.N:] {
		cpy.T.Zks = append(cpy.T.Zks, z)
	}
	for i, xk := range cpy.T.Xks {
		if xk > cpy.T.N {
			cpy.T.Xks[i] = xk + 1
		}
	}
	cpy.T.N++
	s += cpy.T.S()

	var conv bool
	o, _ := cpy.Optimo()
	if !o {
		s += "\n\nComo la nueva tabla no es optima, hay que iterar:"
		s += cpy.Iterate()
		conv = cpy.T.Z > p.T.Z
		if conv {
			s += "\n\nComo Z = " + strconv.FormatFloat(cpy.T.Z, 'g', -1, 64) + " > " + strconv.FormatFloat(p.T.Z, 'g', -1, 64) + ", es conveniente fabricar este nuevo producto."
		} else {
			s += "\n\nComo Z = " + strconv.FormatFloat(cpy.T.Z, 'g', -1, 64) + " < " + strconv.FormatFloat(p.T.Z, 'g', -1, 64) + ", no es conveniente fabricar este nuevo producto."
		}
	} else {
		s += "\n\nComo la tabla continua siendo optima, significa que la columna del nuevo producto no entrara a la base, lo que equivale a decir que no conviene fabricarlo."
	}
	return conv, cpy.T.Z, s
}
