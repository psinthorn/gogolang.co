package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encrypt(input string) string {

	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))

}
