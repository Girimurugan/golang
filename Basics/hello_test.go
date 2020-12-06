package main

import "testing"

func TestSum(t *testing.T) {

	result := sum(10, 20)

	if result == 30 {
		t.Logf("Correct")
	} else {
		t.Errorf("Error")
	}
}
