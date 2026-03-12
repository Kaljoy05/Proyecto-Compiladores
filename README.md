# Proyecto-Compiladores
# SimpliScore: Transpilador de Lenguaje Musical a LilyPond

**Fase Actual:** Fase 1 - Analizador Léxico (Scanner) 100% Funcional.

## Descripción del Proyecto
SimpliScore es un transpilador escrito en Go diseñado para traducir un Lenguaje de Dominio Específico (DSL) basado en notación musical en español, hacia el lenguaje de grabado musical profesional **LilyPond** (`.ly`).

El objetivo de esta primera fase es el **Análisis Léxico**. El programa lee un archivo de texto con código fuente en el dialecto de SimpliScore, ignora los espacios o saltos de línea, e identifica y agrupa los caracteres en unidades lógicas llamadas **Tokens** (Keywords, Notas, Duraciones, Alteraciones, Números y Símbolos).

## Estructura del Proyecto

📁 simpliscore/
 ├── 📄 go.mod           # Definición del módulo de Go
 ├── 📄 main.go          # Punto de entrada y lectura de archivos
 ├── 📄 pruebas.txt      # Archivo de pruebas con casos de uso del DSL
 ├── 📄 README.md        # Documentación del proyecto
 │
 ├── 📁 token/
 │    └── 📄 token.go    # Diccionario de datos y definición de la estructura Token
 │
 └── 📁 lexer/
      └── 📄 lexer.go    # Motor del Scanner (lectura por caracteres y evaluación)


## Requisitos
* Tener instalado **Go 1.18** o superior.

## ¿Cómo ejecutar el Analizador Léxico?

El programa está diseñado para ejecutarse desde la terminal pasándole como argumento el archivo de texto que contiene el código fuente musical.

1. Abre la terminal en la raíz del proyecto.
2. Ejecuta el siguiente comando:

```bash
go run main.go pruebas.txt
