package main

import (
	"fmt"
	"github.com/Dishfo/sharding_parse/sharding"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	"github.com/pingcap/parser/format"
	_ "github.com/pingcap/parser/test_driver"
	"log"
	"strings"
)

func parse(sql string) (*ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return &stmtNodes[0], nil
}
//join t2 on t2.id = t.rid join t3 on t3.id = t2.rid
func main() {
	astNode, err := parse("SELECT a, b FROM t where workspace_id = 'dawdwadwadd'")
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}
	var builder strings.Builder
	restoreCtx := format.NewRestoreCtx(format.RestoreKeyWordUppercase, &builder)
	selNode := (*astNode).(*ast.SelectStmt)

	err = selNode.Restore(restoreCtx)
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}

	log.Println(builder.String())

	fmt.Printf("%v\n", *astNode)
	v := sharding.ShardingKeyVistor{}

	selNode.Accept(&v)
	log.Println(v.Key())

}
