package parser

import (
	"github.com/philjestin/go-interpreter/ast"
	"github.com/philjestin/go-interpreter/lexer"
	"github.com/philjestin/go-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(*lexer.Lexer) *Parser {
	p := &Parser{}

	// read two tokens so that curToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
