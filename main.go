package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	mrand "math/rand"
	"os"
	"strconv"
)

// 用 crypto/rand 做洗牌，比 math/rand 更适合不想被预测的场景
func shuffle(lines []string) {
	for i := len(lines) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			// 真出概率极低，退回不洗这一轮
			continue
		}
		k := int(j.Int64())
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
	flag.Parse()

	var lines []string
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
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
	if *num > 0 && *num < len(lines) {
		lines = lines[:*num]
	}
	for _, l := range lines {
		fmt.Println(l)
	}
}
