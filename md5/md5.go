package md5

import (
	"crypto/md5"
	"fmt"
	"io"
)

func EncryptString(original string) (key string) {
	Md5 := md5.New()
	io.WriteString(Md5, original)
	key = fmt.Sprintf("%x", md5.Sum(nil))
	return
}
