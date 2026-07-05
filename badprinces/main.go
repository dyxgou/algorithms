package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 5
	// 6
	// 3 9 4 6 7 5
	// 1
	// 1000000
	// 2
	// 2 1
	// 10
	// 31 41 59 26 53 58 97 93 23 84
	// 7
	// 3 2 1 2 3 4 5

	s := bufio.NewScanner(os.Stdin)

	var bad int
	for s.Scan() {
		line := s.Text()

		prices := strings.Split(line, " ")
		if len(prices) <= 1 {
			continue
		}

		prevPrice, err := strconv.Atoi(prices[0])
		if err != nil {
			panic(err.Error())
		}

		for p := range pricesToInt(prices) {
			if prevPrice < p {
				// slog.Info("bad prices", "prev", prevPrice, "cur", p)
				bad++
			}

			prevPrice = p
		}

		fmt.Println(bad)
		bad = 0
	}

	if err := s.Err(); err != nil {
		panic(err)
	}
}

func pricesToInt(prices []string) iter.Seq[int] {
	return func(yield func(int) bool) {
		for _, p := range prices {
			n, err := strconv.Atoi(p)

			if err != nil {
				panic(err)
			}

			if !yield(n) {
				return
			}
		}
	}
}
