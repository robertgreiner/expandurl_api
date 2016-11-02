package main

import (
	"testing"
)

func TestCanExpandEncodedURL(t *testing.T) {
	result := ExpandURL("http:%2F%2Fbit.ly%2F2eXwN8A")

	if result.ExpandedURL != "http://robertgreiner.com/" {
		t.Error("Returned incorrect URL")
	}
}

func TestCanExpandURL(t *testing.T) {
	result := ExpandURL("http://bit.ly/2eXwN8A")

	if result.ExpandedURL != "http://robertgreiner.com/" {
		t.Error("Returned incorrect URL")
	}
}
