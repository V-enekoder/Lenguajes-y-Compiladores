package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

// limpiarPalabra limpia una palabra de puntuación y la convierte a minúsculas.
func limpiarPalabra(palabra string) string {
	var resultado strings.Builder
	for _, r := range palabra {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			resultado.WriteString(strings.ToLower(string(r)))
		}
	}
	return resultado.String()
}

func main() {
	// Texto de ejemplo (reemplaza con la entrada del usuario si lo deseas)
	texto := `#include <iostream>

using namespace std;

int main() {
  double num1, num2, resultado;
  char operador;

  cout << "Introduce el primer número: ";
  cin >> num1;
  cout << "Introduce el operador (+, -, *, /): ";
  cin >> operador;
  cout << "Introduce el segundo número: ";
  cin >> num2;

  switch (operador) {
    case '+':
      resultado = num1 + num2;
      break;
    case '-':
      resultado = num1 - num2;
      break;
    case '*':
      resultado = num1 * num2;
      break;
    case '/':
      resultado = num1 / num2;
      break;
    default:
      cout << "Operador inválido." << endl;
      return 1; // Indica un error
  }

  cout << "El resultado es: " << resultado << endl;
  return 0;
}`

	// Lee las palabras reservadas del usuario.
	fmt.Print("Introduce las palabras reservadas (separadas por espacios): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Eliminar espacios en blanco al principio y al final

	palabrasReservadas := strings.Fields(input)
	for i, palabra := range palabrasReservadas {
		palabrasReservadas[i] = limpiarPalabra(palabra)
	}

	// Procesar el texto
	palabrasTexto := regexp.MustCompile(`\b\w+\b`).FindAllString(texto, -1) // Encuentra todas las palabras
	ocurrencias := make(map[string]int)                                     // Mapa para almacenar las ocurrencias

	for _, palabra := range palabrasTexto {
		limpia := limpiarPalabra(palabra)
		ocurrencias[limpia]++
	}

	fmt.Println("\nPalabras reservadas encontradas y sus ocurrencias:")
	for _, reservada := range palabrasReservadas {
		if count, ok := ocurrencias[reservada]; ok {
			fmt.Printf("%s: %d\n", reservada, count)
		}
	}

	if len(palabrasReservadas) == 0 {
		fmt.Println("\nNo se ingresaron palabras reservadas.")
	}
}

/*
auto break case char const continue default do double else enum extern float for goto if int long register return short signed sizeof static struct switch typedef union unsigned void volatile while

*/
