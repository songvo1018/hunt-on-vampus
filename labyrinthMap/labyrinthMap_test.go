package labirinthMap

import (
	"labirynth/game/entity"
	"labirynth/game/items"
	"labirynth/game/location"
	"reflect"
	"testing"
)

func TestLabyrinthEnvironment_changePosition(t *testing.T) {
	type fields struct {
		labyrinthProperties     MapProperties
		userLabyrinthLocation   location.MapLocation
		npcLabyrinthLocation    location.MapLocation
		itemLabyrinthProperties items.ItemLabyrinthProperties
	}
	type args struct {
		target      string
		newLocation location.MapLocation
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
					GridSize:      25,
					EntitiesOnMap: map[location.MapLocation]entity.Entity{},
				},
				userLabyrinthLocation:   location.MapLocation{X: 1, Y: 1},
				npcLabyrinthLocation:    location.MapLocation{X: 5, Y: 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target:      "user",
				newLocation: location.MapLocation{X: 3, Y: 4}},
			want:    true,
			wantErr: false,
		},
		{
			name: "go to busy place",
			fields: fields{
				labyrinthProperties: MapProperties{
					GridSize: 25,
					EntitiesOnMap: map[location.MapLocation]entity.Entity{
						location.MapLocation{X: 1, Y: 1}: entity.WallEntity,
					}},
				userLabyrinthLocation:   location.MapLocation{X: 1, Y: 1},
				npcLabyrinthLocation:    location.MapLocation{X: 5, Y: 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target:      "user",
				newLocation: location.MapLocation{X: 1, Y: 1}},
			want:    false,
			wantErr: true,
		},
		{
			name: "go to half busy place",
			fields: fields{
				labyrinthProperties: MapProperties{
					GridSize:      25,
					EntitiesOnMap: map[location.MapLocation]entity.Entity{}},
				userLabyrinthLocation:   location.MapLocation{X: 1, Y: 1},
				npcLabyrinthLocation:    location.MapLocation{X: 5, Y: 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target:      "user",
				newLocation: location.MapLocation{X: 2, Y: 2}},
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

			l.labyrinthProperties.fillEmptyLocation()
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
		userLabyrinthLocation   location.MapLocation
		npcLabyrinthLocation    location.MapLocation
		itemLabyrinthProperties items.ItemLabyrinthProperties
	}
	type args struct {
		target string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    location.MapLocation
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get user position",
			fields: fields{
				labyrinthProperties: MapProperties{
					GridSize:      25,
					EntitiesOnMap: map[location.MapLocation]entity.Entity{}},
				userLabyrinthLocation:   location.MapLocation{X: 1, Y: 1},
				npcLabyrinthLocation:    location.MapLocation{X: 5, Y: 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target: "user",
			},
			want:    location.MapLocation{X: 1, Y: 1},
			wantErr: false,
		},
		{
			name: "get nps position",
			fields: fields{
				labyrinthProperties: MapProperties{
					GridSize:      25,
					EntitiesOnMap: map[location.MapLocation]entity.Entity{}},
				userLabyrinthLocation:   location.MapLocation{X: 1, Y: 1},
				npcLabyrinthLocation:    location.MapLocation{X: 5, Y: 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target: "npc",
			},
			want:    location.MapLocation{X: 5, Y: 5},
			wantErr: false,
		},
		{
			name: "pseudovampus position err",
			fields: fields{
				labyrinthProperties: MapProperties{
					GridSize:      25,
					EntitiesOnMap: map[location.MapLocation]entity.Entity{}},
				userLabyrinthLocation:   location.MapLocation{X: 1, Y: 1},
				npcLabyrinthLocation:    location.MapLocation{X: 5, Y: 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{
				target: "pseudovampus",
			},
			want:    location.MapLocation{},
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
		userLabyrinthLocation   location.MapLocation
		npcLabyrinthLocation    location.MapLocation
		itemLabyrinthProperties items.ItemLabyrinthProperties
	}
	type args struct {
		location location.MapLocation
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   MapAroundEntities
	}{
		// TODO: Add test cases.
		{
			name: "around lookup: no wall",
			fields: fields{
				labyrinthProperties: MapProperties{
					GridSize:      25,
					EntitiesOnMap: map[location.MapLocation]entity.Entity{}},
				userLabyrinthLocation:   location.MapLocation{X: 1, Y: 1},
				npcLabyrinthLocation:    location.MapLocation{X: 5, Y: 5},
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{location: location.MapLocation{X: 2, Y: 2}},
			want: MapAroundEntities{
				Position: location.MapLocation{X: 2, Y: 2},
				nEntity:  entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 2, Y: 3}},
				neEntity: entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 3, Y: 3}},
				eEntity:  entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 3, Y: 2}},
				seEntity: entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 3, Y: 1}},
				sEntity:  entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 2, Y: 1}},
				swEntity: entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 1, Y: 1}},
				wEntity:  entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 1, Y: 2}},
				nwEntity: entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 1, Y: 3}},
			},
		},
		{
			name: "around lookup: wall from south ",
			fields: fields{
				labyrinthProperties: MapProperties{
					GridSize: 25,
					EntitiesOnMap: map[location.MapLocation]entity.Entity{
						location.MapLocation{X: 2, Y: 1}: entity.Entity{Class: entity.WallType, IsBarrier: true, Location: location.MapLocation{X: 2, Y: 1}},
					}},
				userLabyrinthLocation: location.MapLocation{X: 2, Y: 2},
				npcLabyrinthLocation:  location.MapLocation{X: 5, Y: 5},
				//todo store entities ids for check barrier or not
				itemLabyrinthProperties: items.ItemLabyrinthProperties{Items: nil},
			},
			args: args{location: location.MapLocation{X: 2, Y: 2}},
			want: MapAroundEntities{
				Position: location.MapLocation{X: 2, Y: 2},
				nEntity:  entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 2, Y: 3}},
				neEntity: entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 3, Y: 3}},
				eEntity:  entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 3, Y: 2}},
				seEntity: entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 3, Y: 1}},
				sEntity:  entity.Entity{Class: entity.WallType, IsBarrier: true, Location: location.MapLocation{X: 2, Y: 1}},
				swEntity: entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 1, Y: 1}},
				wEntity:  entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 1, Y: 2}},
				nwEntity: entity.Entity{Class: entity.PassType, IsBarrier: false, Location: location.MapLocation{X: 1, Y: 3}},
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

			m.labyrinthProperties.fillEmptyLocation()
			if got := m.lookAround(tt.args.location); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lookAround()\n got\n = %v, \n want\n = %v \n\n", got, tt.want)
			}
		})
	}
}

