package maze

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		width  int
		height int
	}

	oneByOne := createStructure(1, 1)
	emptyStructure := [][]bool{}

	tests := []struct {
		name string
		args args
		want Maze
	}{
		{
			name: "Negative size creates zero sized maze",
			args: args{width: 1, height: 1},
			want: Maze{width: 1, height: 1, structure: oneByOne},
		},
		{
			name: "Negative size creates zero sized maze",
			args: args{width: -1, height: -1},
			want: Maze{width: 0, height: 0, structure: emptyStructure},
		},
		{
			name: "Zero size creates zero sized maze",
			args: args{width: 0, height: 0},
			want: Maze{width: 0, height: 0, structure: emptyStructure},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.width, tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createStructure(width int, height int) [][]bool {
	result := make([][]bool, height)
	for i := range result {
		result[i] = make([]bool, width)
	}
	return result
}
