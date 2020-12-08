package lob

import "testing"

func TestClient_IsTest(t *testing.T) {
	c := NewClient("test_blah")

	if !c.IsTest() {
		t.Fatal("client should be test")
	}

	c = NewClient("abc123")
	if c.IsTest() {
		t.Fatal("client should not be test")
	}
}
