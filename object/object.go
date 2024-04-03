package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"monkey/ast"
	"strings"
)

type TypeObject string
type BuiltinFunction func(args ...Object) Object

const (
	IntegerObj         = "INTEGER"
	BooleanObj         = "BOOLEAN"
	NullObj            = "NULL"
	ReturnValueObj     = "RETURN_VALUE"
	ErrorObj           = "ERROR"
	FunctionObj        = "FUNCTION"
	StingObj           = "STRING"
	BuiltinFunctionObj = "BUILTIN"
	ArrayObj           = "ARRAY"
	HashObj            = "HASH"
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

type String struct {
	Value string
}

type HashKey struct {
	Type  TypeObject
	Value uint64
}

type Builtin struct {
	Fn BuiltinFunction
}
type Array struct {
	Elements []Object
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hashable interface {
	HashKey() HashKey
}

type Hash struct {
	Pairs map[HashKey]HashPair
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

func (s *String) Type() TypeObject {
	return StingObj
}

func (s *String) Inspect() string {
	return s.Value
}

func (b *Builtin) Type() TypeObject {
	return BuiltinFunctionObj
}

func (b *Builtin) Inspect() string {
	return "builtin function"
}

func (ao *Array) Type() TypeObject { return ArrayObj }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	var elements []string
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (b *Boolean) HashKey() HashKey {
	var value uint64
	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	_, err := h.Write([]byte(s.Value))
	if err != nil {
		return HashKey{}
	}

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

func (h *Hash) Type() TypeObject { return HashObj }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	var pairs []string

	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s:%s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}
