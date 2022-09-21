package array

import (
	"github.com/samber/lo"
	"strconv"
	"strings"
)

func Int32ArrayToIntArray(a []int32) []int {
	return lo.Map(a, func(x int32, _ int) int {
		return int(x)
	})
}

func InArrayString(search string, slice []string) bool {
	return inArray(search, slice)
}

func InArrayInt(search int, slice []int) bool {
	return inArray(search, slice)
}

func InArrayInt64(search int64, slice []int64) bool {
	return inArray(search, slice)
}

func InArrayInt32(search int32, slice []int32) bool {
	return inArray(search, slice)
}

func IntMap(slice []int) map[int]bool {
	rs := map[int]bool{}
	for _, it := range slice {
		rs[it] = true
	}
	return rs
}

func inArray(needle interface{}, haystacks interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range haystacks.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range haystacks.([]int) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range haystacks.([]int64) {
			if key == item {
				return true
			}
		}
	case int32:
		for _, item := range haystacks.([]int32) {
			if key == item {
				return true
			}
		}
	default:
		panic("unsupported type")
	}
	return false
}

func DistinctIntArray(array []int) []int {
	m := map[int]struct{}{}
	rs := make([]int, 0)
	for _, i := range array {
		if _, ok := m[i]; ok {
			continue
		}
		m[i] = struct{}{}
		rs = append(rs, i)
	}
	return rs
}

func DistinctInt64Array(arr []int64) []int64 {
	m := map[int64]struct{}{}
	rs := make([]int64, 0)
	for _, i := range arr {
		if _, ok := m[i]; ok {
			continue
		}
		m[i] = struct{}{}
		rs = append(rs, i)
	}
	return rs
}

func DistinctInt32Array(arr []int32) []int32 {
	m := map[int32]struct{}{}
	rs := make([]int32, 0)
	for _, i := range arr {
		if _, ok := m[i]; ok {
			continue
		}
		m[i] = struct{}{}
		rs = append(rs, i)
	}
	return rs
}

func DistinctStringArray(arr []string) []string {
	m := map[string]struct{}{}
	var rs []string
	for _, s := range arr {
		if _, ok := m[s]; ok {
			continue
		}
		m[s] = struct{}{}
		rs = append(rs, s)
	}
	return rs
}

func Int64ToIntArr(input []int64) []int {
	var rs []int
	for _, it := range input {
		rs = append(rs, int(it))
	}
	return rs
}

func Int32ToIntArr(input []int32) []int {
	var rs []int
	for _, it := range input {
		rs = append(rs, int(it))
	}
	return rs
}

// ToInt32Array WARNING: 精度丢失
func ToInt32Array(input []int) []int32 {
	var rs []int32
	for _, it := range input {
		rs = append(rs, int32(it))
	}
	return rs
}

func JoinIntArr(arr []int, separator string) string {
	var a []string
	for _, it := range arr {
		a = append(a, strconv.Itoa(it))
	}
	return strings.Join(a, separator)
}
