package generator

import (
	"bytes"
	"fmt"
	"simpliscore/ast"
)

var notasMap = map[string]string{
	"do": "c", "re": "d", "mi": "e", "fa": "f",
	"sol": "g", "la": "a", "si": "b",
}

var duracionesMap = map[string]string{
	"redonda": "1", "blanca": "2", "negra": "4", "corchea": "8",
}

var alteracionesMap = map[string]string{
	"sostenido": "is", "bemol": "es",
}

var octavasMap = map[string]string{
	"aguda": "'",
	"grave": ",",
}

func Generate(node ast.Node) string {
	switch n := node.(type) {
	case *ast.Program:
		return generateProgram(n)
	case *ast.NotaStatement:
		return generateNota(n)
	case *ast.SilencioStatement:
		return generateSilencio(n)
	case *ast.TempoStatement:
		return generateTempo(n)
	case *ast.AcordeStatement:
		return generateAcorde(n)
	default:
		return ""
	}
}

func generateProgram(p *ast.Program) string {
	var out bytes.Buffer

	out.WriteString("\\version \"2.24.1\"\n")
	out.WriteString("\\score {\n")
	out.WriteString("  \\fixed c' {\n")
	out.WriteString("    ")

	for _, stmt := range p.Statements {
		out.WriteString(Generate(stmt))
		out.WriteString(" ")
	}

	out.WriteString("\n  }\n")
	out.WriteString("  \\layout { }\n")
	out.WriteString("  \\midi { }\n")
	out.WriteString("}\n")

	return out.String()
}

func generateNota(n *ast.NotaStatement) string {
	nota := notasMap[n.Nota.Literal]

	alteracion := ""
	if n.Alteracion.Type != "" {
		alteracion = alteracionesMap[n.Alteracion.Literal]
	}

	octava := ""
	if n.Octava.Type != "" {
		octava = octavasMap[n.Octava.Literal]
	}

	duracion := duracionesMap[n.Duracion.Literal]

	return nota + alteracion + octava + duracion
}

func generateSilencio(s *ast.SilencioStatement) string {
	return "r" + duracionesMap[s.Duracion.Literal]
}

func generateTempo(t *ast.TempoStatement) string {

	return fmt.Sprintf("\\tempo 4 = %s", t.Valor)
}

func generateAcorde(a *ast.AcordeStatement) string {
	raiz := notasMap[a.Raiz.Literal]

	tipo := ""
	if a.Tipo.Literal == "menor" {
		tipo = ":m"
	}
	return fmt.Sprintf("\\chordmode { %s%s }", raiz, tipo)
}
