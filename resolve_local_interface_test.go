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
			t.Errorf("Expected ResolveLocalInterface(%q) = %q but actual = %q", expected, expected, actual)
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

		res3, err := ResolveLocalInterface("[::]")
		if err != nil {
			t.Errorf("Expected automatic interface resolution to succeed but got err=%T/%s", err, err)
		}

		if !reflect.DeepEqual(res1, res2) {
			t.Errorf("Expected res1(%v) == res2(%v), but they differed", res1, res2)
		}

		if !reflect.DeepEqual(res1, res3) {
			t.Errorf("Expected res1(%v) == res3(%v), but they differed", res1, res3)
		}

		res4, err := ResolveLocalInterface("", res1.String())
		// NB: An error here may not mean anything; for example if there was only 1
		// interface from res1 then none will be found.
		if err == nil {
			if reflect.DeepEqual(res1, res4) {
				t.Errorf("Expected res4=%v to be different from res1=%v", res4, res1)
			}
		}
	}
}

func TestIsEmptyBindSpec(t *testing.T) {
	testCases := []struct {
		bind        string
		expectEmpty bool
	}{
		{
			bind:        "0.0.0.0",
			expectEmpty: true,
		},
		{
			bind:        "[::]",
			expectEmpty: true,
		},
		{
			bind:        "0:0:0:0:0:0:0:0",
			expectEmpty: true,
		},
		{
			bind:        "127.0.0.1",
			expectEmpty: false,
		},
	}

	for i, testCase := range testCases {
		if expected, actual := testCase.expectEmpty, IsEmptyBindSpec(testCase.bind); actual != expected {
			t.Errorf("[i=%v/testCase=%+v] Expected bind=%v empty=%v but actual=%v", i, testCase, testCase.bind, expected, actual)
		}
	}
}
