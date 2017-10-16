package main

import (
	"reflect"
	"testing"
)

func Test_csvFirstColumnToSlice(t *testing.T) {
	type args struct {
		csvSlice [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "TestA",
			want: []string{"A1", "A2", "A3"},
			args: args{
				csvSlice: [][]string{
					{"A1", "B1", "C1"},
					{"A2", "B2", "C2"},
					{"A3", "B3", "C3"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := csvFirstColumnToSlice(tt.args.csvSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("csvFirstColumnToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
