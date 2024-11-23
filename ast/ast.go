package ast

import "golox/lexer"

/*
	 GRAMMAR RULES
	 expression     → literal |
		| unary
		| binary
		| grouping ;

	 literal        → NUMBER | STRING | "true" | "false" | "nil" ;
	 grouping       → "(" expression ")" ;
	 unary          → ( "-" | "!" ) expression ;
	 binary         → expression operator expression ;
	 operator       → "==" | "!=" | "<" | "<=" | ">" | ">=" | "+"  | "-"  | "*" | "/" ;
*/

type Expr interface {
	accept(visitor Visitor)
}

type Visitor interface {
	visit(expr Expr)
}

type Interpreter struct{}

func (i Interpreter) visit(expr Expr) {}

type Logger struct{}

func (l Logger) visit(expr Expr) {}

func Run() {
	v := Interpreter{}
	l := Logger{}

	b := Binary{}
	u := Unary{}

	process(b, v)
	process(u, l)
}
func process(expression Expr, v Visitor) {
	expression.accept(v)
}

type Unary struct {
	operator lexer.Token
	Right    Expr
}

func (u Unary) accept(v Visitor) {
	v.visit(u)
}

type Binary struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (b Binary) accept(v Visitor) {
	v.visit(b)
}

type Grouping struct {
	OPEN       string
	expression Expr
	CLOSE      string
}

type Literal struct {
	Value interface{}
}
