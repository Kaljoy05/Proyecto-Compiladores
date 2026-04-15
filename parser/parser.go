package parser

import (
	"fmt"
	"simpliscore/ast"
	"simpliscore/lexer"
	"simpliscore/token"
)

type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string { return p.errors }

func (p *Parser) peekTokenIs(t token.TokenType) bool { return p.peekToken.Type == t }

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	msg := fmt.Sprintf("se esperaba el token %s, pero se obtuvo %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
	return false
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}
	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.NOTA_KW:
		return p.parseNotaStatement()
	case token.ACORDE:
		return p.parseAcordeStatement()
	case token.SILENCIO:
		return p.parseSilencioStatement()
	case token.TEMPO:
		return p.parseTempoStatement()
	case token.LUEGO:
		p.nextToken()
		return p.parseStatement()
	default:
		p.errors = append(p.errors, fmt.Sprintf("comando no reconocido: %s", p.curToken.Literal))
		return nil
	}
}

func (p *Parser) parseNotaStatement() *ast.NotaStatement {
	stmt := &ast.NotaStatement{Token: p.curToken}
	p.nextToken()

	if !isNotaValida(p.curToken.Type) {
		p.errors = append(p.errors, fmt.Sprintf("se esperaba una nota, se obtuvo '%s'", p.curToken.Literal))
		return nil
	}
	stmt.Nota = p.curToken
	p.nextToken()


	if p.curToken.Type == token.SOSTENIDO || p.curToken.Type == token.BEMOL {
		stmt.Alteracion = p.curToken
		p.nextToken()
	}

	
	if p.curToken.Type == token.INT {
		stmt.Octava = p.curToken
		p.nextToken()
	}

	if !isDuracionValida(p.curToken.Type) {
		p.errors = append(p.errors, fmt.Sprintf("se esperaba una duración, se obtuvo '%s'", p.curToken.Literal))
		return nil
	}
	stmt.Duracion = p.curToken

	if !p.expectPeek(token.PUNTO_COMA) {
		return nil
	}
	return stmt
}

func (p *Parser) parseAcordeStatement() *ast.AcordeStatement {
	stmt := &ast.AcordeStatement{Token: p.curToken}
	p.nextToken()

	if !isNotaValida(p.curToken.Type) {
		p.errors = append(p.errors, fmt.Sprintf("se esperaba raíz del acorde, obtuvo '%s'", p.curToken.Literal))
		return nil
	}
	stmt.Raiz = p.curToken
	p.nextToken()

	// Octava opcional para la raíz del acorde
	if p.curToken.Type == token.INT {
		stmt.Octava = p.curToken
		p.nextToken()
	}

	if p.curToken.Type != token.MAYOR && p.curToken.Type != token.MENOR {
		p.errors = append(p.errors, fmt.Sprintf("se esperaba mayor/menor, obtuvo '%s'", p.curToken.Literal))
		return nil
	}
	stmt.Tipo = p.curToken
	p.nextToken()

	if !isDuracionValida(p.curToken.Type) {
		p.errors = append(p.errors, fmt.Sprintf("se esperaba duración del acorde, obtuvo '%s'", p.curToken.Literal))
		return nil
	}
	stmt.Duracion = p.curToken

	if !p.expectPeek(token.PUNTO_COMA) {
		return nil
	}
	return stmt
}

func (p *Parser) parseSilencioStatement() *ast.SilencioStatement {
	stmt := &ast.SilencioStatement{Token: p.curToken}
	p.nextToken()

	if !isDuracionValida(p.curToken.Type) {
		p.errors = append(p.errors, fmt.Sprintf("se esperaba duración del silencio, obtuvo '%s'", p.curToken.Literal))
		return nil
	}
	stmt.Duracion = p.curToken

	if !p.expectPeek(token.PUNTO_COMA) {
		return nil
	}
	return stmt
}

func (p *Parser) parseTempoStatement() *ast.TempoStatement {
	stmt := &ast.TempoStatement{Token: p.curToken}

	if !p.expectPeek(token.EN) {
		return nil
	}
	if !p.expectPeek(token.INT) {
		return nil
	}
	stmt.Valor = p.curToken.Literal

	if !p.expectPeek(token.PUNTO_COMA) {
		return nil
	}
	return stmt
}

func isNotaValida(t token.TokenType) bool {
	return t == token.DO || t == token.RE || t == token.MI || t == token.FA || t == token.SOL || t == token.LA || t == token.SI
}

func isDuracionValida(t token.TokenType) bool {
	return t == token.REDONDA || t == token.BLANCA || t == token.NEGRA || t == token.CORCHEA
}
