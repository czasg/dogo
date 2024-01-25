package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func NewHash(secret []byte) Hash {
	return Hash{secret: secret}
}

type Hash struct {
	secret []byte
}

func (h Hash) Md5(data ...[]byte) string {
	hashMd5 := md5.New()
	for _, p := range data {
		hashMd5.Write(p)
	}
	return strings.ToUpper(hex.EncodeToString(hashMd5.Sum(h.secret)))
}

func (h Hash) Sha1(data ...[]byte) string {
	hashSha1 := sha1.New()
	for _, p := range data {
		hashSha1.Write(p)
	}
	return strings.ToUpper(hex.EncodeToString(hashSha1.Sum(h.secret)))
}

func (h Hash) Sha256(data ...[]byte) string {
	hashSha256 := sha256.New()
	for _, p := range data {
		hashSha256.Write(p)
	}
	return strings.ToUpper(hex.EncodeToString(hashSha256.Sum(h.secret)))
}
