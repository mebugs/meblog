package security

import (
	"testing"
)

// SafeJson 为keys的值进行安全脱敏
func TestSafeJson(t *testing.T) {
	SafeJson(`{"name" : "123",{"name" : "12345",{"phone":"12345678901"}}`, []string{"name", "phone"})
}
