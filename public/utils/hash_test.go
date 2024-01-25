package utils

import "testing"

func TestHash_Md5(t *testing.T) {
	type fields struct {
		secret []byte
	}
	type args struct {
		data [][]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "",
			fields: fields{secret: nil},
			args: args{
				data: [][]byte{
					[]byte("www"),
					[]byte("."),
					[]byte("baidu"),
					[]byte("."),
					[]byte("com"),
				},
			},
			want: "DAB19E82E1F9A681EE73346D3E7A575E",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHash(tt.fields.secret)
			if got := h.Md5(tt.args.data...); got != tt.want {
				t.Errorf("Md5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash_Sha1(t *testing.T) {
	type fields struct {
		secret []byte
	}
	type args struct {
		data [][]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "",
			fields: fields{secret: nil},
			args: args{
				data: [][]byte{
					[]byte("www"),
					[]byte("."),
					[]byte("baidu"),
					[]byte("."),
					[]byte("com"),
				},
			},
			want: "31E50A13CDFA1BC2A6A0C2A31B74E3340B7A2DBC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHash(tt.fields.secret)
			if got := h.Sha1(tt.args.data...); got != tt.want {
				t.Errorf("Sha1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash_Sha256(t *testing.T) {
	type fields struct {
		secret []byte
	}
	type args struct {
		data [][]byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "",
			fields: fields{secret: nil},
			args: args{
				data: [][]byte{
					[]byte("www"),
					[]byte("."),
					[]byte("baidu"),
					[]byte("."),
					[]byte("com"),
				},
			},
			want: "ECEB28B70500B741F3BAAEA651C6EC8298631AC1B3BF401BCDD1527B4A48AD0C",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHash(tt.fields.secret)
			if got := h.Sha256(tt.args.data...); got != tt.want {
				t.Errorf("Sha256() = %v, want %v", got, tt.want)
			}
		})
	}
}
