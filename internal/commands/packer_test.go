package commands

import "testing"

func TestUnpack(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				input: "a4bc2d5e",
			},
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name: "Test 2",
			args: args{
				input: "abcd",
			},
			want:    "abcd",
			wantErr: false,
		},
		{
			name: "Test 3",
			args: args{
				input: "45",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test 4",
			args: args{
				input: "",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Test 5",
			args: args{
				input: "a10",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unpack(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unpack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Unpack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPack(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test 1",
			args: args{
				input: "aaaabccddddde",
			},
			want:    "a4bc2d5e",
			wantErr: false,
		},
		{
			name: "Test 2",
			args: args{
				input: "abcd",
			},
			want:    "abcd",
			wantErr: false,
		},
		{
			name: "Test 3",
			args: args{
				input: "45",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test 4",
			args: args{
				input: "",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Pack(tt.args.input); got != tt.want {
				t.Errorf("Pack() = %v, want %v", got, tt.want)
			}
		})
	}
}
