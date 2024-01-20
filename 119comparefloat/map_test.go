package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// map，没毛病。
func mapParam_04(m map[string]string) {
	m["a"] = "aa"
	m["c"] = "c"

}

func Test_mapParam_04(t *testing.T) {
	t.Run("tt.name", func(t *testing.T) {
		m := map[string]string{"a": "a", "b": "b"}
		mapParam_04(m)
		assert.Equal(t, map[string]string{"a": "aa", "b": "b", "c": "c"}, m)
	})

}
