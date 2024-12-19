package runtime

import (
	"github.com/smlgh/smarti/internal/ast"
	"github.com/smlgh/smarti/internal/packages"
)

type variable struct {
	Type  ast.NodeType
	Value interface{}
	Ref   bool
}

func toPkgVar(v []variable) []packages.Variable {
	var vars []packages.Variable

	for _, vv := range v {
		vars = append(vars, packages.Variable{
			Type:  packages.VarType(vv.Type),
			Value: vv.Value,
		})
	}

	return vars
}

func toNodeType(v packages.VarType) ast.NodeType {
	return ast.NodeType(v)
}
