package main

import (
	"testing"

	"github.com/clsung/fake"
)

func TestVendor(t *testing.T) {
	ret := fake.DoNothing()
	expected := "Ok"
	if expected != ret {
		t.Errorf("fake returned %q, expected %q", ret, expected)
	}
}
