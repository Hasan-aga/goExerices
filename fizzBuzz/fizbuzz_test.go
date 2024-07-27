package main

import (
	"strings"
	"testing"
)

func Test_proceduralFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Check 15th element is FizzBuzz",
			want: "FizzBuzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := proceduralFizzBuzz(1000000)
			elements := strings.Split(got, "\n")
			if len(elements) < 15 {
				t.Errorf("proceduralFizzBuzz() returned less than 15 elements")
			} else if elements[14] != tt.want {
				t.Errorf("15th element of proceduralFizzBuzz() = %v, want %v", elements[14], tt.want)
			}
		})
	}
}
