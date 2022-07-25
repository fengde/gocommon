package toolx

import (
	"reflect"
	"testing"
)

func TestString2bytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{
				s: "hello world",
			},
			want: []byte("hello world"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String2bytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String2bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytes2string(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				b: []byte("hello world"),
			},
			want: "hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes2string(tt.args.b); got != tt.want {
				t.Errorf("Bytes2string() = %v, want %v", got, tt.want)
			}
		})
	}
}
