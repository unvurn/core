package reslices_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/unvurn/core/collections/reslices"
)

func TestContainsAnyTrue(t *testing.T) {
	s := []int{1, 2, 3}

	assert.True(t, reslices.ContainsAny(s, []int{1}))
	assert.True(t, reslices.ContainsAny(s, []int{1, 2}))
	assert.True(t, reslices.ContainsAny(s, []int{1, 2, 3}))
	assert.True(t, reslices.ContainsAny(s, []int{1, 2, 3, 4}))
}

func TestContainsAnyFalse(t *testing.T) {
	s := []int{1, 2, 3}

	assert.False(t, reslices.ContainsAny(s, []int{4}))
	assert.False(t, reslices.ContainsAny(s, []int{4, 5}))
	assert.False(t, reslices.ContainsAny(s, []int{0, 4, 5}))
}

func TestContainsAnyNil(t *testing.T) {
	var s []int
	s1 := []int{1, 2, 3}

	assert.False(t, reslices.ContainsAny(s, s1))

	s2 := []int{1, 2, 3}

	assert.False(t, reslices.ContainsAny(s2, s))
}

func TestCollectionsContainsAnyInt(t *testing.T) {
	cases := []struct {
		array    []int
		search   []int
		expected bool
	}{
		{[]int{}, []int{1, 2, 3}, false},
		{[]int{11}, []int{1, 2, 3}, false},
		{[]int{11, 12, 13}, []int{1, 2, 3}, false},
		{[]int{11, 12, 13}, []int{1}, false},
		{[]int{11, 12, 13}, []int{}, false},
		{[]int{2}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{2}, true},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
	}
	for _, tc := range cases {
		contains := reslices.ContainsAny(tc.array, tc.search)
		assert.Equal(t, tc.expected, contains)
	}
}

func TestCollectionsContainsAnyString(t *testing.T) {
	cases := []struct {
		array    []string
		search   []string
		expected bool
	}{
		{[]string{}, []string{"1", "2", "3"}, false},
		{[]string{"11"}, []string{"1", "2", "3"}, false},
		{[]string{"11", "12", "13"}, []string{"1", "2", "3"}, false},
		{[]string{"11", "12", "13"}, []string{"1"}, false},
		{[]string{"11", "12", "13"}, []string{}, false},
		{[]string{"2"}, []string{"1", "2", "3"}, true},
		{[]string{"1", "2", "3"}, []string{"2"}, true},
		{[]string{"1", "2", "3"}, []string{"1", "2", "3"}, true},
	}
	for _, tc := range cases {
		contains := reslices.ContainsAny(tc.array, tc.search)
		assert.Equal(t, tc.expected, contains)
	}
}
