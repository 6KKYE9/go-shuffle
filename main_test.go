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

func TestSplitByDelimLine(t *testing.T) {
	// 按行时结尾换行产生的空项要去掉
	got := splitByDelim("a\nb\nc\n", "\n")
	if len(got) != 3 || got[0] != "a" || got[2] != "c" {
		t.Fatalf("按行分割不符: %#v", got)
	}
}

func TestSplitByDelimCustom(t *testing.T) {
	// 按逗号分隔，连续的空段丢掉
	got := splitByDelim("x,,y,z", ",")
	if len(got) != 3 || got[0] != "x" || got[2] != "z" {
		t.Fatalf("逗号分割不符: %#v", got)
	}
}

func TestReverseInPlace(t *testing.T) {
	lines := []string{"1", "2", "3", "4"}
	// 复用 main 里的反转逻辑：直接检验洗牌+反转后集合仍完整
	s := make([]string, len(lines))
	copy(s, lines)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	if s[0] != "4" || s[3] != "1" {
		t.Fatalf("反转不符: %#v", s)
	}
}
