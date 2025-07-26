package eval

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// 解析器函数
func Parse(s string) (Expr, error) {
	p := parser{
		scan: scanner.Scanner{},
	}
	p.scan.Init(strings.NewReader(s))
	p.next() // 初始化获取第一个token
	expr := p.parseExpr()
	if p.token != scanner.EOF {
		return nil, fmt.Errorf("unexpected %s", p.describe())
	}
	return expr, nil
}

type parser struct {
	scan  scanner.Scanner
	token rune // 当前token
}

func (p *parser) next() {
	p.token = p.scan.Scan()
}

func (p *parser) describe() string {
	switch p.token {
	case scanner.EOF:
		return "end of input"
	case scanner.Ident:
		return fmt.Sprintf("identifier %s", p.text())
	case scanner.Int, scanner.Float:
		return fmt.Sprintf("number %s", p.text())
	}
	return fmt.Sprintf("%q", rune(p.token)) // 其他rune
}

func (p *parser) text() string {
	return p.scan.TokenText()
}

func (p *parser) parseExpr() Expr {
	return p.parseBinary(1)
}

// 解析二元表达式，prec表示优先级
func (p *parser) parseBinary(prec1 int) Expr {
	lhs := p.parseUnary()
	for {
		op := p.token
		prec := precedence(op)
		if prec < prec1 {
			return lhs
		}
		p.next() // 消费操作符
		rhs := p.parseBinary(prec + 1)
		lhs = binary{op, lhs, rhs}
	}
}

// 解析一元表达式
func (p *parser) parseUnary() Expr {
	if p.token == '+' || p.token == '-' {
		op := p.token
		p.next() // 消费 '+' 或 '-'
		return unary{op, p.parseUnary()}
	}
	return p.parsePrimary()
}

// 解析基本表达式
func (p *parser) parsePrimary() Expr {
	switch p.token {
	case scanner.Ident:
		id := p.text()
		p.next()
		if p.token != '(' {
			return Var(id)
		}
		p.next()
		var args []Expr
		if p.token != ')' {
			for {
				args = append(args, p.parseExpr())
				if p.token != ',' {
					break
				}
				p.next()
			}
			if p.token != ')' {
				msg := fmt.Sprintf("got %s, want ')'", p.describe())
				panic(fmt.Sprintf("parse error: %s", msg))
			}
		}
		p.next()
		return call{id, args}

	case scanner.Int, scanner.Float:
		f, err := strconv.ParseFloat(p.text(), 64)
		if err != nil {
			panic(fmt.Sprintf("number %q is invalid", p.text()))
		}
		p.next()
		return literal(f)

	case '(':
		p.next()
		expr := p.parseExpr()
		if p.token != ')' {
			msg := fmt.Sprintf("got %s, want ')'", p.describe())
			panic(fmt.Sprintf("parse error: %s", msg))
		}
		p.next()
		return expr
	}
	panic(fmt.Sprintf("unexpected %s", p.describe()))
}

// 运算符优先级
func precedence(op rune) int {
	switch op {
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}