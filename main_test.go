package main

import (
	"os"
	"strconv"
	"testing"
)

const (
	testDirHome = "test"
	testDir     = testDirHome + "/something"
)

func createTestStructure() {
	os.MkdirAll(testDir, os.ModePerm)

	for i := 0; i < 10; i++ {
		os.Create(testDir + "/file" + strconv.Itoa(i))
	}
}

func deleteTestStructure() {
	os.RemoveAll(testDirHome)
}

func TestDir(t *testing.T) {
	createTestStructure()

	expected := "test/something/file9"
	actual := findFile("test", "file9")

	if expected != actual {
		t.Errorf("got value: %s, want: %s.", actual, expected)
	}

	// deleteTestStructure()
}
