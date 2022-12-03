package intset

import (
	"bytes"
	"fmt"
)

// An intset is a set of small non-negative integers.
// Ints zero value represents the empty set
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x
// check if there is the word containing x AND
// use bitwise & to check if x is in the word (1 & 1 == 1) (1 & 0 == 0)
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
// divide x by n of bits to get what word will the number be stored
// modulo by n bits to get the bit position of the word were in
// |= 1 << bit  toggles the bith at bit position.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Len() int {
	res := 0
	for _, word := range s.words {
		res += PopCountK(word)
	}
	return res
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &^= (1 << bit) // flip the bit with XOR
}

func (s *IntSet) Clear() {
	*s = IntSet{}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	new := &IntSet{}
	new.words = make([]uint64, len(s.words))
	copy(new.words, s.words)
	return new
}

// String returns the set as a string formatted "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Values() []byte {
	var buf []byte

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if len(buf) > 0 {
					buf = append(buf, ' ')
				}
				buf = append(buf, byte(64*i+j))
				// fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}

	return buf
}
