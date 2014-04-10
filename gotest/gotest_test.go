package gotest

import "testing"

func Test_Div_1(t *testing.T) {
	if i, e := Div(6, 3); i != 2 || e != nil {
		t.Error("第一个测试没能通过")
	} else {
		t.Log("第一个测试通过了！")
	}
}

func Test_Div_2(t *testing.T) {
	if i, e := Div(6, 3); i != 3 || e != nil {
		t.Error("第二个测试没能通过")
	} else {
		t.Log("第二个测试通过了！")
	}
}
