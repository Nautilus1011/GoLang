package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Level はログレベル
type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	default:
		return "UNKNOWN"
	}
}

// Logger は構造化ロギングを行うロガー
type Logger struct {
	output   io.Writer
	minLevel Level
}

// NewLogger は新しいLoggerを作成
func NewLogger(output io.Writer, minLevel Level) *Logger {
	return &Logger{
		output:   output,
		minLevel: minLevel,
	}
}

// LogEntry はログエントリの構造
type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

// log は実際のログ出力を行う内部メソッド
func (l *Logger) log(level Level, ctx context.Context, msg string, keysAndValues ...interface{}) {
	if level < l.minLevel {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level.String(),
		Message:   msg,
		Fields:    make(map[string]interface{}),
	}

	// キーと値のペアを追加
	for i := 0; i < len(keysAndValues); i += 2 {
		if i+1 < len(keysAndValues) {
			key := fmt.Sprintf("%v", keysAndValues[i])
			entry.Fields[key] = keysAndValues[i+1]
		}
	}

	// JSON形式で出力
	data, err := json.Marshal(entry)
	if err != nil {
		fmt.Fprintf(l.output, "ログのマーシャルに失敗: %v\n", err)
		return
	}

	fmt.Fprintf(l.output, "%s\n", data)
}

// Debug はDEBUGレベルのログを出力
func (l *Logger) Debug(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.log(LevelDebug, ctx, msg, keysAndValues...)
}

// Info はINFOレベルのログを出力
func (l *Logger) Info(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.log(LevelInfo, ctx, msg, keysAndValues...)
}

// Warn はWARNレベルのログを出力
func (l *Logger) Warn(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.log(LevelWarn, ctx, msg, keysAndValues...)
}

// Error はERRORレベルのログを出力
func (l *Logger) Error(ctx context.Context, msg string, keysAndValues ...interface{}) {
	l.log(LevelError, ctx, msg, keysAndValues...)
}

// WithFields はフィールドを追加した新しいロガーを返す（チェーン用）
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	// 簡易実装: 実際のプロダクションではより高度な実装が必要
	return l
}