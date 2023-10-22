package vector_inspector

import (
	"github.com/koykov/dyntpl"
	"github.com/koykov/vector"
)

func VectorNodeEmptyCheck(_ *dyntpl.Ctx, val any) bool {
	if node, ok := val.(*vector.Node); ok {
		return node.Type() == vector.TypeNull
	}
	return false
}
