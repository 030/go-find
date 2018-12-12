package find

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
	actual, _ := File("test", "file9")

	if expected != actual {
		t.Errorf("got value: %s, want: %s.", actual, expected)
	}

	deleteTestStructure()
}

func TestErrors(t *testing.T) {
	expected := "File file9 was not found in directory test2"
	_, err := File("test2", "file9")

	if err.Error() != expected {
		t.Errorf("got value: %s, want: %s.", err, expected)
	}
}
