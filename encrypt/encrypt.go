package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Hash(text string) string {
	md5 := md5.New()
	md5.Write([]byte(text))
	return hex.EncodeToString(md5.Sum(nil))
}
