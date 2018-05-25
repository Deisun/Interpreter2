package parser

import (
	"Interpreter2/lexer"
	"Interpreter2/token"
	"Interpreter2/ast"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}


func New(l *lexer.Lexer) *Parser {
	p := &Parser{l:l}

	// read two tokens so both curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}


func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}


func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

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
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}


func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token:p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
