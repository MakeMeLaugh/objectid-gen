package internal

import (
	"testing"
	"time"
)

func TestNewObjectIDFromTimestamp(t *testing.T) {
	t.Parallel()
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "valid", args: args{t: time.Unix(1645218238, 0)}, want: "621009be0000000000000000"},
		{name: "zero", args: args{t: time.Unix(0, 0)}, want: "000000000000000000000000"},
		{name: "negative", args: args{t: time.Unix(-256, 0)}, want: "000000000000000000000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewObjectIDFromTimestamp(tt.args.t); got.String() != tt.want {
				t.Errorf("NewObjectIDFromTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectID_GetTimestamp(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		o    ObjectID
		ts   int64
		want time.Time
	}{
		{name: "valid", ts: 1645218238, want: time.Unix(1645218238, 0)},
		{name: "zero", ts: 0, want: time.Unix(0, 0)},
		{name: "negative", ts: -512, want: time.Unix(0, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o = NewObjectIDFromTimestamp(time.Unix(tt.ts, 0))
			if got := tt.o.GetTimestamp(); got.Unix() != tt.want.Unix() {
				t.Errorf("GetTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewObjectID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want string
	}{
		{name: "valid", want: NewObjectIDFromTimestamp(time.Now()).String()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// we can test only first 8 chars
			if got := NewObjectID(); got[0:8] != tt.want[0:8] {
				t.Errorf("NewObjectID() = %v, want %v", got[0:8], tt.want[0:8])
			}
		})
	}
}
