package main

import (
	"fmt"
	"os"
	"os/exec"
	"simpliscore/generator"
	"simpliscore/lexer"
	"simpliscore/parser"
	"strings"
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

	fmt.Printf("SimpliScore: Transpilando '%s'...\n", nombreArchivo)

	codigoFuente := string(contenido)
	l := lexer.New(codigoFuente)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("¡Ups! Se encontraron errores de sintaxis:")
		for _, msg := range p.Errors() {
			fmt.Printf("\t- %s\n", msg)
		}
		return
	}

	codigoLilyPond := generator.Generate(program)

	nombreSalida := strings.Replace(nombreArchivo, ".txt", ".ly", 1)
	if !strings.Contains(nombreSalida, ".ly") {
		nombreSalida += ".ly"
	}

	err = os.WriteFile(nombreSalida, []byte(codigoLilyPond), 0644)
	if err != nil {
		fmt.Printf("Error al guardar el archivo generado: %v\n", err)
		return
	}

	fmt.Println("-------------------------------------------------------------")
	fmt.Printf("¡Éxito! Archivo fuente generado: %s\n", nombreSalida)
	fmt.Println("SimpliScore: Invocando al motor de LilyPond para generar PDF y MIDI...")

	cmd := exec.Command("lilypond", nombreSalida)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		fmt.Println("\n[ERROR] El compilador interno falló al intentar ejecutar LilyPond.")
		fmt.Printf("Detalle: %v\n", err)
		fmt.Println("Asegúrate de que LilyPond esté instalado y agregado al PATH (Windows) o instalado vía apt (WSL).")
		return
	}

	fmt.Println("-------------------------------------------------------------")
	fmt.Println("¡Compilación finalizada con éxito! Revisa tu carpeta para ver la partitura y escuchar el audio.")
}
