package main

import (
	"fmt"
	"math/big"
)

func pascal(n int) []int {
	if n < 0 {
		panic("n debe ser no negativo")
	}
	if n == 0 {
		return []int{1}
	}
	if n == 1 {
		return []int{1, 1}
	}

	prevRow := pascal(n - 1)
	row := make([]int, n+1)
	row[0] = 1
	row[n] = 1

	for i := 1; i < n; i++ {
		row[i] = prevRow[i-1] + prevRow[i]
	}
	return row
}

func generar_polinomio(n int) string {
	coefficients := pascal(n)
	polynomial := ""
	for i := n; i >= 0; i-- {
		coeff := coefficients[i]
		if coeff == 0 {
			continue
		}
		if coeff == 1 && i != 0 {
			if i == n {
				polynomial += "x^" + fmt.Sprint(i)
			} else {
				polynomial += "+x^" + fmt.Sprint(i)
			}
		} else if coeff == -1 {
			if i == n {
				polynomial += "-x^" + fmt.Sprint(i)
			} else {
				polynomial += "-x^" + fmt.Sprint(i)
			}
		} else {
			if i == n {
				polynomial += fmt.Sprint(coeff) + "x^" + fmt.Sprint(i)
			} else if i == 0 {
				polynomial += fmt.Sprint(coeff)
			} else {
				polynomial += fmt.Sprint(coeff) + "x^" + fmt.Sprint(i)
			}
		}
		if i > 0 {
			polynomial += " + "
		}
	}
	return polynomial
}

func calcular(n int, x int64) int64 {
	coefficients := pascal(n)
	result := int64(0)
	for i := n; i >= 0; i-- {
		term := int64(1)
		for j := 0; j < i; j++ {
			term *= x
		}
		result += term * int64(coefficients[i])
	}
	return result
}

func main() {
	var n int
	fmt.Print("Ingrese el exponente n (entero no negativo): ")
	fmt.Scanln(&n)

	polynomial := generar_polinomio(n)
	fmt.Printf("El polinomio (x+1)^%d es: %s\n", n, polynomial)

	var x int64
	fmt.Print("Ingrese el valor de x: ")
	fmt.Scanln(&x)

	result := calcular(n, x)
	fmt.Printf("El valor del polinomio para x = %d es: %v\n", x, result)

	fmt.Println("\nCálculo paso a paso para x =", x, "y n =", n, ":")
	coefficients := pascal(n)
	var total float64 = 0
	for i := n; i >= 0; i-- {
		term := new(big.Int)
		term.Exp(big.NewInt(x), big.NewInt(int64(i)), nil)
		bigCoeff := big.NewInt(int64(coefficients[i]))
		term.Mul(term, bigCoeff)
		fmt.Printf("Término (%v)*%d^%d = %v\n", coefficients[i], x, i, term)
		total += float64(term.Int64())
	}
	fmt.Printf("(%d + 1)^%d = %.2f", x, n, total)
}
