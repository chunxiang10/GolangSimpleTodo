package utils

import (
	"crypto/md5"
	"encoding/hex"
)

var Salt = "dayuehelpyou"

func EncryMd5(str string) string {
	ctx := md5.New()
	ctx.Write([]byte(str))
	newpwd := hex.EncodeToString(ctx.Sum(nil)) + Salt
	ctx.Write([]byte(newpwd))
	return hex.EncodeToString(ctx.Sum(nil))
}
