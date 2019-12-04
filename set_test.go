package main

import "testing"

func TestCompareSets(t *testing.T) {
	type args struct {
		slice1 []Task
		slice2 []Task
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Zeros",
			args: args{
				slice1: []Task{},
				slice2: []Task{},
			},
			want: true,
		},

		{
			name: "Equal",
			args: args{
				slice1: []Task{
					{
						Method: "GET",
						Url:    "https://google.ru",
					},
					{
						Method: "POST",
						Url:    "https://yandex.ru",
					},
				},
				slice2: []Task{
					{
						Method: "GET",
						Url:    "https://google.ru",
					},
					{
						Method: "POST",
						Url:    "https://yandex.ru",
					},
				},
			},
			want: true,
		},

		{
			name: "Non-equal",
			args: args{
				slice1: []Task{{
					Method: "POST",
					Url:    "https://yandex.ru",
				},
					{
						Method: "GET",
						Url:    "https://google.ru",
					},
				},
				slice2: []Task{
					{
						Method: "GET",
						Url:    "https://google.ru",
					},
					{
						Method: "POST",
						Url:    "https://yandex.ru",
					},
				},
			},
			want: true,
		},
		{
			name: "Not matched",
			args: args{
				slice1: []Task{
					{
						Method: "GET",
						Url:    "https://google.ru",
					},
					{
						Method: "POST",
						Url:    "https://yandex.ru",
					},
				},
				slice2: []Task{
					{
						Method: "DELETE",
						Url:    "https://google.ru",
					},
					{
						Method: "PUTT",
						Url:    "https://yandex.ru",
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareSets(tt.args.slice1, tt.args.slice2); got != tt.want {
				t.Errorf("CompareSets() = %v, want %v", got, tt.want)
			}
		})
	}
}
