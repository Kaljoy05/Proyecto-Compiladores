# SimpliScore 🎵
**Transpilador de Lenguaje Musical a LilyPond**

**Fase Actual:** Proyecto 100% Funcional (Fase Léxica, Sintáctica, Generación de Código y Orquestación).

## Descripción del Proyecto
SimpliScore es un transpilador escrito en Go (Golang) diseñado para traducir un Lenguaje de Dominio Específico (DSL) basado en notación musical natural en español, hacia el lenguaje de marcado técnico y profesional **LilyPond** (`.ly`).

El objetivo principal del proyecto es ofrecer una experiencia de usuario (DX) fluida y perdonadora para músicos y compositores, abstrayendo la compleja curva de aprendizaje de LilyPond. El programa lee el código fuente, construye un **Árbol de Sintaxis Abstracta (AST)** tolerante a fallos, y finalmente automatiza la generación de partituras (PDF) y pistas de audio (MIDI) interactuando directamente con el sistema operativo.

## ✨ Características Principales
* **Sintaxis en Español y Natural:** Redacción de partituras como si fuera un dictado (ej. `Nota do aguda negra;`).
* **Tolerancia a Fallos:** El analizador léxico es insensible a mayúsculas/minúsculas.
* **Azúcar Sintáctica:** Uso de conectores opcionales como `LUEGO` para mejorar la fluidez de lectura (el AST los ignora para mantenerse puro).
* **Punto y Coma Opcional:** Soporte para separar instrucciones con saltos de línea.
* **Ejecución Zero-Config:** Orquestación automática de subprocesos (`os/exec`). Genera el PDF y el MIDI sin que el usuario deba interactuar con herramientas externas.

## 🗂️ Estructura del Proyecto

📁 simpliscore/
 ├── 📄 go.mod           # Definición del módulo de Go
 ├── 📄 main.go          # Punto de entrada y orquestador del compilador
 ├── 📄 prueba.txt       # Archivo con casos de uso del DSL (Input)
 ├── 📄 README.md        # Documentación del proyecto
 │
 ├── 📁 token/
 │    └── 📄 token.go    # Diccionario de datos y definición de la estructura Token
 │
 ├── 📁 lexer/
 │    └── 📄 lexer.go    # Motor del Scanner (tokenización y normalización del texto)
 │
 ├── 📁 ast/
 │    └── 📄 ast.go      # Definición de los Nodos y Estructuras del Árbol Sintáctico
 │
 ├── 📁 parser/
 │    └── 📄 parser.go   # Analizador Sintáctico Descendente para construir el AST
 │
 └── 📁 generator/
      └── 📄 generator.go # Backend: Traducción a LilyPond e inyección de plantillas

## ⚙️ Requisitos Previos e Instalación

Para ejecutar SimpliScore, necesitas tener instalados **Go** y **LilyPond**.

### 1. Instalar Go
Asegúrate de tener Go instalado (versión 1.20 o superior).

### 2. Instalar LilyPond (¡Importante!)
El transpilador depende del motor externo de LilyPond para renderizar el resultado final. 

#### 🐧 Para Linux (Ubuntu / Debian / WSL):
```bash
sudo apt update
sudo apt install lilypond

Para Windows:
Descarga la versión .zip desde la página oficial de LilyPond.

Extrae la carpeta en tu disco duro (ej. C:\lilypond-2.24.1).

Agregar al PATH (Obligatorio):

Abre "Variables de entorno" en Windows.

En "Variables del sistema", edita la variable Path.

Agrega la ruta completa hacia la carpeta bin de LilyPond (ej. C:\lilypond-2.24.1\bin).

Abre una nueva terminal para que los cambios apliquen.

🚀 Uso
Crea un archivo de texto (ej. prueba.txt) y escribe tu composición.

Abre la terminal en la raíz del proyecto y ejecuta: go run main.go prueba.txt

SimpliScore escaneará el texto, validará la gramática, transpilara el código a .ly y ejecutará LilyPond automáticamente para generar prueba.pdf y prueba.midi en tu carpeta.

📖 Guía Rápida del Lenguaje
El lenguaje soporta comentarios utilizando //.

Estructuras Básicas:

Tempo en [X]: Define la velocidad de la pieza.

Nota [altura] [octava opcional] [duración]: Define una nota individual.

Acorde [altura] [mayor/menor] [duración]: Define un acorde.

Silencio [duración]: Define una pausa.

Diccionario Permitido:

Alturas: do, re, mi, fa, sol, la, si.

Alteraciones (Opcionales): sostenido, bemol.

Octavas (Opcionales): aguda, grave.

Duraciones: redonda, blanca, negra, corchea.