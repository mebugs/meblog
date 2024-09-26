package security

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 哈希
func SHA256(str string) string {
	sig := sha256.Sum256([]byte(str))
	return fmt.Sprintf("%x", sig[:])
}
