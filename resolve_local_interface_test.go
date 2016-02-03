package netlib

import (
	"reflect"
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
		res0, err := ResolveLocalInterface("")
		if err != nil {
			t.Errorf("Expected automatic interface resolution to succeed but got err=%T/%s", err, err)
		}

		res1, err := ResolveLocalInterface("0.0.0.0")
		if err != nil {
			t.Errorf("Expected automatic interface resolution to succeed but got err=%T/%s", err, err)
		}

		if !reflect.DeepEqual(res0, res1) {
			t.Errorf("Expected res0(%v) == res1(%v), but they differed", res0, res1)
		}

		res2, err := ResolveLocalInterface("0:0:0:0:0:0:0:0")
		if err != nil {
			t.Errorf("Expected automatic interface resolution to succeed but got err=%T/%s", err, err)
		}

		if !reflect.DeepEqual(res1, res2) {
			t.Errorf("Expected res1(%v) == res2(%v), but they differed", res1, res2)
		}
	}
}
