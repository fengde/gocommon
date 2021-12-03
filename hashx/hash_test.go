package hashx

import (
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test1",
			args: args{data: []byte("hello")},
			want: 14688674573012802306,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.data); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5(t *testing.T) {
	Md5([]byte("hello"))
}

func TestMd5Hex(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{data: []byte("hello")},
			want: "5d41402abc4b2a76b9719d911017c592",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Hex(tt.args.data); got != tt.want {
				t.Errorf("Md5Hex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSha256(t *testing.T) {
	Sha256([]byte("hello"))
}

func TestSha256Hex(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{data: []byte("hello")},
			want: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha256Hex(tt.args.data); got != tt.want {
				t.Errorf("Sha256Hex() = %v, want %v", got, tt.want)
			}
		})
	}
}
