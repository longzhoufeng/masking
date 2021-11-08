package masking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToStar(t *testing.T) {
	tests := []struct {
		input  interface{}
		preNum int
		sufNum int
		expect string
	}{
		{12345678, 4, 2, "1234**78"},
		{12345678, 4, 4, "12345678"},
		{12345678, 4, 5, "12345678"},
		{12345678, 9, 5, "12345678"},
		{12345678, 0, 0, "********"},
		{12345678, 0, 3, "*****678"},
		{12345678, 3, 0, "123*****"},
		{12345678, -2, 0, "********"},
		{12345678, 0, -3, "********"},
		{12345678, -2, -3, "********"},
		{123456.78, 4, 2, "1234***78"},
		{-123456.78, 4, 3, "-123***.78"},
		{"12345678", 4, 2, "1234**78"},
		{[]byte("123 45678"), 4, 2, "123 ***78"},
		{true, 0, 0, "****"},
		{false, 0, 0, "*****"},
		{nil, 0, 5, ""},
	}

	for _, test := range tests {
		errmsg := fmt.Sprintf("test = %v", test)
		v := Star(test.input, test.preNum, test.sufNum)
		assert.Equal(t, test.expect, v, errmsg)
	}
}
