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
	for i, palabra := range palabrasTexto {
		palabrasTexto[i] = limpiarPalabra(palabra)
	}

	fmt.Println("Palabras reservadas encontradas:")
	for _, reservada := range palabrasReservadas {
		for _, palabraActual := range palabrasTexto {
			if palabraActual == reservada {
				fmt.Println(reservada)
				goto nextReservedWord // Sale del bucle interno una vez que se encuentra la palabra
			}
		}
	nextReservedWord: // Etiqueta para el goto
	}

	if len(palabrasReservadas) == 0 {
		fmt.Println("No se ingresaron palabras reservadas.")
	}
}
