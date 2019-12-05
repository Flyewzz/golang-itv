package main

import "testing"

func TestCheckMethodValid(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Check empty",
			args: args{
				"",
			},
			want: false,
		},
		{
			name: "Check incorrect method",
			args: args{
				"INCORREct",
			},
			want: false,
		},
		{
			name: "Check GET",
			args: args{
				"gEt",
			},
			want: true,
		},
		{
			name: "Check POST",
			args: args{
				"pOST",
			},
			want: true,
		},
		{
			name: "Check PUT",
			args: args{
				"put",
			},
			want: true,
		},
		{
			name: "Check DELETE",
			args: args{
				"DELETE",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckMethodValid(tt.args.name); got != tt.want {
				t.Errorf("CheckMethodValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
