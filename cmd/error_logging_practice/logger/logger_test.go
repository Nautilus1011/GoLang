package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestLoggerLevels(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(&buf, LevelInfo)

	ctx := context.Background()

	// Debugは出力されない（LevelInfo以上のみ）
	log.Debug(ctx, "デバッグメッセージ")
	if buf.Len() > 0 {
		t.Error("Debugメッセージが出力されてしまった")
	}

	// Infoは出力される
	log.Info(ctx, "情報メッセージ")
	if buf.Len() == 0 {
		t.Error("Infoメッセージが出力されていない")
	}
}

func TestLoggerOutput(t *testing.T) {
	var buf bytes.Buffer
	log := NewLogger(&buf, LevelDebug)

	ctx := context.Background()

	log.Info(ctx, "テストメッセージ", "key1", "value1", "key2", 123)

	output := buf.String()
	if output == "" {
		t.Fatal("ログが出力されていない")
	}

	// JSON形式の検証
	var entry LogEntry
	err := json.Unmarshal([]byte(output), &entry)
	if err != nil {
		t.Fatalf("JSONのパースに失敗: %v", err)
	}

	// 内容の検証
	if entry.Level != "INFO" {
		t.Errorf("期待されるレベル: INFO, 実際: %s", entry.Level)
	}

	if entry.Message != "テストメッセージ" {
		t.Errorf("期待されるメッセージ: テストメッセージ, 実際: %s", entry.Message)
	}

	if entry.Fields["key1"] != "value1" {
		t.Errorf("期待される値: value1, 実際: %v", entry.Fields["key1"])
	}

	// key2は数値型なのでfloat64として解釈される
	if entry.Fields["key2"] != float64(123) {
		t.Errorf("期待される値: 123, 実際: %v", entry.Fields["key2"])
	}
}

func TestLoggerAllLevels(t *testing.T) {
	tests := []struct {
		name     string
		logFunc  func(*Logger, context.Context, string, ...interface{})
		expected string
	}{
		{
			name: "Debug",
			logFunc: func(l *Logger, ctx context.Context, msg string, kvs ...interface{}) {
				l.Debug(ctx, msg, kvs...)
			},
			expected: "DEBUG",
		},
		{
			name: "Info",
			logFunc: func(l *Logger, ctx context.Context, msg string, kvs ...interface{}) {
				l.Info(ctx, msg, kvs...)
			},
			expected: "INFO",
		},
		{
			name: "Warn",
			logFunc: func(l *Logger, ctx context.Context, msg string, kvs ...interface{}) {
				l.Warn(ctx, msg, kvs...)
			},
			expected: "WARN",
		},
		{
			name: "Error",
			logFunc: func(l *Logger, ctx context.Context, msg string, kvs ...interface{}) {
				l.Error(ctx, msg, kvs...)
			},
			expected: "ERROR",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log := NewLogger(&buf, LevelDebug)
			ctx := context.Background()

			tt.logFunc(log, ctx, "テスト")

			output := buf.String()
			if !strings.Contains(output, tt.expected) {
				t.Errorf("出力に%sが含まれていない: %s", tt.expected, output)
			}
		})
	}
}

func TestLoggerMinLevel(t *testing.T) {
	tests := []struct {
		name      string
		minLevel  Level
		logLevel  Level
		shouldLog bool
	}{
		{"Debug on Info", LevelInfo, LevelDebug, false},
		{"Info on Info", LevelInfo, LevelInfo, true},
		{"Warn on Info", LevelInfo, LevelWarn, true},
		{"Error on Info", LevelInfo, LevelError, true},
		{"Info on Error", LevelError, LevelInfo, false},
		{"Error on Error", LevelError, LevelError, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			log := NewLogger(&buf, tt.minLevel)
			ctx := context.Background()

			log.log(tt.logLevel, ctx, "テスト")

			hasOutput := buf.Len() > 0
			if hasOutput != tt.shouldLog {
				t.Errorf("期待される出力有無: %v, 実際: %v", tt.shouldLog, hasOutput)
			}
		})
	}
}
