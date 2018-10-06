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

	tests := []struct {
		name string
		args args
		want Maze
	}{
		{
			name: "Maze with given size is created in given size",
			args: args{width: 1, height: 3},
			want: createMazeFromString(1, 3, "XXX"),
		},
		{
			name: "Maze with just one field gets created",
			args: args{width: 1, height: 1},
			want: createMazeFromString(1, 1, "X"),
		},
		{
			name: "Negative size parameters create zero sized maze",
			args: args{width: -1, height: -1},
			want: createMazeFromString(0, 0, ""),
		},
		{
			name: "Zero as parameter creates a zero sized maze",
			args: args{width: 0, height: 0},
			want: createMazeFromString(0, 0, ""),
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

func TestSet(t *testing.T) {
	type args struct {
		x         int
		y         int
		isPassway bool
	}
	tests := []struct {
		name  string
		given Maze
		args  args
		want  Maze
	}{
		{
			name: "Passway is set correctly on a maze without any passways",
			given: createMazeFromString(2, 3,
				""+
					"XX"+
					"XX"+
					"XX"),
			args: args{x: 2, y: 2, isPassway: true},
			want: createMazeFromString(2, 3,
				""+
					"XX"+
					"X_"+
					"XX"),
		},
		{
			name: "Wall is set correctly on a maze without any walls",
			given: createMazeFromString(2, 3,
				""+
					"__"+
					"__"+
					"__"),
			args: args{x: 2, y: 2, isPassway: false},
			want: createMazeFromString(2, 3,
				""+
					"__"+
					"_X"+
					"__"),
		},
		{
			name: "When a wall is set to be a wall, it stays a wall",
			given: createMazeFromString(2, 3,
				""+
					"XX"+
					"XX"+
					"XX"),
			args: args{x: 2, y: 2, isPassway: false},
			want: createMazeFromString(2, 3,
				""+
					"XX"+
					"XX"+
					"XX"),
		},
		{
			name: "When a passway is set to be a passway, it stays a passway",
			given: createMazeFromString(2, 3,
				""+
					"XX"+
					"X_"+
					"XX"),
			args: args{x: 2, y: 2, isPassway: true},
			want: createMazeFromString(2, 3,
				""+
					"XX"+
					"X_"+
					"XX"),
		}, {
			name: "Field at 0,0 can be set",
			given: createMazeFromString(2, 2,
				""+
					"XX"+
					"XX"),
			args: args{x: 1, y: 1, isPassway: true},
			want: createMazeFromString(2, 2,
				""+
					"_X"+
					"XX"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maze := tt.given
			maze.Set(tt.args.x, tt.args.y, tt.args.isPassway)

		})
	}
}

func createMazeFromString(width, height int, input string) Maze {
	maze := New(width, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			charPosition := y*width + x
			char := input[charPosition]
			if char == '_' {
				maze.Set(x, y, false)
			}
		}
	}
	return maze
}
