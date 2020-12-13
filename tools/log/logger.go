package log

import (
	"bytes"
	"fmt"
	"github.com/ggdream/gopack/tools/color"
	"io"
	"reflect"
	"time"
	"unsafe"
)

type Logger struct {
	level  int
	color  bool
	timer  bool
	writer io.Writer
}

func (l *Logger) SetLevel(level int) {
	l.level = level
}

func (l *Logger) SetWriter(writer io.Writer) {
	l.writer = writer
}

func (l *Logger) SetColorSupport(isSupport bool) {
	l.color = isSupport
}

func (l *Logger) OpenTimer() {
	l.timer = true
}

// @format: 模板字符串
// @args: 位置参数
func (l *Logger) ConsoleLog(format string, args ...interface{}) {
	text := fmt.Sprintf(format, args...)
	textDeep := (*reflect.StringHeader)(unsafe.Pointer(&text))
	sliceDeep := reflect.SliceHeader{
		Data: textDeep.Data,
		Len:  textDeep.Len,
		Cap:  textDeep.Len,
	}

	_, _ = l.writer.Write(*(*[]byte)(unsafe.Pointer(&sliceDeep)))
}

func (l *Logger) WrapLog(level int, colour, format string, args ...interface{}) {
	if level > l.level {
		return
	}

	if l.color {
		var buf bytes.Buffer
		buf.WriteString(colour)
		buf.WriteString(format)
		buf.WriteString(color.ColorEnd)
		format = buf.String()
	}

	if l.timer {
		format = fmt.Sprintf("[%s] ", time.Now().Format("2006-01-02 15:04:05")) + format
	}

	l.ConsoleLog(format, args...)
}
