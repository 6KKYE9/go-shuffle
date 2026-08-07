package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"math/big"
	mrand "math/rand"
	"os"
	"strconv"
	"strings"
)

// 用 crypto/rand 做洗牌，比 math/rand 更适合不想被预测的场景
func shuffle(lines []string) {
	for i := len(lines) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		k := i
		if err == nil {
			k = int(j.Int64())
		}
		// crypto/rand 极少出错，真出了就原地不动，保证每轮都走到
		lines[i], lines[k] = lines[k], lines[i]
	}
}

// 用固定种子洗牌，结果可复现（同一 seed 出同一顺序），方便测试和复现
func shuffleSeeded(lines []string, seed int64) {
	r := mrand.New(mrand.NewSource(seed))
	r.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})
}

func main() {
	num := flag.Int("n", 0, "只输出前 N 行（0 表示全部）")
	seed := flag.String("seed", "", "用固定种子洗牌，结果可复现（比如 -seed 123）")
	reverse := flag.Bool("reverse", false, "洗牌后把顺序反过来（相当于倒序输出）")
	delim := flag.String("delim", "\n", "按什么分隔输入，默认按行；可设空格、逗号等")
	outPath := flag.String("o", "", "结果写到这个文件，不给就打到标准输出")
	flag.Parse()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "读入失败: %v\n", err)
		os.Exit(1)
	}
	lines := splitByDelim(string(input), *delim)

	if *seed != "" {
		s, err := strconv.ParseInt(*seed, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "seed 必须是数字: %v\n", err)
			os.Exit(1)
		}
		shuffleSeeded(lines, s)
	} else {
		shuffle(lines)
	}
	if *reverse {
		// 原地反转，简单直接
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}
	if *num > 0 && *num < len(lines) {
		lines = lines[:*num]
	}

	var out strings.Builder
	for _, l := range lines {
		if *delim == "\n" {
			out.WriteString(l)
			out.WriteByte('\n')
		} else {
			// 非按行时，用原分隔符拼回
			if out.Len() > 0 {
				out.WriteString(*delim)
			}
			out.WriteString(l)
		}
	}
	if *delim != "\n" {
		out.WriteByte('\n')
	}

	if *outPath != "" {
		if err := os.WriteFile(*outPath, []byte(out.String()), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "写文件失败: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("已写出 %s\n", *outPath)
		return
	}
	fmt.Print(out.String())
}

// splitByDelim 按分隔符切分，按行时保留每行的实际内容（去掉空的尾随项）
func splitByDelim(s, delim string) []string {
	if delim == "\n" {
		// 按行：去掉结尾的换行产生的空项，但中间的空行保留
		lines := strings.Split(s, "\n")
		if len(lines) > 0 && lines[len(lines)-1] == "" {
			lines = lines[:len(lines)-1]
		}
		return lines
	}
	parts := strings.Split(s, delim)
	var out []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		out = append(out, p)
	}
	return out
}
