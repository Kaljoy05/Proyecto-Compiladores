package parser

import (
	"simpliscore/ast"
	"simpliscore/lexer"
	"testing"
)

func TestParseNotaStatement(t *testing.T) {
	// Definimos varios escenarios para poner a prueba el parser
	tests := []struct {
		input              string
		expectedNota       string
		expectedAlteracion string
		expectedDuracion   string
	}{
		{"nota do negra;", "do", "", "negra"},
		{"nota fa sostenido blanca;", "fa", "sostenido", "blanca"},
		{"nota si bemol corchea;", "si", "bemol", "corchea"},
		{"nota sol redonda;", "sol", "", "redonda"},
		// Opcional: Probar con mayúsculas si tu lexer ya lo soporta bien
		{"nota RE blanca;", "RE", "", "blanca"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()

		// Verificamos que no haya errores de sintaxis
		checkParserErrors(t, p)

		if program == nil {
			t.Fatalf("ParseProgram() devolvió nil")
		}
		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements no contiene 1 sentencia. se obtuvo=%d", len(program.Statements))
		}

		stmt := program.Statements[0]
		notaStmt, ok := stmt.(*ast.NotaStatement)
		if !ok {
			t.Fatalf("stmt no es *ast.NotaStatement. se obtuvo=%T", stmt)
		}

		if notaStmt.TokenLiteral() != "nota" {
			t.Errorf("notaStmt.TokenLiteral no es 'nota'. se obtuvo=%q", notaStmt.TokenLiteral())
		}

		if notaStmt.Nota.Literal != tt.expectedNota {
			t.Errorf("notaStmt.Nota.Literal no es '%s'. se obtuvo=%q", tt.expectedNota, notaStmt.Nota.Literal)
		}

		if notaStmt.Alteracion.Literal != tt.expectedAlteracion {
			t.Errorf("notaStmt.Alteracion.Literal no es '%s'. se obtuvo=%q", tt.expectedAlteracion, notaStmt.Alteracion.Literal)
		}

		if notaStmt.Duracion.Literal != tt.expectedDuracion {
			t.Errorf("notaStmt.Duracion.Literal no es '%s'. se obtuvo=%q", tt.expectedDuracion, notaStmt.Duracion.Literal)
		}
	}
}

// Función auxiliar para imprimir los errores del parser si la prueba falla
func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("el parser tiene %d errores", len(errors))
	for _, msg := range errors {
		t.Errorf("error del parser: %q", msg)
	}
	t.FailNow()
}

func TestParseAcordeStatement(t *testing.T) {
	input := `
	acorde sol mayor ;
	acorde RE menor ;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements no contiene 2 sentencias. se obtuvo=%d", len(program.Statements))
	}

	// Prueba 1: acorde sol mayor
	acorde1 := program.Statements[0].(*ast.AcordeStatement)
	if acorde1.Raiz.Literal != "sol" || acorde1.Tipo.Literal != "mayor" {
		t.Errorf("Error en acorde 1. Raiz: %s, Tipo: %s", acorde1.Raiz.Literal, acorde1.Tipo.Literal)
	}

	// Prueba 2: acorde RE menor (recordando lo de las mayúsculas)
	acorde2 := program.Statements[1].(*ast.AcordeStatement)
	if acorde2.Raiz.Literal != "RE" || acorde2.Tipo.Literal != "menor" {
		t.Errorf("Error en acorde 2. Raiz: %s, Tipo: %s", acorde2.Raiz.Literal, acorde2.Tipo.Literal)
	}
}

func TestParseSilencioYTempo(t *testing.T) {
	input := `
	silencio blanca ;
	tempo en 120 ;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements no contiene 2 sentencias. se obtuvo=%d", len(program.Statements))
	}

	// Prueba Silencio
	silencio := program.Statements[0].(*ast.SilencioStatement)
	if silencio.Duracion.Literal != "blanca" {
		t.Errorf("Error en silencio. Duracion: %s", silencio.Duracion.Literal)
	}

	// Prueba Tempo
	tempo := program.Statements[1].(*ast.TempoStatement)
	if tempo.Valor != "120" {
		t.Errorf("Error en tempo. Valor: %s", tempo.Valor)
	}
}
