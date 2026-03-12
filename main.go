package main

import (
	"fmt"
	"os"
	"simpliscore/lexer"
	"simpliscore/token"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Error: Debes indicar el archivo de texto a compilar.")
		fmt.Println("Uso correcto: go run main.go <archivo.txt>")
		return
	}

	nombreArchivo := os.Args[1]

	contenido, err := os.ReadFile(nombreArchivo)
	if err != nil {
		fmt.Printf("Error al intentar leer el archivo '%s': %v\n", nombreArchivo, err)
		return
	}

	fmt.Printf("SimpliScore: Analizando archivo '%s' \n", nombreArchivo)
	fmt.Println("-------------------------------------------------------------")

	codigoFuente := string(contenido)
	l := lexer.New(codigoFuente)

	fmt.Println("--- INICIO DE TOKENIZACIÓN ---")
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Printf("{ Tipo: %-15s | Literal: '%s' }\n", tok.Type, tok.Literal)
	}
	fmt.Println("--- FIN DE TOKENIZACIÓN ---")
}
