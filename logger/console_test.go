package logger

import (
	"reflect"
	"testing"
)

func TestNewLogger(t *testing.T) {
	tests := []struct {
		name  string
		lvstr string
		want  *CLogger
	}{
		{
			name:  "DEBUG",
			lvstr: "DEBUG",
			want:  &CLogger{Lv: DEBUG},
		},
		{
			name:  "TRACE",
			lvstr: "TRACE",
			want:  &CLogger{Lv: TRACE},
		},
		{
			name:  "INFO",
			lvstr: "INFO",
			want:  &CLogger{Lv: INFO},
		},
		{
			name:  "WARN",
			lvstr: "WARN",
			want:  &CLogger{Lv: WARN},
		},
		{
			name:  "ERROR",
			lvstr: "ERROR",
			want:  &CLogger{Lv: ERROR},
		},
		{
			name:  "FATAL",
			lvstr: "FATAL",
			want:  &CLogger{Lv: FATAL},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCLogger(tt.lvstr); !reflect.DeepEqual(got, tt.want) {
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
		l          *CLogger
		identifier string
		format     string
		args       []interface{}
	}{
		{
			name:       "DEBUG",
			l:          &CLogger{Lv: DEBUG},
			identifier: "debug",
			format:     "This is a debug log, id: %d, name:%s\n",
		},
		{
			name:       "TRACE",
			l:          &CLogger{Lv: TRACE},
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
		l    *CLogger
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
		l    *CLogger
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
		l    *CLogger
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
		l    *CLogger
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
		l    *CLogger
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
		l    *CLogger
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