func TestMapProperties_fillEmptyLocation(t *testing.T) {
	type fields struct {
		GridSize      int
		EntitiesOnMap map[location.MapLocation]entity.Entity
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
		{
			name: "fill empty location paths",
			fields: fields{
				GridSize: 25,
				EntitiesOnMap: map[location.MapLocation]entity.Entity{
					location.MapLocation{X: 1, Y: 1}: entity.WallEntity,
					location.MapLocation{X: 1, Y: 2}: entity.WallEntity,
					location.MapLocation{X: 1, Y: 3}: entity.WallEntity,
					location.MapLocation{X: 1, Y: 4}: entity.WallEntity,
					location.MapLocation{X: 1, Y: 5}: entity.WallEntity,
				},
			},
			want: true,
		},
		{
			name: "wall after filling must be wall",
			fields: fields{
				GridSize: 25,
				EntitiesOnMap: map[location.MapLocation]entity.Entity{
					location.MapLocation{X: 1, Y: 1}: entity.Entity{Class: entity.WallType, IsBarrier: true, Location: location.MapLocation{X: 1, Y: 1}},
					location.MapLocation{X: 1, Y: 2}: entity.WallEntity,
					location.MapLocation{X: 1, Y: 3}: entity.WallEntity,
					location.MapLocation{X: 1, Y: 4}: entity.WallEntity,
					location.MapLocation{X: 1, Y: 5}: entity.WallEntity,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &MapProperties{
				GridSize:      tt.fields.GridSize,
				EntitiesOnMap: tt.fields.EntitiesOnMap,
			}

			p.fillEmptyLocation()
			if tt.want && p.GridSize*p.GridSize != len(p.EntitiesOnMap) {
				t.Errorf("%v; got %v, want %v", tt.name, p.GridSize, len(p.EntitiesOnMap))
			} else if !tt.want && p.EntitiesOnMap[location.MapLocation{X: 1, Y: 1}].Class != entity.WallType {
				t.Errorf("%v; got %v, want %v", tt.name, p.GridSize, len(p.EntitiesOnMap))
			}
		})
	}
}
