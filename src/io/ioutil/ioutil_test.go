// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ioutil_test

import (
	. "io/ioutil"
	"os"
	"testing"
)

func checkSize(t *testing.T, path string, size int64) {
	dir, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat %q (looking for size %d): %s", path, size, err)
	}
	if dir.Size() != size {
		t.Errorf("Stat %q: size %d want %d", path, dir.Size(), size)
	}
}

func TestReadDir(t *testing.T) {
	dirname := "rumpelstilzchen"
	_, err := ReadDir(dirname)
	if err == nil {
		t.Fatalf("ReadDir %s: error expected, none found", dirname)
	}

	dirname = ".."
	list, err := ReadDir(dirname)
	if err != nil {
		t.Fatalf("ReadDir %s: %v", dirname, err)
	}

	foundFile := false
	foundSubDir := false
	for _, dir := range list {
		switch {
		case !dir.IsDir() && dir.Name() == "io_test.go":
			foundFile = true
		case dir.IsDir() && dir.Name() == "ioutil":
			foundSubDir = true
		}
	}
	if !foundFile {
		t.Fatalf("ReadDir %s: io_test.go file not found", dirname)
	}
	if !foundSubDir {
		t.Fatalf("ReadDir %s: ioutil directory not found", dirname)
	}
}
