package grafana_test

import (
	"encoding/json"
	"fmt"

	"github.com/grafov/autograf/grafana"
)

func ExampleNewBoard() {
	board := grafana.NewBoard("Sample dashboard title")
	board.ID = 1
	row1 := board.AddRow("Sample row title")
	row1.Add(grafana.NewGraph("Sample graph"))
	graphWithDs := grafana.NewGraph("Sample graph 2")
	target := grafana.Target{
		RefID:      "A",
		Datasource: "Sample Source 1",
		Expr:       "sample request 1"}
	graphWithDs.AddTarget(&target)
	row1.Add(graphWithDs)
	data, _ := json.MarshalIndent(board, "", "  ")
	fmt.Printf("%s", data)
	// Output:
	// 	{
	//   "id": 1,
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
	//           "aliasColors": null,
	//           "bars": false,
	//           "fill": 0,
	//           "grid": {
	//             "leftLogBase": null,
	//             "leftMax": null,
	//             "leftMin": null,
	//             "rightLogBase": null,
	//             "rightMax": null,
	//             "rightMin": null,
	//             "threshold1": null,
	//             "threshold1Color": "",
	//             "threshold2": null,
	//             "threshold2Color": "",
	//             "thresholdLine": false
	//           },
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
	//           "seriesOverrides": null,
	//           "stack": false,
	//           "steppedLine": false,
	//           "timeFrom": null,
	//           "timeShift": null,
	//           "tooltip": {
	//             "shared": false,
	//             "value_type": ""
	//           },
	//           "x-axis": true,
	//           "y-axis": true,
	//           "y_formats": null
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
	//           "aliasColors": null,
	//           "bars": false,
	//           "fill": 0,
	//           "grid": {
	//             "leftLogBase": null,
	//             "leftMax": null,
	//             "leftMin": null,
	//             "rightLogBase": null,
	//             "rightMax": null,
	//             "rightMin": null,
	//             "threshold1": null,
	//             "threshold1Color": "",
	//             "threshold2": null,
	//             "threshold2Color": "",
	//             "thresholdLine": false
	//           },
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
	//           "seriesOverrides": null,
	//           "stack": false,
	//           "steppedLine": false,
	//           "targets": [
	//             {
	//               "refId": "A",
	//               "datasource": "Sample Source 1",
	//               "expr": "sample request 1",
	//               "intervalFactor": 0,
	//               "step": 0,
	//               "legendFormat": ""
	//             }
	//           ],
	//           "timeFrom": null,
	//           "timeShift": null,
	//           "tooltip": {
	//             "shared": false,
	//             "value_type": ""
	//           },
	//           "x-axis": true,
	//           "y-axis": true,
	//           "y_formats": null
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
	//   "schemiaVersion": 0,
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
