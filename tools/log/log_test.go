package log

import "testing"

func TestLogger_Info(t *testing.T) {
	l := New()
	l.Debug("hello, rust!\n")
	l.Info("hello, golang!\n")
	l.Warning("hello, typescript!\n")
	l.Error("hello, python!\n")
}
