package helper

import (
	"path"
	"strings"
)

func FileNameWithoutExt(fileName string) string {
	return strings.TrimSuffix(path.Base(fileName), path.Ext(fileName))
}
