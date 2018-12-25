package CSVReader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractContents(tt *testing.T) {
	type args struct {
		filePath string
	}

	type want struct {
		lines [][]string
		err   error
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "Successfully read csv",
			args: args{
				filePath: "test_file.csv",
			},
			want: want{
				lines: [][]string{
					[]string{"Column1", "Column2", "Column3"},
					[]string{"1", "name1", "value1"},
					[]string{"2", "name2", "value2"},
					[]string{"3", "name3", "value3"},
					[]string{"4", "name4", "value4"},
					[]string{"5", "name5", "value5"},
					[]string{"6", "name6", "value6"},
					[]string{"7", "name7", "value7"},
					[]string{"8", "name8", "value8"},
					[]string{"9", "name9", "value9"},
					[]string{"10", "name10", "value10"},
				},
				err: nil,
			},
		},
	}

	for _, test := range tests {
		tt.Run(test.name, func(t *testing.T) {
			lines, err := ExtractContents(test.args.filePath)
			assert.Equal(t, test.want.lines, lines)
			assert.Equal(t, test.want.err, err)
		})
	}
}

func BenchmarkExtractContents(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ExtractContents("test_file.csv")
	}
}
