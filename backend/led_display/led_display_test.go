package led_display

import (
	"testing"
	"github.com/maylukas/go-wordclock/backend/display_model/model"
	"github.com/maylukas/go-wordclock/backend/display_model/api"
	"golang.org/x/text/language"
)

func TestLedDisplay_mapIndex(t *testing.T) {
	type fields struct {
		model  model.Display
		//strip  ws281x.Matrix
		config Config
	}
	type args struct {
		ledIdx int
	}

	/**
	Test alternating:

	 0  1  2  3  4  5  6  7  8  9 10        0  1  2  3  4  5  6  7  8  9 10
	11 12 13 14 15 16 17 18 19 20 21   =>  21 20 19 18 17 16 15 14 13 12 11
	22 23 ...                              22 23
	 */
	alternateEvenConfig := Config{
		AlternateFlow: true,
		AlternateEven: true,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Alternate even 10",
			fields: fields{
				config: alternateEvenConfig,
				model: api.GetDisplayModelForLanguage(language.German),
			},
			args: args{
				ledIdx: 10,
			},
			want: 10,
		},
		{
			name: "Alternate even 11",
			fields: fields{
				config: alternateEvenConfig,
				model: api.GetDisplayModelForLanguage(language.German),
			},
			args: args{
				ledIdx: 11,
			},
			want: 21,
		},
		{
			name: "Alternate even 11",
			fields: fields{
				config: alternateEvenConfig,
				model: api.GetDisplayModelForLanguage(language.German),
			},
			args: args{
				ledIdx: 21,
			},
			want: 11,
		},
		{
			name: "Alternate even 11",
			fields: fields{
				config: alternateEvenConfig,
				model: api.GetDisplayModelForLanguage(language.German),
			},
			args: args{
				ledIdx: 22,
			},
			want: 22,
		},
		{
			name: "Alternate even 11",
			fields: fields{
				config: alternateEvenConfig,
				model: api.GetDisplayModelForLanguage(language.German),
			},
			args: args{
				ledIdx: 18,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ledDisplayImpl{
				model:  tt.fields.model,
				//strip:  tt.fields.strip,
				config: &tt.fields.config,
			}
			if got := l.mapIndex(tt.args.ledIdx); got != tt.want {
				t.Errorf("LedDisplay.mapIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
