# Proyecto-Compiladores
# SimpliScore: Transpilador de Lenguaje Musical a LilyPond

**Fase Actual:** Fase 2 - Analizador Sintáctico (Parser) y AST 100% Funcional.

## Descripción del Proyecto
SimpliScore es un transpilador escrito en Go diseñado para traducir un Lenguaje de Dominio Específico (DSL) basado en notación musical en español, hacia el lenguaje de grabado musical profesional **LilyPond** (`.ly`).

El proyecto ha superado el **Análisis Léxico** y actualmente implementa un **Análisis Sintáctico** completo. El programa lee el código fuente, genera los tokens y los agrupa lógicamente para construir un **Árbol de Sintaxis Abstracta (AST)**. El parser es capaz de entender comandos estructurados de notas (con octavas y alteraciones opcionales), acordes con duraciones, silencios y métricas de tempo, detectando errores gramaticales e ignorando inteligentemente el "azúcar sintáctica" (como la palabra `luego`) para mantener el árbol de compilación limpio.

## Estructura del Proyecto

📁 simpliscore/
 ├── 📄 go.mod           # Definición del módulo de Go
 ├── 📄 main.go          # Punto de entrada y ejecución del compilador
 ├── 📄 pruebas.txt      # Archivo con casos de uso válidos e inválidos del DSL
 ├── 📄 README.md        # Documentación del proyecto
 │
 ├── 📁 token/
 │    └── 📄 token.go    # Diccionario de datos y definición de la estructura Token
 │
 ├── 📁 lexer/
 │    └── 📄 lexer.go    # Motor del Scanner (tokenización del texto)
 │
 ├── 📁 ast/
 │    ├── 📄 ast.go      # Definición de los Nodos y Estructuras del Árbol Sintáctico
 │    └── 📄 ast_test.go # Pruebas unitarias de representación del AST
 │
 └── 📁 parser/
      ├── 📄 parser.go   # Analizador Sintáctico para construir el AST
      └── 📄 parser_test.go # Pruebas unitarias de las sentencias musicales

## Requisitos
* Tener instalado **Go 1.18** o superior.

## ¿Cómo ejecutar el Analizador Sintáctico?

El programa está diseñado para ejecutarse desde la terminal pasándole como argumento el archivo de texto que contiene el código fuente musical. Al ejecutarlo, validará las reglas de la gramática e imprimirá la estructura lógica del AST generado, o en su defecto, reportará los errores de sintaxis exactos.

1. Abre la terminal en la raíz del proyecto.
2. Ejecuta el siguiente comando:

```bash
go run main.go pruebas.txt
