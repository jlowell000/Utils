package utils

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	size = 1000
)

func Test_ForSquare(t *testing.T) {
	expected := [size][size]int{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			expected[i][j] = i * j
		}
	}

	actual := [size][size]int{}
	ForSquare(size, func(i, j int) {
		actual[i][j] = i * j
	})
	assert.Equal(t, expected, actual, "For^2 Loop matches ForSquare")
}

func Test_For(t *testing.T) {
	expected := []int{}
	for i := 0; i < size; i++ {
		expected = append(expected, i)
	}

	actual := []int{}
	For(size, func(i int) {
		actual = append(actual, i)
	})
	assert.Equal(t, expected, actual, "For Loop matches For")
}

func Test_ForEach(t *testing.T) {
	testData := []int{}
	For(size, func(i int) { testData = append(testData, i) })

	expected := []int{}
	for _, t := range testData {
		expected = append(expected, t*t)
	}

	actual := []int{}
	ForEach(testData, func(t int) {
		actual = append(actual, t*t)
	})
	assert.Equal(t, expected, actual, "For Loop matches For")
}

func Test_Map(t *testing.T) {
	testData := make([]int, size)
	For(size, func(i int) { testData[i] = i })
	expected := make([]int, size)
	for i, t := range testData {
		expected[i] = t * t
	}

	actual := Map(testData, func(t int) int {
		return t * t
	})
	assert.Equal(t, expected, actual, "Map func matches expected")
}

func Test_Filter(t *testing.T) {
	testData := make([]int, size)
	For(size, func(i int) { testData[i] = i })
	expected := []int{}
	for _, t := range testData {
		if t%2 == 0 {
			expected = append(expected, t)
		}
	}

	actual := Filter(testData, func(t int) bool {
		return t%2 == 0
	})
	assert.Equal(t, expected, actual, "Filter func matches expected")
}

func Test_ForEachWG(t *testing.T) {
	testData := []int{}
	For(size, func(i int) { testData = append(testData, i) })

	expected := []int{}
	for _, t := range testData {
		expected = append(expected, t*t)
	}

	actual := make([]int, len(testData))
	ForEachWG(testData, func(t int) {
		actual[t] = t * t
	})

	assert.Equal(t, len(expected), len(actual), "Arrays not same size")
	for _, e := range expected {
		assert.True(t, slices.Contains(actual, e), fmt.Sprintf("actual does not contain expected: %d", e))
	}
}

func Test_MapWG(t *testing.T) {
	testData := []int{}
	For(size, func(i int) { testData = append(testData, i) })

	expected := []int{}
	for _, t := range testData {
		expected = append(expected, t*t)
	}

	actual := MapWG(testData, func(t int) int {
		return t * t
	})

	assert.Equal(t, len(expected), len(actual), "Arrays not same size")
	for _, e := range expected {
		assert.True(t, slices.Contains(actual, e), fmt.Sprintf("actual does not contain expected: %d", e))
	}
}
func Test_OrderedMapWG(t *testing.T) {
	testData := make([]int, size)
	For(size, func(i int) { testData[i] = i })
	expected := make([]int, size)
	for i, t := range testData {
		expected[i] = t * t
	}

	actual := OrderedMapWG(testData, func(t int) int {
		return t * t
	})
	assert.Equal(t, expected, actual, "Mapping working group match sequential mapping")
}

func Test_FilterWG(t *testing.T) {
	testData := make([]int, size)
	For(size, func(i int) { testData[i] = i })
	expected := []int{}
	for _, t := range testData {
		if t%2 == 0 {
			expected = append(expected, t)
		}
	}

	actual := FilterWG(testData, func(t int) bool {
		return t%2 == 0
	})
	assert.Equal(t, len(expected), len(actual), "Arrays not same size")
	for _, e := range expected {
		assert.True(t, slices.Contains(actual, e), fmt.Sprintf("actual does not contain expected: %d", e))
	}
}

func Test_ActionWG(t *testing.T) {
	testData := []int{}
	For(size, func(i int) { testData = append(testData, i) })

	expected := []int{}
	for _, t := range testData {
		expected = append(expected, t*t)
	}
	actual := make([]int, len(testData))
	actions := []func(){}
	For(size, func(i int) { actions = append(actions, func() { actual[i] = i * i }) })
	ActionWG(actions)

	assert.Equal(t, len(expected), len(actual), "Arrays not same size")
	for _, e := range expected {
		assert.True(t, slices.Contains(actual, e), fmt.Sprintf("actual does not contain expected: %d", e))
	}
}
