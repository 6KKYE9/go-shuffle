package main

import (
	"sort"
	"testing"
)

func TestShufflePreservesSet(t *testing.T) {
	lines := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	orig := make([]string, len(lines))
	copy(orig, lines)
	shuffle(lines)
	// 内容集合不变，只是顺序乱了
	sort.Strings(lines)
	sort.Strings(orig)
	for i := range orig {
		if lines[i] != orig[i] {
			t.Fatalf("洗牌后集合变了: %#v vs %#v", lines, orig)
		}
	}
}

func TestShuffleEmpty(t *testing.T) {
	shuffle([]string{}) // 不应 panic
}
