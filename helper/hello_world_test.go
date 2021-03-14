package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ega")

	if result != "Hello Ega" {
		t.Fail()
	}

	fmt.Println("Test hello ")
}

func TestHelloWorldXixi(t *testing.T) {
	result := HelloWorld("Xixi")

	if result != "Hello Xixi" {
		t.FailNow()
	}

	fmt.Println("Test Xixi")
}
