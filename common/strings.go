package common

import (
	"strconv"
	"strings"
)

// Split returns a []int by splitting str by sep .
func Split(str string, sep string) ([]int, error) {
	strs := strings.Split(str, sep)
	ints := make([]int, len(strs))
	var err error
	for i := range strs {
		ints[i], err = strconv.Atoi(strs[i])
		if err != nil {
			return []int{}, err
		}
	}
	return ints, nil
}
