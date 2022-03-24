package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

const (
	space      = ' '
	defaultTop = 10
)

type count struct {
	Word  string
	Count int
}

func main() {
	// words := scanWords(os.Stdin)
	words := countWords(os.Stdin)

	var counts []count
	for w, c := range words {
		counts = append(counts, count{
			Word:  w,
			Count: c,
		})
	}

	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count
	})

	length := len(counts)
	if length > defaultTop {
		length = defaultTop
	}

	for _, w := range counts[:length] {
		fmt.Printf("%d %s\n", w.Count, w.Word)
	}
}

func scanWords(rd io.Reader) map[string]int {
	words := make(map[string]int)
	scanner := bufio.NewScanner(rd)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words[scanner.Text()]++
	}
	return words
}

func countWords(rd io.Reader) map[string]int {
	var word []byte
	buf := make([]byte, bufio.MaxScanTokenSize)
	words := make(map[string]int)

	for {
		n, err := rd.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			c := buf[i]
			if c <= space {
				if len(word) > 0 {
					words[string(word)]++
					word = word[:0]
				}
				continue
			}

			word = append(word, c)
		}
	}

	if len(word) > 0 {
		words[string(word)]++
	}

	return words
}
