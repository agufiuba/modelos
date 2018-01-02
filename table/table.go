package table

type T struct {
	N    int
	M    int
	Cks  []float64
	Xks  []int
	Bks  []float64
	Z    float64
	Zks  []float64
	Cols []C
	Dual bool
	May  []bool
}

type C struct {
	Values []float64
}

func New(n int, m int, d bool) *T {
	t := T{}
	t.N = n
	t.M = m
	t.Cks = make([]float64, n+m)
	t.Xks = make([]int, m)
	t.Bks = make([]float64, m)
	t.Cols = make([]C, m+n)
	t.Zks = make([]float64, m+n)
	t.Dual = !d
	if t.Dual {
		t.May = make([]bool, n)
	} else {
		t.May = make([]bool, m)
	}
	for i := range t.Cols {
		t.Cols[i] = NewCol(m)
	}
	return &t
}

func NewCol(m int) C {
	c := C{}
	c.Values = make([]float64, m)
	return c
}

func (t *T) SetCols(vs ...float64) {
	i := 0
	j := 0
	c := NewCol(t.M)
	for _, v := range vs {
		if i < t.M {
			c.Values[i] = v
			i++
		} else {
			t.Cols[j] = c
			c = NewCol(t.M)
			i = 1
			c.Values[0] = v
			j++
		}
	}
	t.Cols[j] = c
}
