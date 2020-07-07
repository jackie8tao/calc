package ast

import "testing"

// 100 + 50 + 200
func Test_AstWalk(t *testing.T) {
	a1 := NewIntExpr(100)
	a2 := NewIntExpr(50)
	a3 := NewIntExpr(200)
	sum1 := NewAddExpr(a1, a2)
	sum2 := NewAddExpr(sum1, a3)

	if sum2.Walk() != 350 {
		t.Fail()
	}
}
