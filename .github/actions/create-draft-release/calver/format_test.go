package main

import (
	"strings"
	"testing"
	"time"
)

func TestGetNextVersion_Date(t *testing.T) {
	tag, _ := GetNextVersion("20180525.1")

	const layout = "20060102"
	today := time.Now().Format(layout)
	expected := today + ".1"
	if tag != expected {
		t.Errorf("Output=%q, Expected=%q", tag, expected)
	}
}

func TestGetNextVersion_DatePlus(t *testing.T) {
	const layout = "20060102"
	today := time.Now().Format(layout)
	tag, _ := GetNextVersion(today + ".1")

	expected := today + ".2"
	if tag != expected {
		t.Errorf("Output=%q, Expected=%q", tag, expected)
	}
}

func TestGetNextVersion_DateWithPrefix(t *testing.T) {
	const prefix = "release_"
	tag, _ := GetNextVersion(prefix + "20180525.1")

	const layout = "20060102"
	today := time.Now().Format(layout)
	expected := prefix + today + ".1"
	if tag != expected {
		t.Errorf("Output=%q, Expected=%q", tag, expected)
	}
}

func TestGetNextVersion_DatePlusWithPrefix(t *testing.T) {
	const prefix = "release_"
	const layout = "20060102"
	today := time.Now().Format(layout)
	tag, _ := GetNextVersion(prefix + today + ".1")

	expected := prefix + today + ".2"
	if tag != expected {
		t.Errorf("Output=%q, Expected=%q", tag, expected)
	}
}

func TestGetNextVersion_DateError(t *testing.T) {
	const layout = "20060102"
	today := time.Now().Format(layout)
	_, err := GetNextVersion(today + ".date")

	expected := "invalid syntax"
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("Output=%q, Expected=%q", err.Error(), expected)
	}
}
