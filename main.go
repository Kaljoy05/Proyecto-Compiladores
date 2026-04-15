package main

import (
	"fmt"
	"os"
	"simpliscore/lexer"
	"simpliscore/parser" // ¡No olvides importar tu nuevo paquete parser!
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

	fmt.Printf("SimpliScore: Analizando sintaxis de '%s' \n", nombreArchivo)

	codigoFuente := string(contenido)
	l := lexer.New(codigoFuente)
	p := parser.New(l)

	// Iniciamos el análisis sintáctico
	program := p.ParseProgram()

	// 1. Verificamos si el usuario cometió errores de sintaxis
	if len(p.Errors()) != 0 {
		fmt.Println("Se encontraron errores de sintaxis en el archivo:")
		for _, msg := range p.Errors() {
			fmt.Printf("\t- %s\n", msg)
		}
		return
	}

	// 2. Si no hay errores, mostramos el AST resultante
	fmt.Println("Análisis sintáctico exitoso El Árbol generado es:")
	fmt.Println("--- INICIO DE INSTRUCCIONES ---")

	for i, stmt := range program.Statements {
		// Al llamar a String() estamos usando los métodos que definimos en ast.go
		fmt.Printf("Instrucción %d: %s\n", i+1, stmt.String())
	}

	fmt.Println("--- FIN DE INSTRUCCIONES ---")
}
