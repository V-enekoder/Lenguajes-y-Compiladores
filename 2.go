package main

import (
	"fmt"
	"math/big"
)

func pascal(n int) []big.Int {
	if n < 0 {
		panic("n debe ser no negativo")
	}
	if n == 0 {
		return []big.Int{*big.NewInt(1)}
	}
	if n == 1 {
		return []big.Int{*big.NewInt(1), *big.NewInt(1)}
	}

	prevRow := pascal(n - 1)
	row := make([]big.Int, n+1)
	row[0] = *big.NewInt(1)
	row[n] = *big.NewInt(1)

	for i := 1; i < n; i++ {
		row[i].Add(&prevRow[i-1], &prevRow[i])
	}
	return row
}

func generar_polinomio(n int) string {
	coefficients := pascal(n)
	polynomial := ""
	for i := n; i >= 0; i-- {
		coeff := coefficients[i]
		if coeff.Cmp(big.NewInt(0)) == 0 {
			continue // Saltar coeficientes cero
		}
		if coeff.Cmp(big.NewInt(1)) == 0 && i != 0 {
			if i == n {
				polynomial += "x^" + fmt.Sprint(i)
			} else {
				polynomial += "+x^" + fmt.Sprint(i)

			}

		} else if coeff.Cmp(big.NewInt(-1)) == 0 {
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

func calcular(n int, x int64) *big.Int {
	coefficients := pascal(n)
	result := big.NewInt(0)
	for i := n; i >= 0; i-- {
		term := new(big.Int)
		term.Exp(big.NewInt(x), big.NewInt(int64(i)), nil)
		term.Mul(term, &coefficients[i])
		result.Add(result, term)
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
	for i := n; i >= 0; i-- {
		term := new(big.Int)
		term.Exp(big.NewInt(x), big.NewInt(int64(i)), nil)
		term.Mul(term, &coefficients[i])
		fmt.Printf("Término (%v)x^%d = %v\n", coefficients[i], i, term) // Imprime correctamente big.Int
	}
}
