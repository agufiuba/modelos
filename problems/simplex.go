package problems

import (
	"modelos/table"
)

func (p *P) Iterate() string {
	s := ""
	if p.Minmax == Max {
		positives, j := AllPositives(p.T.Zks, p.T.N+p.T.M)
		for !positives {
			s += "\n\n" + i(p.T, j)
			positives, j = AllPositives(p.T.Zks, p.T.N+p.T.M)
		}
	}
	return s
}

func (p *P) IterateD() string {
	s := ""
	if p.Minmax == Max {
		negatives, j := AllNegatives(p.D.Zks, p.D.N+p.D.M)
		for !negatives {
			s += "\n\n" + i(p.D, j)
			negatives, j = AllNegatives(p.D.Zks, p.D.N+p.D.M)
		}
	}
	return s
}

func (p *P) IterateDCambio(col int, value float64) string {
	p.D.Cks[col] = value
	var zcol float64
	for i := 0; i < p.D.M; i++ {
		zcol += p.D.Cols[col].Values[i] * p.D.Cks[p.D.Xks[i]-1]
	}
	zcol -= value
	updateZs(p.D)
	return "\n" + p.D.S() + p.IterateD()
}

func i(t *table.T, cpivot int) string {
	var min float64
	min = 1000
	fpivot := 0
	for i := 0; i < t.M; i++ {
		if t.Cols[cpivot].Values[i] != 0 {
			div := t.Bks[i] * t.Cols[cpivot].Values[i]
			if div > 0 && div < min {
				min = div
				fpivot = i
			}
		}
	}
	celpivot := t.Cols[cpivot].Values[fpivot]
	multi := 1 / celpivot
	for i := 0; i < t.N+t.M; i++ {
		t.Cols[i].Values[fpivot] *= multi
	}
	t.Bks[fpivot] *= multi
	for i := 0; i < t.M; i++ {
		if i != fpivot {
			k := t.Cols[cpivot].Values[i]
			for j := 0; j < t.N+t.M; j++ {
				t.Cols[j].Values[i] -= k * t.Cols[j].Values[fpivot]
			}
			t.Bks[i] -= k * t.Bks[fpivot]
		}
	}
	t.Xks[fpivot] = cpivot + 1
	updateZs(t)
	return t.S()
}

func AllNegatives(xs []float64, s int) (bool, int) {
	i := 0
	for i < s {
		if xs[i] > 0 {
			return false, i
		} else {
			i++
		}
	}
	return true, 0
}

func AllPositives(xs []float64, s int) (bool, int) {
	i := 0
	for i < s {
		if xs[i] < 0 {
			return false, i
		} else {
			i++
		}
	}
	return true, 0
}

func updateZs(t *table.T) {
	t.Z = 0
	for i := 0; i < t.M; i++ {
		t.Z += t.Cks[t.Xks[i]-1] * t.Bks[i]
	}
	for i := 0; i < t.N+t.M; i++ {
		t.Zks[i] = 0
		for j := 0; j < t.M; j++ {
			t.Zks[i] += t.Cols[i].Values[j] * t.Cks[t.Xks[j]-1]
		}
		t.Zks[i] -= t.Cks[i]
	}
}

func (p *P) GetInversa() []table.C {
	cols := make([]table.C, p.T.M)
	for i := p.T.N; i < p.T.N+p.T.M; i++ {
		c := table.NewCol(p.T.M)
		for j := 0; j < p.T.M; j++ {
			if p.T.May[i-p.T.N] {
				c.Values[j] = -p.T.Cols[i].Values[j]
			} else {
				c.Values[j] = p.T.Cols[i].Values[j]
			}
		}
		cols[i-p.T.N] = c
	}
	return cols
}

func (p *P) GetInversaD() []table.C {
	cols := make([]table.C, p.D.M)
	for i := p.D.N; i < p.D.N+p.D.M; i++ {
		c := table.NewCol(p.D.M)
		for j := 0; j < p.D.M; j++ {
			if p.T.Xks[i-p.D.N] > 0 {
				c.Values[j] = -p.D.Cols[i].Values[j]
			} else {
				c.Values[j] = p.D.Cols[i].Values[j]
			}
		}
		cols[i-p.D.N] = c
	}
	return cols
}

func (p *P) Optimo() (bool, string) {
	if p.Minmax == Max {
		o, _ := AllPositives(p.T.Zks, p.T.N+p.T.M)
		return o, "positivos"
	} else {
		o, _ := AllNegatives(p.T.Zks, p.T.N+p.T.M)
		return o, "negativos"
	}
}

func (p *P) OptimoD() (bool, string) {
	if p.Minmax == Max {
		o, _ := AllPositives(p.D.Zks, p.D.N+p.D.M)
		return o, "positivos"
	} else {
		o, _ := AllNegatives(p.D.Zks, p.D.N+p.D.M)
		return o, "negativos"
	}
}
