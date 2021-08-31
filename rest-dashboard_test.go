package sdk_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"github.com/grafana-tools/sdk"
	"github.com/stretchr/testify/assert"
)

func TestClient_SearchDashboards(t *testing.T) {
	type in struct {
		Query   string
		Starred bool
		Tags    []string
	}
	type testCase struct {
		In  in
		Out url.Values
	}
	dashboard := string(sdk.SearchTypeDashboard)
	for i, tc := range []testCase{
		{
			In:  in{},
			Out: url.Values{"type": []string{dashboard}, "starred": []string{"false"}},
		},
		{
			In:  in{Starred: true},
			Out: url.Values{"type": []string{dashboard}, "starred": []string{"true"}},
		},
		{
			In:  in{Query: "Foo"},
			Out: url.Values{"type": []string{dashboard}, "query": []string{"Foo"}, "starred": []string{"false"}},
		},
		{
			In:  in{Tags: []string{"Foo", "Bar"}},
			Out: url.Values{"type": []string{dashboard}, "starred": []string{"false"}, "tag": []string{"Foo", "Bar"}},
		},
	} {
		ts := httptest.NewServer(http.HandlerFunc(testSearchQuery(t, i, tc.Out)))
		client, _ := sdk.NewClient(ts.URL, "", ts.Client())
		ctx := context.Background()
		_, err := client.SearchDashboards(ctx, tc.In.Query, tc.In.Starred, tc.In.Tags...)
		ts.Close()
		if err != nil {
			t.Fatalf("SearchDashboards test %d failed: %s", i, err)
		}
	}
}

func TestClient_Search(t *testing.T) {
	type testCase struct {
		In  []sdk.SearchParam
		Out url.Values
	}
	for i, tc := range []testCase{
		{
			In:  []sdk.SearchParam{},
			Out: url.Values{},
		},
		{
			// Test all options given their correct usage.
			In: []sdk.SearchParam{
				sdk.SearchDashboardID(234),
				sdk.SearchDashboardID(432),
				sdk.SearchFolderID(123),
				sdk.SearchFolderID(321),
				sdk.SearchLimit(10),
				sdk.SearchPage(99),
				sdk.SearchQuery("Q"),
				sdk.SearchStarred(true),
				sdk.SearchTag("Foo"),
				sdk.SearchTag("Bar"),
				sdk.SearchType(sdk.SearchTypeFolder),
			},
			Out: url.Values{
				"dashboardIds": []string{"234", "432"},
				"folderIds":    []string{"123", "321"},
				"limit":        []string{"10"},
				"page":         []string{"99"},
				"query":        []string{"Q"},
				"starred":      []string{"true"},
				"tag":          []string{"Foo", "Bar"},
				"type":         []string{string(sdk.SearchTypeFolder)},
			},
		},
		{
			// Test non-repeatable options.
			In: []sdk.SearchParam{
				sdk.SearchLimit(10),
				sdk.SearchLimit(100),
				sdk.SearchPage(88),
				sdk.SearchPage(99),
				sdk.SearchQuery("Q1"),
				sdk.SearchQuery("Q2"),
				sdk.SearchStarred(true),
				sdk.SearchStarred(false),
				sdk.SearchType(sdk.SearchTypeFolder),
				sdk.SearchType(sdk.SearchTypeDashboard),
			},
			Out: url.Values{
				"limit":   []string{"100"},
				"page":    []string{"99"},
				"query":   []string{"Q2"},
				"starred": []string{"false"},
				"type":    []string{string(sdk.SearchTypeDashboard)},
			},
		},
	} {
		ts := httptest.NewServer(http.HandlerFunc(testSearchQuery(t, i, tc.Out)))
		client, _ := sdk.NewClient(ts.URL, "", ts.Client())
		ctx := context.Background()
		_, err := client.Search(ctx, tc.In...)
		ts.Close()
		if err != nil {
			t.Fatalf("SearchDashboards test %d failed: %s", i, err)
		}

	}
}

func testSearchQuery(t *testing.T, testID int, exp url.Values) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if m := r.Method; m != http.MethodGet {
			t.Fatalf("unexpected http method for case %d: expected %s, got %s", testID, http.MethodGet, m)
		}
		if e := "/api/search"; r.URL.Path != e {
			t.Fatalf("unexpected http handler called for case %d: expected %s, got %s", testID, r.URL.Path, e)
		}
		qv, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			t.Fatalf("failed to parse query for case %d: %s, got err %s", testID, r.URL.RawQuery, err)
		}
		if !reflect.DeepEqual(exp, qv) {
			t.Fatalf("unexpected query arguments for case %d: expected %v, got %v", testID, exp, qv)
		}
		if _, err := w.Write([]byte("[]")); err != nil {
			t.Fatalf("failed to write http answer for case %d: %s", testID, err)
		}
	}
}

