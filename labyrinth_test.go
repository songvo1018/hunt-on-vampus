package game

import (
	"reflect"
	"testing"
)

func TestLabyrinthEnvironment_changePosition(t *testing.T) {
	type fields struct {
		labyrinthProperties     LabyrinthProperties
		userLabyrinthLocation   LabyrinthMapLocation
		npcLabyrinthLocation    LabyrinthMapLocation
		itemLabyrinthProperties ItemLabyrinthProperties
	}
	type args struct {
		target      string
		newLocation LabyrinthMapLocation
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
				labyrinthProperties: LabyrinthProperties{
					mapLongitude: map[int]bool{
						1: true,
						2: false,
						3: false,
						4: false,
						5: false,
					},
					mapLatitude: map[int]bool{
						1: true,
						2: false,
						3: false,
						4: false,
						5: false,
					},
				},
				userLabyrinthLocation:   LabyrinthMapLocation{1, 1},
				npcLabyrinthLocation:    LabyrinthMapLocation{5, 5},
				itemLabyrinthProperties: ItemLabyrinthProperties{items: nil},
			},
			args: args{
				target:      "user",
				newLocation: LabyrinthMapLocation{3, 4}},
			want:    true,
			wantErr: false,
		},
		{
			name: "go to busy place",
			fields: fields{
				labyrinthProperties: LabyrinthProperties{
					mapLongitude: map[int]bool{
						1: true,
						2: false,
						3: false,
						4: false,
						5: false,
					},
					mapLatitude: map[int]bool{
						1: true,
						2: false,
						3: false,
						4: false,
						5: false,
					}},
				userLabyrinthLocation:   LabyrinthMapLocation{1, 1},
				npcLabyrinthLocation:    LabyrinthMapLocation{5, 5},
				itemLabyrinthProperties: ItemLabyrinthProperties{items: nil},
			},
			args: args{
				target:      "user",
				newLocation: LabyrinthMapLocation{1, 1}},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LabyrinthEnvironment{
				labyrinthProperties:     tt.fields.labyrinthProperties,
				userLabyrinthLocation:   tt.fields.userLabyrinthLocation,
				npcLabyrinthLocation:    tt.fields.npcLabyrinthLocation,
				itemLabyrinthProperties: tt.fields.itemLabyrinthProperties,
			}
			changedLocation, err := l.changePosition(tt.args.target, tt.args.newLocation)
			if (err != nil) != tt.wantErr {
				t.Errorf("changePosition() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			newPosition, _ := l.getPosition(tt.args.target)

			got := changedLocation == tt.args.newLocation && changedLocation == newPosition

			if got != tt.want {
				t.Errorf("changePosition() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLabyrinthEnvironment_getPosition1(t *testing.T) {
	type fields struct {
		labyrinthProperties     LabyrinthProperties
		userLabyrinthLocation   LabyrinthMapLocation
		npcLabyrinthLocation    LabyrinthMapLocation
		itemLabyrinthProperties ItemLabyrinthProperties
	}
	type args struct {
		target string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    LabyrinthMapLocation
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "get user position",
			fields: fields{
				labyrinthProperties: LabyrinthProperties{
					mapLongitude: map[int]bool{
						1: true,
						2: false,
						3: false,
						4: false,
						5: false,
					},
					mapLatitude: map[int]bool{
						1: true,
						2: false,
						3: false,
						4: false,
						5: false,
					}},
				userLabyrinthLocation:   LabyrinthMapLocation{1, 1},
				npcLabyrinthLocation:    LabyrinthMapLocation{5, 5},
				itemLabyrinthProperties: ItemLabyrinthProperties{items: nil},
			},
			args: args{
				target: "user",
			},
			want:    LabyrinthMapLocation{1, 1},
			wantErr: false,
		},
		{
			name: "get nps position",
			fields: fields{
				labyrinthProperties: LabyrinthProperties{
					mapLongitude: map[int]bool{
						1: true,
						2: false,
						3: false,
						4: false,
						5: false,
					},
					mapLatitude: map[int]bool{
						1: true,
						2: false,
						3: false,
						4: false,
						5: false,
					}},
				userLabyrinthLocation:   LabyrinthMapLocation{1, 1},
				npcLabyrinthLocation:    LabyrinthMapLocation{5, 5},
				itemLabyrinthProperties: ItemLabyrinthProperties{items: nil},
			},
			args: args{
				target: "npc",
			},
			want:    LabyrinthMapLocation{5, 5},
			wantErr: false,
		},
		{
			name: "pseudovampus position err",
			fields: fields{
				labyrinthProperties: LabyrinthProperties{
					mapLongitude: map[int]bool{
						1: true,
						2: false,
						3: false,
						4: false,
						5: false,
					},
					mapLatitude: map[int]bool{
						1: true,
						2: false,
						3: false,
						4: false,
						5: false,
					}},
				userLabyrinthLocation:   LabyrinthMapLocation{1, 1},
				npcLabyrinthLocation:    LabyrinthMapLocation{5, 5},
				itemLabyrinthProperties: ItemLabyrinthProperties{items: nil},
			},
			args: args{
				target: "pseudovampus",
			},
			want:    LabyrinthMapLocation{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LabyrinthEnvironment{
				labyrinthProperties:     tt.fields.labyrinthProperties,
				userLabyrinthLocation:   tt.fields.userLabyrinthLocation,
				npcLabyrinthLocation:    tt.fields.npcLabyrinthLocation,
				itemLabyrinthProperties: tt.fields.itemLabyrinthProperties,
			}
			got, err := l.getPosition(tt.args.target)
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
