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

func TestShuffleSeededReproducible(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	b := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	shuffleSeeded(a, 12345)
	shuffleSeeded(b, 12345)
	for i := range a {
		if a[i] != b[i] {
			t.Fatalf("同种子应得到相同顺序: %#v vs %#v", a, b)
		}
	}
}

func TestShuffleSeededDifferent(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5"}
	b := []string{"1", "2", "3", "4", "5"}
	shuffleSeeded(a, 1)
	shuffleSeeded(b, 2)
	// 不同种子大概率不同，这里只验证都能跑完不 panic
	_ = a
	_ = b
}
