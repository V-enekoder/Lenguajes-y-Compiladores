package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isValidFEN(fen string) bool {
	parts := strings.Split(fen, " ")
	if len(parts) != 6 {
		return false // Debe tener 6 partes
	}

	// 1. Posición de las piezas
	if !isValidPiecePlacement(parts[0]) {
		return false
	}

	// 2. Jugador que juega
	if parts[1] != "w" && parts[1] != "b" {
		return false // Debe ser "w" (blancas) o "b" (negras)
	}

	// 3. Enroque
	if !isValidCastling(parts[2]) {
		return false
	}

	// 4. En passant
	if !isValidEnPassant(parts[3]) {
		return false
	}

	// 5. Movimiento del 50-movimiento
	if !isValidFiftyMove(parts[4]) {
		return false
	}

	// 6. Movimiento del reloj
	if !isValidFullmove(parts[5]) {
		return false
	}

	return true
}

func isValidPiecePlacement(placement string) bool {
	rows := strings.Split(placement, "/")
	if len(rows) != 8 {
		return false
	}

	pieceRegex := regexp.MustCompile(`^[rnbqkpRNBQKP]+$`)

	for _, row := range rows {
		count := 0
		for _, char := range row {
			if char >= '1' && char <= '8' {
				count += int(char - '0')
				if count > 8 {
					return false
				}
			} else if !pieceRegex.MatchString(string(char)) { //Validacion de pieza aqui
				return false // Caracter inválido (no es una pieza)
			} else {
				count++
			}
		}
		if count != 8 {
			return false
		}
	}
	return true
}

func isValidCastling(castling string) bool {
	if castling == "" {
		return true
	}
	re := regexp.MustCompile(`^([KQkq]{1,4})$`)
	return re.MatchString(castling)
}

func isValidEnPassant(enPassant string) bool {
	re := regexp.MustCompile(`^([a-h][36]|-)$`) // a-h[36] o "-"
	return re.MatchString(enPassant)
}
func isValidFiftyMove(fiftyMove string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(fiftyMove)
}

func isValidFullmove(fullMove string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(fullMove)
}

func main() {
	fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 0 1"
	fmt.Println(fen, "es válido:", isValidFEN(fen))

	fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1  extra"
	fmt.Println(fen, "es válido:", isValidFEN(fen))

	fen = "rnbqkbnr/ppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	fmt.Println(fen, "es válido:", isValidFEN(fen))

	fen = "rnbqkbnr/pppp1ppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	fmt.Println(fen, "es válido:", isValidFEN(fen))

	fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR p KQkq - 0 1"
	fmt.Println(fen, "es válido:", isValidFEN(fen))

	fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w Kkq - 0 1"
	fmt.Println(fen, "es válido:", isValidFEN(fen))

	fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkqa - 0 1" //Enroque invalido
	fmt.Println(fen, "es válido:", isValidFEN(fen))

	fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w  - 0 1"
	fmt.Println(fen, "es válido:", isValidFEN(fen))

	fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR b KQkq - 50 10"
	fmt.Println(fen, "es válido:", isValidFEN(fen))

	fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	fmt.Println(fen, "es válido:", isValidFEN(fen))

	fen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 a"
	fmt.Println(fen, "es válido:", isValidFEN(fen))

}
