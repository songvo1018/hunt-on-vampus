package labirinthMap

import (
	"labirynth/game/entity"
	"labirynth/game/items"
	"reflect"
	"testing"
)

func TestLabyrinthEnvironment_changePosition(t *testing.T) {
	type fields struct {
		labyrinthProperties     MapProperties
		userLabyrinthLocation   MapLocation
		npcLabyrinthLocation    MapLocation
		itemLabyrinthProperties items.ItemLabyrinthProperties
	}
	type args struct {
		target      string
		newLocation MapLocation
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "go to empty place",
			fields: fields{
				labyrinthProperties: MapProperties{
					mapLongitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					},
					mapLatitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					},
				},
				userLabyrinthLocation:   MapLocation{1, 1},
				npcLabyrinthLocation:    MapLocation{5, 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target:      "user",
				newLocation: MapLocation{3, 4}},
			want:    true,
			wantErr: false,
		},
		{
			name: "go to busy place",
			fields: fields{
				labyrinthProperties: MapProperties{
					mapLongitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					},
					mapLatitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					}},
				userLabyrinthLocation:   MapLocation{1, 1},
				npcLabyrinthLocation:    MapLocation{5, 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target:      "user",
				newLocation: MapLocation{1, 1}},
			want:    false,
			wantErr: true,
		},
		{
			name: "go to half busy place",
			fields: fields{
				labyrinthProperties: MapProperties{
					mapLongitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Wall,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					},
					mapLatitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					}},
				userLabyrinthLocation:   MapLocation{1, 1},
				npcLabyrinthLocation:    MapLocation{5, 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target:      "user",
				newLocation: MapLocation{2, 2}},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := MapEnvironment{
				labyrinthProperties:     tt.fields.labyrinthProperties,
				userLabyrinthLocation:   tt.fields.userLabyrinthLocation,
				npcLabyrinthLocation:    tt.fields.npcLabyrinthLocation,
				itemLabyrinthProperties: tt.fields.itemLabyrinthProperties,
			}
			changedLocation, err := l.ChangePosition(tt.args.target, tt.args.newLocation)
			if (err != nil) != tt.wantErr {
				t.Errorf("changePosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			newPosition, _ := l.GetPosition(tt.args.target)

			got := changedLocation == tt.args.newLocation && changedLocation == newPosition

			if got != tt.want {
				t.Errorf("changePosition() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLabyrinthEnvironment_getPosition1(t *testing.T) {
	type fields struct {
		labyrinthProperties     MapProperties
		userLabyrinthLocation   MapLocation
		npcLabyrinthLocation    MapLocation
		itemLabyrinthProperties items.ItemLabyrinthProperties
	}
	type args struct {
		target string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    MapLocation
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get user position",
			fields: fields{
				labyrinthProperties: MapProperties{
					mapLongitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					},
					mapLatitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					}},
				userLabyrinthLocation:   MapLocation{1, 1},
				npcLabyrinthLocation:    MapLocation{5, 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target: "user",
			},
			want:    MapLocation{1, 1},
			wantErr: false,
		},
		{
			name: "get nps position",
			fields: fields{
				labyrinthProperties: MapProperties{
					mapLongitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					},
					mapLatitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					}},
				userLabyrinthLocation:   MapLocation{1, 1},
				npcLabyrinthLocation:    MapLocation{5, 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target: "npc",
			},
			want:    MapLocation{5, 5},
			wantErr: false,
		},
		{
			name: "pseudovampus position err",
			fields: fields{
				labyrinthProperties: MapProperties{
					mapLongitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					},
					mapLatitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					}},
				userLabyrinthLocation:   MapLocation{1, 1},
				npcLabyrinthLocation:    MapLocation{5, 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target: "pseudovampus",
			},
			want:    MapLocation{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := MapEnvironment{
				labyrinthProperties:     tt.fields.labyrinthProperties,
				userLabyrinthLocation:   tt.fields.userLabyrinthLocation,
				npcLabyrinthLocation:    tt.fields.npcLabyrinthLocation,
				itemLabyrinthProperties: tt.fields.itemLabyrinthProperties,
			}
			got, err := l.GetPosition(tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPosition() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapEnvironment_lookAround(t *testing.T) {
	type fields struct {
		labyrinthProperties     MapProperties
		userLabyrinthLocation   MapLocation
		npcLabyrinthLocation    MapLocation
		itemLabyrinthProperties items.ItemLabyrinthProperties
	}
	type args struct {
		location MapLocation
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MapAroundEntities
	}{
		// TODO: Add test cases.
		{
			name: "around lookup: wall",
			fields: fields{
				labyrinthProperties: MapProperties{
					mapLongitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					},
					mapLatitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					}},
				userLabyrinthLocation:   MapLocation{1, 1},
				npcLabyrinthLocation:    MapLocation{5, 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{location: MapLocation{3, 3}},
			want: MapAroundEntities{
				Position: MapLocation{3, 3},
				nEntity:  entity.Pass,
				neEntity: entity.Pass,
				nwEntity: entity.Pass,
				eEntity:  entity.Pass,
				sEntity:  entity.Pass,
				seEntity: entity.Pass,
				swEntity: entity.Pass,
				wEntity:  entity.Pass,
			},
		},
		{
			name: "around lookup: nw & n is wall",
			fields: fields{
				labyrinthProperties: MapProperties{
					mapLongitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Wall,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					},
					mapLatitude: map[int]entity.Entity{
						1: entity.Wall,
						2: entity.Pass,
						3: entity.Pass,
						4: entity.Pass,
						5: entity.Pass,
					}},
				userLabyrinthLocation:   MapLocation{4, 4},
				npcLabyrinthLocation:    MapLocation{5, 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{location: MapLocation{2, 2}},
			want: MapAroundEntities{
				Position: MapLocation{2, 2},
				nEntity:  entity.Wall,
				neEntity: entity.Pass,
				eEntity:  entity.Pass,
				seEntity: entity.Pass,
				sEntity:  entity.Pass,
				swEntity: entity.Pass,
				wEntity:  entity.Pass,
				nwEntity: entity.Wall,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MapEnvironment{
				labyrinthProperties:     tt.fields.labyrinthProperties,
				userLabyrinthLocation:   tt.fields.userLabyrinthLocation,
				npcLabyrinthLocation:    tt.fields.npcLabyrinthLocation,
				itemLabyrinthProperties: tt.fields.itemLabyrinthProperties,
			}
			if got := m.lookAround(tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lookAround() = %v, want %v", got, tt.want)
			}
		})
	}
}
