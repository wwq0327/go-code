package file

import (
	"path"
)

func Basename(file string) string {
	return path.Base(file)
}

func Dirname(file string) string {
	return path.Dir(file)
}
