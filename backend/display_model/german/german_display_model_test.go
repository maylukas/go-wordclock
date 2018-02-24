package german

import (
	"reflect"
	"testing"
	"time"
	"sort"
)

type testFields struct {
	nonHourWords map[string][]int
	hourWords    map[string][]int
}
var emptyFields = testFields{
nonHourWords: make(map[string][]int),
hourWords:    make(map[string][]int),
}
func TestLedDisplayModel_translateTimeToWords(t *testing.T) {
	type args struct {
		time time.Time
	}

	tests := []struct {
		name string
		args args
		fields testFields
		wantControlWords []string
		wantHourWords []string
	}{
		{
			name: "Test Midnight",
			args: args {
				time: time.Date(2018, 1,1,0,0,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "UHR"},
			wantHourWords: []string {"ZWÖLF"},
			fields: emptyFields,
		},
		{
			name: "Test half past midnight",
			args: args {
				time: time.Date(2018, 1,1,0,30,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "HALB"},
			wantHourWords: []string {"EINS"},
			fields: emptyFields,
		},
		{
			name: "Test half past noon 24 hour time",
			args: args {
				time: time.Date(2018, 1,1,12,30,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "HALB"},
			wantHourWords: []string {"EINS"},
			fields: emptyFields,
		},
		{
			name: "Test half past midnight with offset after",
			args: args {
				time: time.Date(2018, 1,1,0,34,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "HALB"},
			wantHourWords: []string {"EINS"},
			fields: emptyFields,
		},
		{
			name: "Test half past midnight with offset before",
			args: args {
				time: time.Date(2018, 1,1,0,26,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "FÜNF", "VOR", "HALB"},
			wantHourWords: []string {"EINS"},
			fields: emptyFields,
		},
		{
			name: "Test five after 2",
			args: args {
				time: time.Date(2018, 1,1,2,5,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "FÜNF", "NACH"},
			wantHourWords: []string {"ZWEI"},
			fields: emptyFields,
		},
		{
			name: "Test ten after 3",
			args: args {
				time: time.Date(2018, 1,1,3,10,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "ZEHN", "NACH"},
			wantHourWords: []string {"DREI"},
			fields: emptyFields,
		},
		{
			name: "Test quarter after 4",
			args: args {
				time: time.Date(2018, 1,1,4,17,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "VIERTEL", "NACH"},
			wantHourWords: []string {"VIER"},
			fields: emptyFields,
		},
		{
			name: "Test 20 after 5",
			args: args {
				time: time.Date(2018, 1,1,5,21,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "ZWANZIG", "NACH"},
			wantHourWords: []string {"FÜNF"},
			fields: emptyFields,
		},
		{
			name: "Test half past 6",
			args: args {
				time: time.Date(2018, 1,1,6,33,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "HALB"},
			wantHourWords: []string {"SIEBEN"},
			fields: emptyFields,
		},
		{
			name: "Test 35 past 7",
			args: args {
				time: time.Date(2018, 1,1,7,38,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "FÜNF", "NACH", "HALB"},
			wantHourWords: []string {"ACHT"},
			fields: emptyFields,
		},
		{
			name: "Test 40 past 8",
			args: args {
				time: time.Date(2018, 1,1,8,44,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "ZWANZIG", "VOR"},
			wantHourWords: []string {"NEUN"},
			fields: emptyFields,
		},
		{
			name: "Test 45 past 9",
			args: args {
				time: time.Date(2018, 1,1,9,49,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "VIERTEL", "VOR"},
			wantHourWords: []string {"ZEHN"},
			fields: emptyFields,
		},
		{
			name: "Test 50 past 10",
			args: args {
				time: time.Date(2018, 1,1,10,52,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "ZEHN", "VOR"},
			wantHourWords: []string {"ELF"},
			fields: emptyFields,
		},
		{
			name: "Test 55 past 11",
			args: args {
				time: time.Date(2018, 1,1,11,58,0,0, time.UTC),
			},
			wantControlWords: []string{"ES", "IST", "FÜNF", "VOR"},
			wantHourWords: []string {"ZWÖLF"},
			fields: emptyFields,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disp := &LedDisplayModel{
				nonHourWords: tt.fields.nonHourWords,
				hourWords:    tt.fields.hourWords,
			}

			controlWords, hourWords := disp.translateTimeToWords(tt.args.time)
			sort.Strings(controlWords)
			sort.Strings(hourWords)
			sort.Strings(tt.wantHourWords)
			sort.Strings(tt.wantControlWords)

			if !reflect.DeepEqual(controlWords, tt.wantControlWords) {
				t.Errorf("LedDisplayModel.translateTimeToWords() Control words not equal: (Case:%s) = %v, want %v", tt.name, controlWords, tt.wantControlWords)
			}
			if !reflect.DeepEqual(hourWords, tt.wantHourWords) {
				t.Errorf("LedDisplayModel.translateTimeToWords() Hour words not equal: (Case:%s) = %v, want %v", tt.name, hourWords, tt.wantHourWords)
			}
		})
	}
}

func TestLedDisplayModel_getNonHourWord(t *testing.T) {
	type args struct {
		word string
	}

	tests := []struct {
		name   string
		fields testFields
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
			disp := &LedDisplayModel{
				nonHourWords: tt.fields.nonHourWords,
				hourWords:    tt.fields.hourWords,
			}
			if got := disp.getNonHourWord(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LedDisplayModel.getNonHourWord() (Case:%s) = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}
