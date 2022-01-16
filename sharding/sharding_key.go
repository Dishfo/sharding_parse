package sharding

import (
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/opcode"
	"github.com/pingcap/parser/test_driver"
)

type ShardingColVisitor struct {
	foundCol bool
	val      string
}

func (s *ShardingColVisitor) Enter(n ast.Node) (node ast.Node, skipChildren bool) {
	switch real := n.(type) {
	case *ast.ColumnNameExpr:
		if real.Name.Name.L == "workspace_id" || real.Name.Name.L == "company_id" {
			s.foundCol = true
			return n, true
		}
	case *test_driver.ValueExpr:
		s.val = real.GetString()
	}

	return n, false
}

func (s *ShardingColVisitor) Leave(n ast.Node) (node ast.Node, ok bool) {
	return n, true
}

type ShardingKeyVistor struct {
	shardingKeyValue string
}

func (s *ShardingKeyVistor) Key() string {
	return s.shardingKeyValue
}

func (s *ShardingKeyVistor) Enter(n ast.Node) (node ast.Node, skipChildren bool) {
	switch real := n.(type) {
	case *ast.BinaryOperationExpr:
		if real.Op == opcode.EQ {
			v := &ShardingColVisitor{}
			real.Accept(v)
			if v.foundCol {
				s.shardingKeyValue = v.val
				return n, true
			}
		}
	}

	return n, false
}

func (s *ShardingKeyVistor) Leave(n ast.Node) (node ast.Node, ok bool) {
	return n, len(s.shardingKeyValue) == 0
}
