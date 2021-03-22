package checksum

import (
	"crypto/md5"
	"fmt"
	"io"
)

// 计算单个字符串MD5
func MD5Str(message string) string {
	return MD5Strs(message)
}

// 计算多个字符串的MD5
func MD5Strs(messages ...string) string {
	h := md5.New()
	for _, msg := range messages {
		io.WriteString(h, msg)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
