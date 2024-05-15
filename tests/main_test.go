package main

import (
	"os"
	"testing"
)

func TestRun1(t *testing.T) {
	result1 := Run("tests/input/in1.txt")
	expect1, _ := os.ReadFile("tests/output/out1.txt")
	if string(expect1) != result1 {
		t.Errorf("Incorrect result. Expect %s, got %s", expect1, result1)
	}
}

func TestRun2(t *testing.T) {
	result2 := Run("tests/input/in2.txt")
	expect2, _ := os.ReadFile("tests/output/out2.txt")
	if string(expect2) != result2 {
		t.Errorf("Incorrect result. Expect %s, got %s", expect2, result2)
	}
}

func TestRun3(t *testing.T) {
	result3 := Run("tests/input/in3.txt")
	expect3, _ := os.ReadFile("tests/output/out3.txt")
	if string(expect3) != result3 {
		t.Errorf("Incorrect result. Expect %s, got %s", expect3, result3)
	}
}

func TestRun4(t *testing.T) {
	result4 := Run("tests/input/in4.txt")
	expect4, _ := os.ReadFile("tests/output/out4.txt")
	if string(expect4) != result4 {
		t.Errorf("Incorrect result. Expect %s, got %s", expect4, result4)
	}
}

func TestRun5(t *testing.T) {
	result5 := Run("tests/input/in5.txt")
	expect5, _ := os.ReadFile("tests/output/out5.txt")
	if string(expect5) != result5 {
		t.Errorf("Incorrect result. Expect %s, got %s", expect5, result5)
	}
}

func TestRun6(t *testing.T) {
	result6 := Run("tests/input/in6.txt")
	expect6, _ := os.ReadFile("tests/output/out6.txt")
	if string(expect6) != result6 {
		t.Errorf("Incorrect result. Expect %s, got %s", expect6, result6)
	}
}
