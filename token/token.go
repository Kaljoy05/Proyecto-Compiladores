package token

import "strings"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL    = "ILLEGAL"
	EOF        = "EOF"
	IDENT      = "IDENT"
	INT        = "INT"
	PUNTO_COMA = ";"

	NOTA_KW  = "NOTA"
	ACORDE   = "ACORDE"
	SILENCIO = "SILENCIO"
	TEMPO    = "TEMPO"
	EN       = "EN"
	LUEGO    = "LUEGO"

	DO  = "DO"
	RE  = "RE"
	MI  = "MI"
	FA  = "FA"
	SOL = "SOL"
	LA  = "LA"
	SI  = "SI"

	REDONDA = "REDONDA"
	BLANCA  = "BLANCA"
	NEGRA   = "NEGRA"
	CORCHEA = "CORCHEA"

	SOSTENIDO = "SOSTENIDO"
	BEMOL     = "BEMOL"
	MAYOR     = "MAYOR"
	MENOR     = "MENOR"
)

var keywords = map[string]TokenType{

	"nota":     NOTA_KW,
	"acorde":   ACORDE,
	"silencio": SILENCIO,
	"tempo":    TEMPO,
	"en":       EN,
	"luego":    LUEGO,

	"do":  DO,
	"re":  RE,
	"mi":  MI,
	"fa":  FA,
	"sol": SOL,
	"la":  LA,
	"si":  SI,

	"redonda": REDONDA,
	"blanca":  BLANCA,
	"negra":   NEGRA,
	"corchea": CORCHEA,

	"sostenido": SOSTENIDO,
	"bemol":     BEMOL,
	"mayor":     MAYOR,
	"menor":     MENOR,
}

func LookupIdent(ident string) TokenType {
	identMinuscula := strings.ToLower(ident)
	if tok, ok := keywords[identMinuscula]; ok {
		return tok
	}
	return IDENT
}
