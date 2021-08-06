package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	uuid "github.com/satori/go.uuid"
)

// MD5V md5加密
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func UUIDV5() uuid.UUID {
	return uuid.NewV5(uuid.NamespaceOID, uuid.NewV4().String()+time.Now().Format(time.RFC3339))
}
