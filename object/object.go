package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

type TypeObject string

const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
	FunctionObj    = "FUNCTION"
)

type Object interface {
	Type() TypeObject
	Inspect() string
}

type Integer struct {
	Value int64
}

type Boolean struct {
	Value bool
}

type Null struct{}

type ReturnValue struct {
	Value Object
}

type Error struct {
	Message string
}

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() TypeObject {
	return IntegerObj
}

func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Type() TypeObject {
	return BooleanObj
}

func (n *Null) Inspect() string {
	return "null"
}

func (n *Null) Type() TypeObject {
	return NullObj
}

func (rv *ReturnValue) Type() TypeObject {
	return ReturnValueObj
}

func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

func (e *Error) Type() TypeObject {
	return ErrorObj
}

func (e *Error) Inspect() string {
	return "ERROR: " + e.Message
}

func (f *Function) Type() TypeObject {
	return FunctionObj
}

func (f *Function) Inspect() string {
	var out bytes.Buffer
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}
