package netlib

import (
	"testing"
)

func TestResolveLocalInterface(t *testing.T) {
	{
		expected := "127.0.0.1"
		ip, err := ResolveLocalInterface(expected)
		if err != nil {
			t.Fatal(err)
		} else if actual := ip.String(); actual != expected {
			t.Errorf("Expected ResolveLocalInterface(%q) = %q but actual = %q", expected, actual)
		}
	}
	{
		if _, err := ResolveLocalInterface(""); err != nil {
			t.Errorf("Expected automatic interface resolution to succeed but got err=%T/%s", err, err)
		}
	}
}
