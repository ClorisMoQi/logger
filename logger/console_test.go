package logger

import (
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name  string
		lvstr string
		want  *Logger
	}{
		{
			name:  "DEBUG",
			lvstr: "DEBUG",
			want:  &Logger{Lv: DEBUG},
		},
		{
			name:  "TRACE",
			lvstr: "TRACE",
			want:  &Logger{Lv: TRACE},
		},
		{
			name:  "INFO",
			lvstr: "INFO",
			want:  &Logger{Lv: INFO},
		},
		{
			name:  "WARN",
			lvstr: "WARN",
			want:  &Logger{Lv: WARN},
		},
		{
			name:  "ERROR",
			lvstr: "ERROR",
			want:  &Logger{Lv: ERROR},
		},
		{
			name:  "FATAL",
			lvstr: "FATAL",
			want:  &Logger{Lv: FATAL},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(tt.lvstr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogger_Debug(t *testing.T) {
	id := 10010
	name := "DebugLog"
	tests := []struct {
		name       string
		l          *Logger
		identifier string
		format     string
		args       []interface{}
	}{
		{
			name:       "DEBUG",
			l:          &Logger{Lv: DEBUG},
			identifier: "debug",
			format:     "This is a debug log, id: %d, name:%s\n",
		},
		{
			name:       "TRACE",
			l:          &Logger{Lv: TRACE},
			identifier: "trace",
			format:     "This is a trace log, id: %d, name:%s\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Debug(tt.identifier, tt.format, id, name)
		})
	}
}

func TestLogger_Trace(t *testing.T) {
	type args struct {
		identifier string
		format     string
		args       []interface{}
	}
	tests := []struct {
		name string
		l    *Logger
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Trace(tt.args.identifier, tt.args.format, tt.args.args...)
		})
	}
}

func TestLogger_Info(t *testing.T) {
	type args struct {
		identifier string
		format     string
		args       []interface{}
	}
	tests := []struct {
		name string
		l    *Logger
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Info(tt.args.identifier, tt.args.format, tt.args.args...)
		})
	}
}

func TestLogger_Warn(t *testing.T) {
	type args struct {
		identifier string
		format     string
		args       []interface{}
	}
	tests := []struct {
		name string
		l    *Logger
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Warn(tt.args.identifier, tt.args.format, tt.args.args...)
		})
	}
}

func TestLogger_Error(t *testing.T) {
	type args struct {
		identifier string
		format     string
		args       []interface{}
	}
	tests := []struct {
		name string
		l    *Logger
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Error(tt.args.identifier, tt.args.format, tt.args.args...)
		})
	}
}

func TestLogger_Fatal(t *testing.T) {
	type args struct {
		identifier string
		format     string
		args       []interface{}
	}
	tests := []struct {
		name string
		l    *Logger
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Fatal(tt.args.identifier, tt.args.format, tt.args.args...)
		})
	}
}

func TestLogger_log(t *testing.T) {
	type args struct {
		lv         LogLevel
		identifier string
		format     string
		args       []interface{}
	}
	tests := []struct {
		name string
		l    *Logger
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.log(tt.args.lv, tt.args.identifier, tt.args.format, tt.args.args...)
		})
	}
}
