package logger

import (
	"fmt"
	"io"
	"os"
)

type LogLevel string

const (
	LevelInfo    LogLevel = "INFO"
	LevelWarning LogLevel = "WARNING"
	LevelError   LogLevel = "ERROR"
)

type Logger struct {
	formatter Formatter
	output    io.Writer
}

type LoggerOption func(*Logger)

// Thiết lập một formatter tùy chỉnh cho Logger.
func WithFormatter(f Formatter) LoggerOption {
	return func(l *Logger) {
		l.formatter = f
	}
}

// Tạo một instance Logger mới.
func NewLogger(options ...LoggerOption) *Logger {
	l := &Logger{
		formatter: DefaultFormatter(), // Mặc định sử dụng DefaultFormatter
		output:    os.Stdout,
	}

	for _, opt := range options {
		opt(l)
	}

	return l
}

// Hàm cốt lõi để ghi log.
// Nó NHẬN THÔNG BÁO ĐÃ ĐƯỢC ĐỊNH DẠNG (message), không phải format string và args.
func (l *Logger) logInternal(level LogLevel, message string) {
	formattedMessage := l.formatter(string(level), message)
	fmt.Fprintln(l.output, formattedMessage)
}

// Ghi một thông báo log ở cấp độ INFO.
func (l *Logger) Info(message string) {
	l.logInternal(LevelInfo, message)
}

// Ghi một thông báo log ở cấp độ WARNING.
func (l *Logger) Warning(message string) {
	l.logInternal(LevelWarning, message)
}

// Ghi một thông báo log ở cấp độ ERROR.
func (l *Logger) Error(message string) {
	l.logInternal(LevelError, message)
}

// Ghi log với định dạng giống như fmt.Printf.
func (l *Logger) Logf(level LogLevel, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	// Truyền tin nhắn đã định dạng vào logInternal
	l.logInternal(level, message)
}