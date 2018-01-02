package table

import (
	"math/big"
	"strconv"
)

func (t *T) S() string {
	cols := make([][]string, t.M+t.N+3)
	cols[0] = make([]string, t.M+1)
	for i := 1; i < 3; i++ {
		cols[i] = make([]string, t.M+2)
	}
	for i := 3; i < t.N+t.M+3; i++ {
		cols[i] = make([]string, t.M+3)
	}
	cols[0][0] = "Ck"
	if !t.Dual {
		cols[1][0] = "Xk"
	} else {
		cols[1][0] = "Yk"
	}
	cols[2][0] = "Bk"
	for i := 1; i <= t.N+t.M; i++ {
		cols[i+2][1] = "A" + strconv.Itoa(i)
	}
	for i := 0; i < t.M; i++ {
		if !t.Dual {
			cols[1][i+1] = "X" + strconv.Itoa(t.Xks[i])
		} else {
			cols[1][i+1] = "Y" + strconv.Itoa(t.Xks[i])
		}
	}
	for i := 0; i < t.M; i++ {
		j, _ := strconv.Atoi(cols[1][i+1][len(cols[1][i+1])-1:])
		ck := new(big.Rat).SetFloat64(t.Cks[j-1]).RatString()
		cols[0][i+1] = ck
	}
	for i := 0; i < t.M; i++ {
		cols[2][i+1] = new(big.Rat).SetFloat64(t.Bks[i]).RatString()
	}
	for i := 0; i < t.N+t.M; i++ {
		for j := 0; j < t.M; j++ {
			cols[i+3][j+2] = new(big.Rat).SetFloat64(t.Cols[i].Values[j]).RatString()
		}
		if t.Cks[i] != 0 {
			cols[i+3][0] = new(big.Rat).SetFloat64(t.Cks[i]).RatString()
		}
	}
	cols[1][t.M+1] = "Z"
	cols[2][t.M+1] = new(big.Rat).SetFloat64(t.Z).RatString()
	for i := 0; i < t.N+t.M; i++ {
		cols[i+3][t.M+2] = new(big.Rat).SetFloat64(t.Zks[i]).RatString()
	}
	maxtotal := 9
	init := 3
	for i := 0; i < t.N+t.M+3; i++ {
		max := 0
		for j := 0; j < len(cols[i]); j++ {
			if len(cols[i][j]) > max {
				max = len(cols[i][j])
			}
		}
		if i == 0 {
			init += max
			maxtotal += max
			for j := 0; j < t.M+1; j++ {
				space := max - len(cols[i][j])
				for k := 0; k < space+3; k++ {
					cols[i][j] = cols[i][j] + " "
				}
			}
		} else {
			if i < 3 {
				maxtotal += max
				for j := 0; j <= t.M+1; j++ {
					space := max - len(cols[i][j])
					for k := 0; k < space+3; k++ {
						cols[i][j] = cols[i][j] + " "
					}
				}
			} else {
				for j := 0; j <= t.M+2; j++ {
					space := max - len(cols[i][j])
					for k := 0; k < space+3; k++ {
						cols[i][j] = cols[i][j] + " "
					}
				}
			}
		}
	}
	s := ""
	for i := 0; i < maxtotal; i++ {
		s += " "
	}
	for i := 0; i < t.N+t.M; i++ {
		s += cols[i+3][0]
	}
	s += "\n"
	for i := 0; i <= t.M; i++ {
		for j := 0; j < t.N+t.M+3; j++ {
			if j > 2 {
				s += cols[j][1+i]
			} else {
				s += cols[j][0+i]
			}
		}
		s += "\n"
	}
	for i := 0; i < init; i++ {
		s += " "
	}
	for i := 1; i <= t.N+t.M+2; i++ {
		if i < 3 {
			s += cols[i][t.M+1]
		} else {
			s += cols[i][t.M+2]
		}
	}
	return s
}

func ColsString(cols []C) string {
	cs := make([][]string, len(cols))
	for i, c := range cols {
		cs[i] = make([]string, len(c.Values))
		max := 0
		for j, v := range c.Values {
			cs[i][j] = new(big.Rat).SetFloat64(v).RatString()
			if len(cs[i][j]) > max {
				max = len(cs[i][j])
			}
		}
		for j, v := range cs[i] {
			for k := 0; k < max-len(v); k++ {
				cs[i][j] += " "
			}
		}
	}
	s := ""
	for i := 0; i < len(cs[0]); i++ {
		for j := 0; j < len(cs); j++ {
			s += cs[j][i] + " "
		}
		s += "\n"
	}
	return s
}
