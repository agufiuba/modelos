package problems

import (
	"math"
	"math/big"
	"strconv"
)

func (p *P) ComprarRecurso(recurso int, cantidad float64, costoTotal float64) (bool, float64, string) {
	usoActual := p.T.Zks[p.T.N+recurso-1]
	if usoActual != 0 {
		ratZOld := new(big.Rat).SetFloat64(p.T.Z).RatString()
		z := p.T.Z - costoTotal
		ratZ := new(big.Rat).SetFloat64(z).RatString()
		ratCosto := new(big.Rat).SetFloat64(costoTotal).RatString()
		// TOASK: Alcanza con "z = z - costo"?
		s := "No es conveniente, ya que el recurso se encuentra en exceso. Esto se puede ver en Z" + strconv.Itoa(p.T.N+recurso) + " != 0 (En la tabla del directo), o bien en B" + strconv.Itoa(recurso) + " != 0 (En la tabla del dual). Por lo tanto resulta que Z = " + ratZOld + " - " + ratCosto + " = " + ratZ + "."
		return false, z, s
	} else {
		// TODO: Respuesta de si conviene
		return true, 0, ""
	}
}

// TOTEST: Probar en minimizacion
func (p *P) VenderRecurso(recurso int, cantidad float64, precioTotal float64) (bool, float64, string) {
	var usoActual float64
	for i := 0; i < p.D.M; i++ {
		if p.D.Xks[i] == recurso {
			usoActual = p.D.Bks[recurso-1]
		}
	}
	if usoActual != 0 && usoActual-cantidad >= 0 {
		ratZOld := new(big.Rat).SetFloat64(p.T.Z).RatString()
		z := p.T.Z + precioTotal
		ratZ := new(big.Rat).SetFloat64(z).RatString()
		ratPrecio := new(big.Rat).SetFloat64(precioTotal).RatString()
		// TOASK: Alcanza con "z = z + precio"?
		s := "Si es conveniente, ya que el recurso se encuentra en exceso. Esto se puede ver en Z" + strconv.Itoa(p.T.N+recurso) + " != 0 (En la tabla del directo), o bien en B" + strconv.Itoa(recurso) + " != 0 (En la tabla del dual). Por lo tanto resulta que Z = " + ratZOld + " + " + ratPrecio + " = " + ratZ + "."
		return true, z, s
	} else {
		var s string
		if usoActual == 0 {
			s = "Como el recurso es limitante, esto se puede ver en B" + strconv.Itoa(recurso) + " = 0 (En la tabla del dual) o en Z" + strconv.Itoa(p.D.N+recurso-1) + " = 0 (En la tabla del directo), sera necesario aplicar el metodo simplex para obtener el nuevo valor de Z.\n"
		} else {
			ratCantidad := new(big.Rat).SetFloat64(cantidad).RatString()
			s = "Como el recurso se encuentra en exceso pero no lo suficiente como para vender sobras, esto se puede ver en B" + strconv.Itoa(recurso) + " < " + ratCantidad + " (En la tabla del dual) o en Z" + strconv.Itoa(p.D.N+recurso-1) + " < " + ratCantidad + " (En la tabla del directo), sera necesario aplicar el metodo simplex para obtener el nuevo valor de Z.\n"
		}
		s += "En este caso, se buscara la tabla optima para la nueva cantidad de recurso.\n"
		cpy := p.Copy()
		s += cpy.IterateDCambio(recurso-1, math.Abs(p.D.Cks[recurso-1])-cantidad)
		ratZ := new(big.Rat).SetFloat64(cpy.D.Z).RatString()
		z := cpy.D.Z + precioTotal
		s += "\n\nPor lo tanto, resulta que Z = " + ratZ + " + " + strconv.FormatFloat(precioTotal, 'g', -1, 64) + " = " + strconv.FormatFloat(cpy.D.Z+precioTotal, 'g', -1, 64) + "."
		return cpy.D.Z > p.D.Z, z, s
	}
}
