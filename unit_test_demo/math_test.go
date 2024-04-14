package main

import(
	"testing"
	// "./math" 
)

func Add(a,b int)int  {
	return a+b
}

func TestAdd(t *testing.T)  {
	// 测试用例 1
	result := Add(1, 2)
	expected := 3
	if result != expected {
		t.Errorf("Add(1, 2) 返回值错误，期望 %d，实际 %d", expected, result)
	}

	// 测试用例 2
	result = Add(0, 0)
	expected = 0
	if result != expected {
		t.Errorf("Add(0, 0) 返回值错误，期望 %d，实际 %d", expected, result)
	}

	// 测试用例 3
	result = Add(-1, 1)
	expected = 0
	if result != expected {
		t.Errorf("Add(-1, 1) 返回值错误，期望 %d，实际 %d", expected, result)
	}
}