package utils

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestInt2String(t *testing.T) {
	var num = int64(35)
	var result = Int2String(num, 5)
	assert.Equal(t, result, "aaaaJ")
	num = 125
	result = Int2String(num, 5)
	assert.Equal(t, result, "aaacb")
}

func TestString2Int(t *testing.T) {
	var str = "J"
	var result = String2Int(str)
	assert.Equal(t, result, int64(35))
	str = "cb"
	result = String2Int(str)
	assert.Equal(t, result, int64(125))
}
