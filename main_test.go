package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_parseInput(t *testing.T) {
	fixture := `35
		20
		15
		25
		47
		40
		62
		55
		65
		95
		102
		117
		150
		182
		127
		219
		299
		277
		309
		576`
	want := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	got, err := parseInput(strings.NewReader(fixture))
	if (err != nil) != false {
		t.Errorf("parseInput() error = %v, wantErr %v", err, false)
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parseInput() = %v, want %v", got, want)
	}
}
func Test_numberIsValid(t *testing.T) {
	type args struct {
		preamble []int
		number   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test 1", args{[]int{35, 20, 15, 25, 47}, 40}, true},
		{"test 2", args{[]int{20, 15, 25, 47, 40}, 62}, true},
		{"test 3", args{[]int{15, 25, 47, 40, 62}, 55}, true},
		{"test 4", args{[]int{95, 102, 117, 150, 182}, 127}, false},
		{"test 5", args{[]int{102, 117, 150, 182, 127}, 219}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberIsValid(tt.args.preamble, tt.args.number); got != tt.want {
				t.Errorf("numberIsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
