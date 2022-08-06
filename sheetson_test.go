package sheetson_test

import (
	"testing"

	"github.com/kecci/go-sheetson"
)

func TestAddRow(t *testing.T) {
	type args struct {
		sheetonClient sheetson.Client
		spreadSheetID string
		sheetName     string
		row           map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "AddRow",
			args: args{
				sheetonClient: sheetson.NewClient("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
				spreadSheetID: "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
				sheetName:     "Sheet1",
				row: map[string]interface{}{
					"name":       "San Francisco",
					"state":      "CA",
					"country":    "USA",
					"population": "860000",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.args.sheetonClient.AddRow(tt.args.spreadSheetID, tt.args.sheetName, tt.args.row); (err != nil) != tt.wantErr {
				t.Errorf("AddRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}
