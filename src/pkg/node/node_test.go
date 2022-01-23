package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeInsert(t *testing.T) {
	root := &Node{}

	root.Insert("foooo")
	root.Insert("baaar")
	root.Insert("baaaz")

	assert.Equal(t, true, root.FindWord(""), "Should match")
	assert.Equal(t, true, root.FindWord("foooo"), "Should match")
	assert.Equal(t, true, root.FindWord("baaar"), "Should match")
	assert.Equal(t, true, root.FindWord("baaaz"), "Should match")

	assert.Equal(t, false, root.FindWord("zzzzz"), "Should match")
}

func TestNodeMatchPattern(t *testing.T) {
	root := &Node{}

	root.Insert("foooo")
	root.Insert("faaaa")
	root.Insert("ffaaa")
	root.Insert("ffaaf")
	root.Insert("baaar")
	root.Insert("baaaz")

	assert.Equal(t, []string{"baaar", "baaaz", "faaaa", "ffaaa", "ffaaf", "foooo"}, root.MatchPattern("_____", nil, nil, nil), "Should match")
	assert.Equal(t, []string{"baaar", "baaaz", "faaaa"}, root.MatchPattern("_a___", nil, nil, nil), "Should match")
	assert.Equal(t, []string{"baaar", "baaaz"}, root.MatchPattern("ba___", nil, nil, nil), "Should match")
	assert.Equal(t, []string{"baaar", "baaaz"}, root.MatchPattern("baaa_", nil, nil, nil), "Should match")
	assert.Equal(t, []string{"baaaz"}, root.MatchPattern("baaa_", map[rune]bool{'r': true}, nil, nil), "Should match")

	// We have the first F and need another but it can't be in the 5th position
	assert.Equal(t, []string{"ffaaa"}, root.MatchPattern("f____", nil, map[rune]int{'f': 5}, map[rune]bool{'f': true}), "Should match")
}
