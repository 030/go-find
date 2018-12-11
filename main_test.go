package main

import (
	"os"
	"testing"
)

const (
	testDirHome = "test"
	testDir     = testDirHome + "/something"
)

func createTestStructure() {
	os.MkdirAll(testDir, os.ModePerm)
	os.Create(testDir + "/file")
}

func deleteTestStructure() {
	os.RemoveAll(testDirHome)
}

func TestDir(t *testing.T) {
	createTestStructure()

	expected := "world"
	actual := find()

	if expected != actual {
		t.Errorf("got value: %s, want: %s.", actual, expected)
	}

	deleteTestStructure()
}
