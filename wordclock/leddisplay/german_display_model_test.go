package leddisplay

import (
	"reflect"
	"testing"
)

func TestGermanLedDisplayModel_getNonHourWord(t *testing.T) {
	type fields struct {
		nonHourWords map[string][]int
		hourWords    map[string][]int
	}
	type args struct {
		word string
	}
	emptyFields := fields{
		nonHourWords: make(map[string][]int),
		hourWords:    make(map[string][]int),
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []int
	}{
		{
			name:   "Test start index",
			fields: emptyFields,
			args: args{
				word: "ES",
			},
			want: []int{0, 1},
		},
		{
			name:   "Test end index",
			fields: emptyFields,
			args: args{
				word: "UHR",
			},
			want: []int{107, 108, 109},
		},
		{
			name:   "Test duplicate word",
			fields: emptyFields,
			args: args{
				word: "FÜNF",
			},
			want: []int{7, 8, 9, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disp := &GermanLedDisplayModel{
				nonHourWords: tt.fields.nonHourWords,
				hourWords:    tt.fields.hourWords,
			}
			if got := disp.getNonHourWord(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GermanLedDisplayModel.getNonHourWord() (Case:%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
