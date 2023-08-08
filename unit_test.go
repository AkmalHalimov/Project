package main

import (
	"testing"
)

func GetTest(t *testing.T) {
	result := GetSum(1, 9)
	if result != 7 {
		panic(result)
	}
}
