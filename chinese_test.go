package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func TestIsChinese(t *testing.T) {
	f := path.Join(strings.Split(os.Getenv("GOPATH"), ":")[0], "src",
		PackageDir, CJKUnifiedIdeographs)

	if s, e := ioutil.ReadFile(f); e != nil {
		t.Fatalf("Failed to load testdata file: %v", e)
	} else {
		if r := StrictlyChinese(string(s)); r != nil {
			t.Errorf("Unexpected error: %v", r)
		}
	}
}

const (
	PackageDir           = "github.com/wangkuiyi/icu/"
	CJKUnifiedIdeographs = "testdata/CJKUnifiedIdeographs"
)
