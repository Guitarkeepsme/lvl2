package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask2(t *testing.T) {
	// Arrange
	testTable := []struct {
		str      PackedString
		expected string
	}{
		{
			str:      "st3k",
			expected: "stttk",
		},
		{
			str:      "s\\2as",
			expected: "s2as",
		},

		{
			str:      "d7kj10i",
			expected: "dddddddkjjjjjjjjjji",
		},

		{
			str:      "d7kj10",
			expected: "dddddddkjjjjjjjjjj",
		},
		{
			str:      "45",
			expected: "",
		},

		{
			str:      "\\\\5ka2lv",
			expected: "\\\\\\\\\\kaalv",
		},

		{
			str:      "\\\\5ka2lv",
			expected: "\\\\\\\\\\kaalv",
		},

		{
			str:      "\\\\5k11i",
			expected: "\\\\\\\\\\kkkkkkkkkkki",
		},
	}

	// Act
	for _, testCase := range testTable {
		res := testCase.str.Unpack()

		t.Logf("Calling Unpacking(%v), result %v\n", testCase.str, res)

		//Assert

		assert.Equal(t, testCase.expected, res,
			fmt.Sprintf("Incorrect result. Expected %v, got %v", testCase.expected, res))
	}

}
