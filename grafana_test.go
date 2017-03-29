package sdk_test

import (
	"encoding/json"
	"fmt"

	"github.com/grafana-tools/sdk"
)

func ExampleNewBoard() {
	board := sdk.NewBoard("Sample dashboard title")
	board.ID = 1
	row1 := board.AddRow("Sample row title")
	row1.Add(sdk.NewGraph("Sample graph"))
	graphWithDs := sdk.NewGraph("Sample graph 2")
	target := sdk.Target{
		RefID:      "A",
		Datasource: "Sample Source 1",
		Expr:       "sample request 1"}
	graphWithDs.AddTarget(&target)
	row1.Add(graphWithDs)
	data, _ := json.MarshalIndent(board, "", "  ")
	fmt.Printf("%s", data)
	// Output:
	// {
	//   "id": 1,
	//   "slug": "",
	//   "title": "Sample dashboard title",
	//   "originalTitle": "",
	//   "tags": null,
	//   "style": "dark",
	//   "timezone": "browser",
	//   "editable": true,
	//   "hideControls": false,
	//   "sharedCrosshair": false,
	//   "rows": [
	//     {
	//       "title": "Sample row title",
	//       "showTitle": false,
	//       "collapse": false,
	//       "editable": true,
	//       "height": "250px",
	//       "panels": [
	//         {
	//           "id": 1,
	//           "title": "Sample graph",
	//           "span": 12,
	//           "renderer": "flot",
	//           "transparent": false,
	//           "type": "graph",
	//           "error": false,
	//           "isNew": true,
	//           "editable": false,
	//           "aliasColors": null,
	//           "bars": false,
	//           "fill": 0,
	//           "legend": {
	//             "alignAsTable": false,
	//             "avg": false,
	//             "current": false,
	//             "hideEmpty": false,
	//             "hideZero": false,
	//             "max": false,
	//             "min": false,
	//             "rightSide": false,
	//             "show": false,
	//             "total": false,
	//             "values": false
	//           },
	//           "lines": false,
	//           "linewidth": 0,
	//           "nullPointMode": "connected",
	//           "percentage": false,
	//           "pointradius": 5,
	//           "points": false,
	//           "stack": false,
	//           "steppedLine": false,
	//           "tooltip": {
	//             "shared": false,
	//             "value_type": ""
	//           },
	//           "x-axis": true,
	//           "y-axis": true,
	//           "xaxis": {
	//             "format": "",
	//             "logBase": 0,
	//             "show": false
	//           },
	//           "yaxes": null
	//         },
	//         {
	//           "id": 2,
	//           "title": "Sample graph 2",
	//           "span": 12,
	//           "renderer": "flot",
	//           "transparent": false,
	//           "type": "graph",
	//           "error": false,
	//           "isNew": true,
	//           "editable": false,
	//           "aliasColors": null,
	//           "bars": false,
	//           "fill": 0,
	//           "legend": {
	//             "alignAsTable": false,
	//             "avg": false,
	//             "current": false,
	//             "hideEmpty": false,
	//             "hideZero": false,
	//             "max": false,
	//             "min": false,
	//             "rightSide": false,
	//             "show": false,
	//             "total": false,
	//             "values": false
	//           },
	//           "lines": false,
	//           "linewidth": 0,
	//           "nullPointMode": "connected",
	//           "percentage": false,
	//           "pointradius": 5,
	//           "points": false,
	//           "stack": false,
	//           "steppedLine": false,
	//           "targets": [
	//             {
	//               "refId": "A",
	//               "datasource": "Sample Source 1",
	//               "expr": "sample request 1"
	//             }
	//           ],
	//           "tooltip": {
	//             "shared": false,
	//             "value_type": ""
	//           },
	//           "x-axis": true,
	//           "y-axis": true,
	//           "xaxis": {
	//             "format": "",
	//             "logBase": 0,
	//             "show": false
	//           },
	//           "yaxes": null
	//         }
	//       ]
	//     }
	//   ],
	//   "templating": {
	//     "list": null
	//   },
	//   "annotations": {
	//     "list": null
	//   },
	//   "schemaVersion": 0,
	//   "version": 0,
	//   "links": null,
	//   "time": {
	//     "from": "",
	//     "to": ""
	//   },
	//   "timepicker": {
	//     "refresh_intervals": null,
	//     "time_options": null
	//   }
	// }
}