func TestSearchQuery(t *testing.T) {
	testStringSearchParam(t, sdk.SearchQuery, "query", []string{"foo", "bar"})
}

func TestSearchTag(t *testing.T) {
	testRepeatableStringSearchParam(t, sdk.SearchTag, "tag", []string{"foo", "bar"})
}

func TestSearchDashboardID(t *testing.T) {
	testRepeatableIntSearchParam(t, sdk.SearchDashboardID, "dashboardIds", []int{100, 200})
}

func TestSearchFolderID(t *testing.T) {
	testRepeatableIntSearchParam(t, sdk.SearchFolderID, "folderIds", []int{100, 200})
}

func TestSearchPage(t *testing.T) {
	testNonZeroUIntSearchParam(t, sdk.SearchPage, "page", []uint{100, 200})
}

func TestSearchLimit(t *testing.T) {
	testNonZeroUIntSearchParam(t, sdk.SearchLimit, "limit", []uint{100, 200})
}

func TestSearchStarred(t *testing.T) {
	testBoolSearchParam(t, sdk.SearchStarred, "starred", []bool{true, false})
}

func TestSearchType(t *testing.T) {
	var (
		sp         = sdk.SearchType
		key        = "type"
		testValues = []sdk.SearchParamType{sdk.SearchTypeFolder, sdk.SearchTypeDashboard}
	)

	v := make(url.Values)
	for _, testValue := range testValues {
		sp(testValue)(&v)
		expectedLen := 1
		gotLen := len(v[key])
		if gotLen != expectedLen {
			t.Errorf("expected length of %s to be %d, but was %d", key, expectedLen, gotLen)
		}
		value := v.Get(key)
		if value != string(testValue) {
			t.Errorf("expected value of %s to be %s, but was %s", key, testValue, value)
		}
	}
}

func testRepeatableStringSearchParam(t *testing.T, sp func(string) sdk.SearchParam, key string, testValues []string) {
	v := make(url.Values)
	for i, testValue := range testValues {
		sp(testValue)(&v)
		expectedLen := i + 1
		gotLen := len(v[key])
		if gotLen != expectedLen {
			t.Errorf("expected length of %s to be %d, but was %d", key, expectedLen, gotLen)
		}
		last := v[key][i]
		if last != testValue {
			t.Errorf("expected last %s to be %s, but was %s", key, testValue, last)
		}
	}
}

func testStringSearchParam(t *testing.T, sp func(string) sdk.SearchParam, key string, testValues []string) {
	v := make(url.Values)
	for _, testValue := range testValues {
		sp(testValue)(&v)
		expectedLen := 1
		gotLen := len(v[key])
		if gotLen != expectedLen {
			t.Errorf("expected length of %s to be %d, but was %d", key, expectedLen, gotLen)
		}
		value := v.Get(key)
		if value != testValue {
			t.Errorf("expected value of %s to be %s, but was %s", key, testValue, value)
		}
	}
}

func testBoolSearchParam(t *testing.T, sp func(bool) sdk.SearchParam, key string, testValues []bool) {
	v := make(url.Values)
	for _, testValue := range testValues {
		sp(testValue)(&v)
		expectedLen := 1
		gotLen := len(v[key])
		if gotLen != expectedLen {
			t.Errorf("expected length of %s to be %d, but was %d", key, expectedLen, gotLen)
		}
		value := v.Get(key)
		if value != strconv.FormatBool(testValue) {
			t.Errorf("expected value of %s to be %t, but was %s", key, testValue, value)
		}
	}
}

func testRepeatableIntSearchParam(t *testing.T, sp func(int) sdk.SearchParam, key string, testValues []int) {
	v := make(url.Values)
	for i, testValue := range testValues {
		sp(testValue)(&v)
		expectedLen := i + 1
		gotLen := len(v[key])
		if gotLen != expectedLen {
			t.Errorf("expected length of %s to be %d, but was %d", key, expectedLen, gotLen)
		}
		last := v[key][i]
		if last != strconv.Itoa(testValue) {
			t.Errorf("expected last %s to be %d, but was %s", key, testValue, last)
		}
	}
}

func testNonZeroUIntSearchParam(t *testing.T, sp func(uint) sdk.SearchParam, key string, testValues []uint) {
	v := make(url.Values)
	sp(0)(&v)
	value := v.Get(key)
	if value != "" {
		t.Errorf("expected value of %s to be unset, but was %s", key, value)
	}
	testUIntSearchParam(t, sp, key, testValues)
}

