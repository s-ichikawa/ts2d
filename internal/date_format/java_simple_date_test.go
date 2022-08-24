package date_format_test

import (
	"reflect"
	"testing"

	df "github.com/s-ichikawa/ts2d/internal/date_format"
)

func TestJavaSimpleDateFormatterLayout(t *testing.T) {
	tests := []struct {
		name, ptn, expected string
	}{
		{
			name:     "simple format",
			ptn:      "yyyyMMdd HH:mm:ss",
			expected: "20060102 15:04:05",
		},
		{
			name:     "simple format with AM/PM & milliseconds",
			ptn:      "yyyyMMdd h:mm:ss.SSS a",
			expected: "20060102 3:04:05.000 PM",
		},
		{
			name:     "RFC3339",
			ptn:      "yyyy-MM-dd'T'HH:mm:ssXXX",
			expected: "2006-01-02T15:04:05-07:00",
		},
		{
			name:     "UnixDate",
			ptn:      "EEE MMM d HH:mm:ss z yyyy",
			expected: "Mon Jan _2 15:04:05 MST 2006",
		},
		{
			name:     "RFC822",
			ptn:      "dd MMM yy HH:mm z",
			expected: "02 Jan 06 15:04 MST",
		},
		{
			name:     "RFC850",
			ptn:      "EEEE, dd-MMM-yy HH:mm:ss z",
			expected: "Monday, 02-Jan-06 15:04:05 MST",
		},
		{
			name:     "RFC1123",
			ptn:      "EEE, dd MMM yyyy HH:mm:ss z",
			expected: "Mon, 02 Jan 2006 15:04:05 MST",
		},
		{
			name:     "With single-quoted words",
			ptn:      "'at' hh 'o''clock' a, z",
			expected: "at 03 o'clock PM, MST",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			f := df.NewJavaSimpleDateConvertor(tt.ptn)
			got := f.ToLayout()
			if !reflect.DeepEqual(tt.expected, got) {
				t.Fatalf("expected: %v, got: %v", tt.expected, got)
			}
		})
	}
}
