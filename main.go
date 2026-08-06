package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
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

func main() {
	flag.Parse()
	var lines []string
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	shuffle(lines)
	for _, l := range lines {
		fmt.Println(l)
	}
}
