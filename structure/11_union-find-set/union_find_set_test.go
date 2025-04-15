package unionfindset

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(10)

	// 测试初始状态
	if uf.Count() != 10 {
		t.Errorf("Expected count 10, got %d", uf.Count())
	}

	// 测试合并
	uf.Union(4, 3)
	uf.Union(3, 8)
	uf.Union(6, 5)
	uf.Union(9, 4)
	uf.Union(2, 1)

	// 测试连通性
	if !uf.Connected(8, 9) {
		t.Error("Expected 8 and 9 to be connected")
	}

	if uf.Connected(5, 4) {
		t.Error("Expected 5 and 4 to be not connected")
	}

	// 测试集合数量
	if uf.Count() != 5 {
		t.Errorf("Expected count 5, got %d", uf.Count())
	}
}
