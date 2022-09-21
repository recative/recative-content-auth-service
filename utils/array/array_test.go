package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistinctIntArray(t *testing.T) {
	i := []int{1, 2, 3, 1}
	assert.Equal(t, []int{1, 2, 3}, DistinctIntArray(i))
}

func TestDistinctStringArray(t *testing.T) {
	i := []string{"a", "b", "aabb", "b"}
	assert.Equal(t, []string{"a", "b", "aabb"}, DistinctStringArray(i))
}

func TestJoinIntArrF(t *testing.T) {
	assert.Equal(t, "1,2,3", JoinIntArr([]int{1, 2, 3}, ","))
}
