package gounit

import (
	"testing"
)

func TestGetSum(t *testing.T) {
	sum := GetSum(2,3)
	if sum == 5{
		t.Log("测试通过")
	}else{
		t.Errorf("有点问题哦！2+3应该是%d，可是等于%d\n", 5, sum)
		return
	}
}
