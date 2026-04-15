package ast

import (
	"bytes"
	"simpliscore/token"
)

// Interfaces base
type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String() + " ")
	}
	return out.String()
}

// NODO PARA NOTAS
type NotaStatement struct {
	Token      token.Token
	Nota       token.Token
	Alteracion token.Token
	Octava     token.Token // Soporte para el número de octava
	Duracion   token.Token
}

func (ns *NotaStatement) statementNode()       {}
func (ns *NotaStatement) TokenLiteral() string { return ns.Token.Literal }
func (ns *NotaStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ns.TokenLiteral() + " " + ns.Nota.Literal)
	if ns.Alteracion.Literal != "" {
		out.WriteString(" " + ns.Alteracion.Literal)
	}
	if ns.Octava.Literal != "" {
		out.WriteString(" " + ns.Octava.Literal)
	}
	out.WriteString(" " + ns.Duracion.Literal + ";")
	return out.String()
}

// NODO PARA ACORDES
type AcordeStatement struct {
	Token    token.Token
	Raiz     token.Token
	Octava   token.Token // Soporte para la octava base del acorde
	Tipo     token.Token
	Duracion token.Token
}

func (as *AcordeStatement) statementNode()       {}
func (as *AcordeStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AcordeStatement) String() string {
	var out bytes.Buffer
	out.WriteString(as.TokenLiteral() + " " + as.Raiz.Literal)
	if as.Octava.Literal != "" {
		out.WriteString(" " + as.Octava.Literal)
	}
	out.WriteString(" " + as.Tipo.Literal + " " + as.Duracion.Literal + ";")
	return out.String()
}

// NODO PARA SILENCIOS
type SilencioStatement struct {
	Token    token.Token
	Duracion token.Token
}

func (ss *SilencioStatement) statementNode()       {}
func (ss *SilencioStatement) TokenLiteral() string { return ss.Token.Literal }
func (ss *SilencioStatement) String() string {
	return ss.TokenLiteral() + " " + ss.Duracion.Literal + ";"
}

// NODO PARA TEMPO
type TempoStatement struct {
	Token token.Token
	Valor string
}

func (ts *TempoStatement) statementNode()       {}
func (ts *TempoStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *TempoStatement) String() string {
	return ts.TokenLiteral() + " en " + ts.Valor + ";"
}
