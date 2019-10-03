package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	sing1 := GetInstance(123)
	sing2 := GetInstance(123)

	if sing1 != sing2{
		t.Errorf("两个实例不一样")
	}else {
		t.Log("两个实例一样")
	}
}
