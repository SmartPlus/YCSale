package web

import (
	"path/filepath"
	"runtime"
)

func FilenameFromTheSameDir(dest string) string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(filename), dest)
}
