package superhasher

import (
	"crypto/md5"
	"fmt"
)

func Superhash(path string) string {
	return path
}

func calculateMD5(data []byte) string {
	result := md5.Sum(data)
	s := fmt.Sprintf("%x", result)
	return s
}
