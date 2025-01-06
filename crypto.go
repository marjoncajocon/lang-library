package lib

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	hashedBytes := hash.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}

func Md5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	hashedBytes := hash.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	return hashedString
}
