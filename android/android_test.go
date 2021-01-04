package android

import (
	"fmt"
	"reaktor/input"
	"reflect"
	"strings"
	"testing"
)

func TestLink_FindPath(t *testing.T) {
	testLink := NewLink()
	fractions := strings.Split(input.ReadInput("./input.txt"), "\n")
	for _, fraction := range fractions {
		testLink.readFraction(fraction)
	}
	type fields struct {
		mapping map[Coord]bool
		start   Coord
		end     map[Coord]bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Testing with the puzzle input",
			fields: fields{
				mapping: testLink.mapping,
				start:   testLink.start,
				end:     testLink.end},
			want: input.ReadInput("./solution.txt"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("Link Start: ")

			link := Link{
				mapping: tt.fields.mapping,
				start:   tt.fields.start,
				end:     tt.fields.end,
			}
			fmt.Println("Link: ", link)
			if got := link.FindPath(); got != tt.want {
				t.Errorf("Link.FindPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoord_getDirectionTo(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		nextCoord Coord
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"Right", fields{1, 1}, args{Coord{2, 1}}, "R"},
		{"Left", fields{1, 1}, args{Coord{0, 1}}, "L"},
		{"Up", fields{1, 1}, args{Coord{1, 0}}, "U"},
		{"Down", fields{1, 1}, args{Coord{1, 2}}, "D"},
		{"Same Coordinate", fields{1, 1}, args{Coord{1, 1}}, "FAILED"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coord := Coord{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := coord.getDirectionTo(tt.args.nextCoord); got != tt.want {
				t.Errorf("Coord.getDirectionTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLink_readFraction(t *testing.T) {
	type fields struct {
		mapping map[Coord]bool
		start   Coord
		end     map[Coord]bool
	}
	type args struct {
		fraction string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Link
	}{
		{
			name: "Down and Right to the Start. No previous data has been added to the link",
			fields: fields{
				mapping: make(map[Coord]bool),
				end:     make(map[Coord]bool),
			},
			args: args{"0,0 D,D,R,S"},
			want: Link{
				mapping: map[Coord]bool{{0, 0}: true, {0, 1}: true, {0, 2}: true, {1, 2}: true},
				start:   Coord{1, 2},
				end:     make(map[Coord]bool)},
		},
		{
			name: "Up and Left to the End. Link has some coordinates and an end set",
			fields: fields{
				mapping: map[Coord]bool{{0, 5}: true, {2, 2}: true},
				end:     map[Coord]bool{{0, 5}: true}},
			args: args{"3,2 L,U,L,F"},
			want: Link{
				mapping: map[Coord]bool{{0, 5}: true, {3, 2}: true, {2, 2}: true, {2, 1}: true, {1, 1}: true},
				start:   Coord{0, 0},
				end:     map[Coord]bool{{0, 5}: true, {1, 1}: true}},
		},
		{
			name: "Path ends without directions",
			fields: fields{
				mapping: map[Coord]bool{{3, 2}: true},
				end:     map[Coord]bool{{0, 5}: true}},
			args: args{"0,0"},
			want: Link{
				mapping: map[Coord]bool{{3, 2}: true, {0, 0}: true},
				start:   Coord{0, 0},
				end:     map[Coord]bool{{0, 5}: true}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			link := Link{
				mapping: tt.fields.mapping,
				start:   tt.fields.start,
				end:     tt.fields.end,
			}
			link.readFraction(tt.args.fraction)
			if !reflect.DeepEqual(link, tt.want) {
				t.Errorf("readFraction() = %v, want %v", link, tt.want)
			}
		})
	}
}
