package project_x

import (
	"gx/ipfs/QmPVkJMTeRC6iBByPWdrRkD3BE5UXsj5HPzb4kPqL186mS/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t,LengthOfLongestSubstring("abcabcbb"),3)
	assert.Equal(t,LengthOfLongestSubstring("bbbbb"),1)
	assert.Equal(t,LengthOfLongestSubstring("pwwkew"),3)
}
