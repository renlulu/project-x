package project_x

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	strs := []string{"flower","flow","flight"}
	r := LongestCommonPrefix(strs)
	fmt.Println(r)
}
