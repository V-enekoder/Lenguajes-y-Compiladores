package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/Knetic/govaluate"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduce la expresión matemática (soporta notación científica y paréntesis): ")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)

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
}
