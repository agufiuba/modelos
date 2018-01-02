package main

import (
	"fmt"
	"modelos/problems"
	"modelos/table"
)

func main() {
	// fmt.Println("Modelos")

	t := table.New(2, 3, true)
	t.Cks[0] = 60
	t.Cks[1] = 40
	t.Xks = []int{1, 4, 2}
	t.Bks = []float64{30, 0, 10}
	t.Z = 2200
	t.Zks = []float64{0, 0, 30, 0, 20}
	t.May[2] = true
	t.SetCols(1, 0, 0, 0, 0, 1, 0.5, -0.5, 0, 0, 1, 0, 1, 1, -1)
	fmt.Println(t.S() + "\n")

	td := table.New(3, 2, false)
	td.Cks[0] = 80
	td.Cks[1] = 50
	td.Cks[2] = -10
	td.Xks = []int{1, 3}
	td.Bks = []float64{30, 20}
	td.Z = 2200
	td.Zks = []float64{0, 0, 0, -30, -10}
	td.May[2] = true
	td.SetCols(1, 0, 0.5, -1, 0, 1, -0.5, -1, 0, 1)
	td.Dual = true
	fmt.Println(td.S() + "\n")

	p := problems.New(t, td, problems.Max)
	_, _, c := p.ComprarRecurso(1, 20, 200)
	fmt.Println(c)
	_, _, c2 := p.VenderRecurso(2, 10, 700)
	fmt.Println(c2)
	_, _, c3 := p.ComprarProductoReducirRestriccion(3, 5, 225)
	fmt.Println(c3)
	_, _, c4 := p.AgregarProducto([]float64{3, 1, 0}, 55)
	fmt.Println(c4)
}
