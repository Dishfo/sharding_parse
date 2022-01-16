package sharding

import (
	"context"
	"github.com/pingcap/parser/ast"
)

//poc of
type ShardingContext struct {
	ctx           context.Context
	specificKeyId string
	stmt          *ast.SelectStmt

}




func (s *ShardingContext) parseShardingKey() {



}
