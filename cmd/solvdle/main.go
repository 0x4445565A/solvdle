package main

import (
	"bufio"
	"log"
	"os"
)

const MAX_LEVEL = 5
const WILDCARD = '_'

var RuneToIndex = map[rune]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
	'i': 8,
	'j': 9,
	'k': 10,
	'l': 11,
	'm': 12,
	'n': 13,
	'o': 14,
	'p': 15,
	'q': 16,
	'r': 17,
	's': 18,
	't': 19,
	'u': 20,
	'v': 21,
	'w': 22,
	'x': 23,
	'y': 24,
	'z': 25,
}

type Node struct {
	Value rune
	Level int
	// We store a-z nodes
	Children [26]*Node
}

func (n *Node) FindWord(s string) bool {
	position := n
	for _, r := range s {
		position = position.Children[RuneToIndex[r]]
		if position == nil {
			return false
		}
	}

	return true
}

func (n *Node) MatchPattern(pattern string, banned map[rune]bool, needed map[rune]int) []string {
	var ret []string

	if pattern == "" {

		// Since we have letters that we need still but are at the ened of the road, skip em
		if len(needed) > 0 {
			return []string{}
		}

		return []string{string(n.Value)}
	}

	c, s := removeFirstChar(pattern)

	var validChildren []*Node

	if c == WILDCARD {
		validChildren = n.Children[:]
	} else {
		validChildren = []*Node{
			n.Children[RuneToIndex[c]],
		}
	}

	force := false
	if MAX_LEVEL-n.Level <= len(needed) {
		force = true
	}

	for _, nn := range validChildren {
		if nn != nil {
			// in our banned list meaning we should not search
			if _, ok := banned[nn.Value]; ok && c == WILDCARD {
				continue
			}

			if lvl, ok := needed[nn.Value]; ok || !force {
				// We need this character but not on this level
				if lvl == nn.Level {
					continue
				}

				// Maps pass by reference so we need to remake it.  Small enough to be fine.
				needed_copy := needed
				if ok {
					needed_copy = map[rune]int{}
					for k, level := range needed {
						if k == nn.Value {
							continue
						}
						needed_copy[k] = level
					}
				}

				// We're going down down in an earlier round, and sugar we're going down swinging
				//                                                             --  Aristotle
				ret = append(ret, nn.MatchPattern(s, banned, needed_copy)...)
			}
		}
	}

	for i, s := range ret {
		ret[i] = string(n.Value) + s
	}
	return ret
}

func removeFirstChar(s string) (rune, string) {
	if s == "" {
		return 0, ""
	}

	c := rune(s[0])

	if len(s) <= 1 {
		return c, ""
	}

	return c, s[1:]
}

func main() {
	root := &Node{}

	file, err := os.Open("words_len5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		position := root
		for i, c := range scanner.Text() {
			if position.Children[RuneToIndex[c]] == nil {

				position.Children[RuneToIndex[c]] = &Node{
					Value: c,
					Level: i + 1,
				}
			}

			position = position.Children[RuneToIndex[c]]
		}
	}

	log.Println(root.MatchPattern("_i_c_", map[rune]bool{
		'm': true,
		'c': true,
		'p': true,
		'y': true,
		's': true,
		'o': true,
		'a': true,
		'b': true,
		't': true,
		'h': true,
		'g': true,
	}, map[rune]int{
		'i': 4,
		'c': 5,
	}))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