func testUIntSearchParam(t *testing.T, sp func(uint) sdk.SearchParam, key string, testValues []uint) {
	v := make(url.Values)
	for _, testValue := range testValues {
		sp(testValue)(&v)
		expectedLen := 1
		gotLen := len(v[key])
		if gotLen != expectedLen {
			t.Errorf("expected length of %s to be %d, but was %d", key, expectedLen, gotLen)
		}
		value := v.Get(key)
		if value != strconv.FormatUint(uint64(testValue), 10) {
			t.Errorf("expected value of %s to be %d, but was %s", key, testValue, value)
		}
	}
}

func Test_DecodeRawBoardPreserveId(t *testing.T) {
	r := getRawBoardRequest(t, true)
	//Serialize object
	data, err := json.Marshal(r)
	assert.Nil(t, err)
	//Deserialize data
	var rawData map[string]interface{} = make(map[string]interface{})
	assert.Nil(t, json.Unmarshal(data, &rawData))
	folderId := rawData["FolderID"]

	assert.Equal(t, folderId.(float64), float64(27))
	assert.Equal(t, rawData["Overwrite"].(bool), true)
	rawData = rawData["dashboard"].(map[string]interface{})
	assert.Equal(t, rawData["id"].(float64), float64(25))
	assert.Equal(t, rawData["title"].(string), "woot")

}

func Test_DecodeRawBoardWipeId(t *testing.T) {
	r := getRawBoardRequest(t, false)
	//Serialize object
	data, err := json.Marshal(r)
	assert.Nil(t, err)
	//Deserialize data
	var rawData map[string]interface{} = make(map[string]interface{})
	err = json.Unmarshal(data, &rawData)
	assert.Nil(t, err)
	folderId := rawData["FolderID"]

	assert.Equal(t, folderId.(float64), float64(27))
	assert.Equal(t, rawData["Overwrite"].(bool), true)
	rawData = rawData["dashboard"].(map[string]interface{})
	assert.Equal(t, rawData["id"].(float64), float64(0))
	assert.Equal(t, rawData["title"].(string), "woot")
}

func getRawBoardRequest(t *testing.T, preserveId bool) sdk.RawBoardRequest {
	r := sdk.RawBoardRequest{}
	board := sdk.Board{}
	board.ID = 25
	board.Title = "woot"
	rawBoard, err := json.Marshal(board)
	assert.Nil(t, err)

	params := sdk.SetDashboardParams{
		FolderID:   27,
		Overwrite:  true,
		PreserveId: preserveId,
	}
	r.Dashboard = rawBoard
	r.Parameters = params

	return r
}

// TestCustomPanelSerialization serializes and de-serializes a sample dashboard
// object a few times and prints out final serialized JSON data to highlight
// that CustomPanel was nested in a strange way.
func TestCustomPanelSerialization(t *testing.T) {
	board := &sdk.Board{
		ID:              0,
		UID:             "",
		Slug:            "",
		Title:           "",
		OriginalTitle:   "",
		Tags:            nil,
		Style:           "",
		Timezone:        "",
		Editable:        false,
		HideControls:    false,
		SharedCrosshair: false,
		Panels: []*sdk.Panel{
			{
				CommonPanel:     sdk.CommonPanel{},
				GraphPanel:      nil,
				TablePanel:      nil,
				TextPanel:       nil,
				SinglestatPanel: nil,
				DashlistPanel:   nil,
				PluginlistPanel: nil,
				RowPanel:        nil,
				AlertlistPanel:  nil,
				CustomPanel: &sdk.CustomPanel{
					"key": "value",
				},
			},
		},
		Rows:       nil,
		Templating: sdk.Templating{},
		Annotations: struct {
			List []sdk.Annotation `json:"list"`
		}{},
		Refresh:       nil,
		SchemaVersion: 0,
		Version:       0,
		Links:         nil,
		Time:          sdk.Time{},
		Timepicker:    sdk.Timepicker{},
		GraphTooltip:  0,
	}

	jb, err := json.Marshal(board)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		newBoard := &sdk.Board{}
		if err := json.Unmarshal(jb, newBoard); err != nil {
			t.Fatal(err)
		}

		jb, err = json.MarshalIndent(newBoard, "", "  ")
		if err != nil {
			t.Fatal(err)
		}
	}

	cnt := strings.Count(string(jb), "CustomPanel")
	if cnt != 0 {
		t.Fatalf("expected to not have any CustomPanel keys, got: %v", cnt)
	}
}
