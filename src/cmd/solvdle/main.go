package main

import (
	"bufio"
	"log"
	"os"

	"github.com/0x4445565a/solvdle/src/pkg/node"
)

func main() {
	root := &node.Node{}

	file, err := os.Open("words_len5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		root.Insert(scanner.Text())
	}

	log.Println(root.MatchPattern("c_imp", map[rune]bool{
		'a': true,
		'd': true,
		'e': true,
		'u': true,
		'h': true,
	}, map[rune]int{}))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
