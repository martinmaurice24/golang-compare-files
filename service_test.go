package diff

import (
	"fmt"
	"testing"
)

func TestCompare(t *testing.T) {
	testsEntries := []struct {
		CurrentFilePath string
		NewFilePath     string
		isSameFile      bool
	}{
		{
			CurrentFilePath: "../assets/old.txt",
			NewFilePath:     "../assets/new.txt",
			isSameFile:      false,
		},
		{
			CurrentFilePath: "../assets/old.txt",
			NewFilePath:     "../assets/old.txt",
			isSameFile:      true,
		},
		{
			CurrentFilePath: "../assets/gopher.png",
			NewFilePath:     "../assets/mario.png",
			isSameFile:      false,
		},
		{
			CurrentFilePath: "../assets/mario.png",
			NewFilePath:     "../assets/mario.png",
			isSameFile:      true,
		},
	}

	for i, entry := range testsEntries {
		t.Run(fmt.Sprintf("%s%d", "test", i), func(st *testing.T) {
			isSame, _ := CompareLines(entry.CurrentFilePath, entry.NewFilePath)
			if isSame != entry.isSameFile {
				t.Errorf("The expected value is %t but we have %t", entry.isSameFile, isSame)
			}
		})
	}
}

func TestFindRemovedLines(t *testing.T) {
	isSame, removedLines := FindRemovedLines("../assets/old.txt", "../assets/new.txt")
	if isSame != false {
		t.Errorf("Expected isSame to be false but it's true")
	}

	if s := len(removedLines); s != 3 {
		t.Errorf("Expected removed lines count to be 3 we get %d", s)
	}
}

func TestFindNewLines(t *testing.T) {
	isSame, newLines := FindNewLines("../assets/old.txt", "../assets/new.txt")
	if isSame != false {
		t.Errorf("Expected isSame to be false but it's true")
	}

	if s := len(newLines); s != 2 {
		t.Errorf("Expected new lines count to be 2 we get %d", s)
	}
}
