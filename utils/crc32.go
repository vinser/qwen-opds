package utils

import (
	"hash/crc32"
	"io"
)

func CalculateCRC32(file io.Reader) uint32 {
	hash := crc32.NewIEEE()
	_, _ = io.Copy(hash, file)
	return hash.Sum32()
}
