package main

import (
	"fmt"
	"os"
	"simpliscore/lexer"
	"simpliscore/parser"
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

	
	program := p.ParseProgram()

	
	if len(p.Errors()) != 0 {
		fmt.Println("Se encontraron errores de sintaxis en el archivo:")
		for _, msg := range p.Errors() {
			fmt.Printf("\t- %s\n", msg)
		}
		return
	}

	fmt.Println("Análisis sintáctico exitoso El Árbol generado es:")
	fmt.Println("--- INICIO DE INSTRUCCIONES ---")

	for i, stmt := range program.Statements {
		
		fmt.Printf("Instrucción %d: %s\n", i+1, stmt.String())
	}

	fmt.Println("--- FIN DE INSTRUCCIONES ---")
}
