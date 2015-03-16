package jchash

import "testing"

func TestJCHash(t *testing.T) {
	cases := []string{
		"This is a test strings for murmur3 hash algorithm",
		"Hello, 世界",
		"",
	}
	for _, c := range cases {
		var n uint64 = 1
		for ; n < 1024*1024; n++ {
			_, err := JumpConsistentHash(c, n)
			if err != nil {
				t.Errorf("JumpConsistentHash(%q), error %q, n value %d", c, err, n)
			}
		}
	}
}
