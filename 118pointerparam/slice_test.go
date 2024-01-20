package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 用指针传递，没毛病。
func sliceParam_04(s *[]string) {
	(*s)[0] = "zzz"
	*s = append(*s, "444", "444", "444", "444")
}

func Test_sliceParam_04(t *testing.T) {
	t.Run("tt.name", func(t *testing.T) {
		s := []string{"111", "222", "333"}
		sliceParam_04(&s)
		// 用指针传递，没毛病。追加，换了底层数组了
		assert.Equal(t, []string{"zzz", "222", "333", "444", "444", "444", "444"}, s)
	})

}

// 不定长变量。
func sliceParam_05(s ...string) {
	s[0] = "zzz"
	s = append(s, "444", "444", "444", "444")
}

func Test_sliceParam_05(t *testing.T) {
	t.Run("tt.name", func(t *testing.T) {
		s := []string{"111", "222", "333"}
		sliceParam_05(s...)
		// 追加，换了底层数组了
		assert.Equal(t, []string{"zzz", "222", "333"}, s)
	})
}

func sliceParam_01(s []string) {
	s[0] = "zzz"
}

func Test_sliceParam_01(t *testing.T) {
	t.Run("tt.name", func(t *testing.T) {
		s := []string{"111", "222", "333"}
		sliceParam_01(s)
		assert.Equal(t, []string{"zzz", "222", "333"}, s)
	})
}

func sliceParam_02(s []string) {
	s[0] = "zzz"
	s = append(s, "444", "444", "444", "444")
}

func Test_sliceParam_02(t *testing.T) {
	t.Run("tt.name", func(t *testing.T) {
		s := []string{"111", "222", "333"}
		sliceParam_02(s)
		// 追加，换了底层数组了
		assert.Equal(t, []string{"zzz", "222", "333"}, s)
	})
	t.Run("tt.name", func(t *testing.T) {
		s := []string{}
		s = append(s, "111", "222", "333")
		sliceParam_02(s)
		// 追加，换了底层数组了
		assert.Equal(t, []string{"zzz", "222", "333"}, s)
	})
}

func sliceParam_03(s []string) {
	s[0] = "zzz"
	s = append(s, "444")
}

func Test_sliceParam_03(t *testing.T) {
	t.Run("tt.name", func(t *testing.T) {
		s := []string{"111", "111", "111"}
		sliceParam_03(s)
		// 追加，换了底层数组了
		assert.Equal(t, []string{"zzz", "111", "111"}, s)
	})
	t.Run("tt.name", func(t *testing.T) {
		s := []string{}
		s = append(s, "111")
		s = append(s, "111")
		s = append(s, "111")
		sliceParam_03(s)
		// 追加，没有 换底层数组
		assert.Equal(t, []string{"zzz", "111", "111"}, s)
	})
}
