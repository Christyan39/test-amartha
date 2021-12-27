package helper

import (
	"regexp"
	"testing"
)

func TestRandString(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success Only",
			args: args{
				length: 6,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandString(tt.args.length)
			ok, _ := regexp.Match(RegexpValidator, []byte(got))
			if !ok {
				t.Errorf("RandString() = %v, want meet expresion %v", got, RegexpValidator)
			}
		})
	}
}
