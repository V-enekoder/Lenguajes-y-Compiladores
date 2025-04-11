package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"

	"github.com/Knetic/govaluate"
)

func main() {
	expression := "(2E10 – 1e15)/5E-2"

	// Registrar una función personalizada para manejar la notación científica
	functions := map[string]govaluate.ExpressionFunction{
		"sci": func(args ...interface{}) (interface{}, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("función sci necesita dos argumentos (base, exponente)")
			}
			base, ok := args[0].(float64)
			if !ok {
				return nil, fmt.Errorf("argumento base debe ser un número")
			}
			exp, ok := args[1].(float64)
			if !ok {
				return nil, fmt.Errorf("argumento exponente debe ser un número")
			}
			return math.Pow(base, exp), nil
		},
	}

	// Reemplazar la notación científica con la función personalizada
	re := regexp.MustCompile(`(\d+(\.\d+)?)([Ee])([\+\-]?\d+)`)
	expression = re.ReplaceAllString(expression, `sci($1,$4)`)
	expression = strings.ReplaceAll(expression, "–", "-") // Reemplazar guiones

	expressionParsed, err := govaluate.NewEvaluableExpressionWithFunctions(expression, functions)
	if err != nil {
		log.Fatalf("Error al analizar la expresión: %v", err)
	}

	result, err := expressionParsed.Evaluate(nil)
	if err != nil {
		log.Fatalf("Error al evaluar la expresión: %v", err)
	}

	fmt.Printf("Resultado: %v\n", result)

	// Ejemplo adicional
	expression2 := "1.23E-5 + 4.56E7"
	expression2 = re.ReplaceAllString(expression2, `sci($1,$4)`)
	expressionParsed2, err := govaluate.NewEvaluableExpressionWithFunctions(expression2, functions)
	if err != nil {
		log.Fatalf("Error al analizar la expresión 2: %v", err)
	}
	result2, err := expressionParsed2.Evaluate(nil)
	if err != nil {
		log.Fatalf("Error al evaluar la expresión 2: %v", err)
	}
	fmt.Printf("Resultado 2: %v\n", result2)
}
