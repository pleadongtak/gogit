package calc

import "testing"

func TestName(t *testing.T) {

	r := Sum(1,2)
	if r  != 3{
		t.Error(" 3이 아닙니다.")
	}

}