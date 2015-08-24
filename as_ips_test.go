package netlib

import (
	"fmt"
	"testing"
)

func Test_AsIps(t *testing.T) {
	actual, err := AsIps("localhost", "localhost:5144", "localhost:21")
	if err != nil {
		t.Fatalf("asIps failed: %s", err)
	}
	expected := []string{"127.0.0.1", "127.0.0.1:5144", "127.0.0.1:21"}
	if fmt.Sprintf("%+v", actual) != fmt.Sprintf("%+v", expected) {
		t.Fatalf("Expected output slice=%+v but instead got %+v", expected, actual)
	}
}
