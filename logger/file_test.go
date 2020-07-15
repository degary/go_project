package logger

import (
	"testing"
)

func TestFileLogger(t *testing.T){
	logger := NewFileLogger(LogLevelDebug,"C:\\Users\\denghui\\Desktop","test")
	logger.Debug("user id[%d] is come from china",2131)
	logger.Warn("test warn")
	logger.Fatal("test fatal")
	logger.Close()
}