package logger

import (
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name string
		want *Logger
	}{
		// TODO: Add test cases.
		{
			name: "TestNewLogger",
			want: &Logger{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger("debug"); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	tests := []struct {
		name       string
		l          *Logger
		identifier string
		msg        string
	}{
		// TODO: Add test cases.
		{
			name:       "Debug",
			l:          &Logger{Lv: FATAL},
			identifier: "Debug",
			msg:        "This is a Debug log.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Debug(tt.identifier, tt.msg)
		})
	}
}

func TestLogger_Trace(t *testing.T) {
	tests := []struct {
		name       string
		l          *Logger
		identifier string
		msg        string
	}{
		// TODO: Add test cases.
		{
			name:       "Debug",
			l:          &Logger{Lv: DEBUG},
			identifier: "Debug",
			msg:        "This is a Debug log.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Trace(tt.identifier, tt.msg)
		})
	}
}

func TestLogger_Fatal(t *testing.T) {
	tests := []struct {
		name string
		l    *Logger
		identifier string
		msg        string
	}{
		// TODO: Add test cases.
		{
			name:       "Debug",
			l:          &Logger{Lv: DEBUG},
			identifier: "Fatal",
			msg:        "This is a Fatal log.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Fatal(tt.identifier, tt.msg)
		})
	}
}
