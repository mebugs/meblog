package log

import (
	"testing"
)

func TestSetLevel(t *testing.T) {
	setLevelIndex(WarnLevel)
	Debug("debug")
	Info("info")
	Warn("warn")

	// 还原
	setLevelIndex(DebugLevel)
}
