package influxdb_test

import (
	"encoding/json"
	"testing"

	"github.com/influxdata/influxdb"
	platformtesting "github.com/influxdata/influxdb/testing"
)

func TestCheck_MarshalJSON(t *testing.T) {
	type args struct {
		check influxdb.Check
	}
	type wants struct {
		json string
	}
	tests := []struct {
		name  string
		args  args
		wants wants
	}{
		{
			args: args{
				check: influxdb.Check{
					ID: platformtesting.MustIDBase16("f01dab1ef005ba11"),
					// Properties: platform.XYViewProperties{
					//   Type: "xy",
					// },
				},
			},
			wants: wants{
				json: `
{
  "id": "f01dab1ef005ba11",
  "name": "hello",
  "properties": {
    "shape": "chronograf-v2",
    "queries": null,
    "axes": null,
    "type": "xy",
    "colors": null,
    "legend": {},
    "geom": "",
    "note": "",
    "showNoteWhenEmpty": false,
    "xColumn": "",
    "yColumn": "",
    "shadeBelow": false
  }
}
`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.MarshalIndent(tt.args.check, "", "  ")
			if err != nil {
				t.Fatalf("error marshalling json")
			}

			eq, err := jsonEqual(string(b), tt.wants.json)
			if err != nil {
				t.Fatalf("error marshalling json %v", err)
			}
			if !eq {
				t.Errorf("JSON did not match\nexpected:%s\ngot:\n%s\n", tt.wants.json, string(b))
			}
		})
	}
}
