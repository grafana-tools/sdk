package sdk_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strconv"
	"testing"

	"github.com/grafana-tools/sdk"
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
		client := sdk.NewClient(ts.URL, "", ts.Client())
		ctx := context.Background()
		_, err := client.SearchDashboards(ctx, tc.In.Query, tc.In.Starred, tc.In.Tags...)
		ts.Close()
		if err != nil {
			t.Fatalf("SearchDashboards test %d failed: %s", i, err)
		}
	}
}

func TestClient_SearchWithParams(t *testing.T) {
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
				sdk.WithSearchDashboardID(234),
				sdk.WithSearchDashboardID(432),
				sdk.WithSearchFolderID(123),
				sdk.WithSearchFolderID(321),
				sdk.WithSearchLimit(10),
				sdk.WithSearchPage(99),
				sdk.WithSearchQuery("Q"),
				sdk.WithSearchStarred(true),
				sdk.WithSearchTag("Foo"),
				sdk.WithSearchTag("Bar"),
				sdk.WithSearchType(sdk.SearchTypeFolder),
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
				sdk.WithSearchLimit(10),
				sdk.WithSearchLimit(100),
				sdk.WithSearchPage(88),
				sdk.WithSearchPage(99),
				sdk.WithSearchQuery("Q1"),
				sdk.WithSearchQuery("Q2"),
				sdk.WithSearchStarred(true),
				sdk.WithSearchStarred(false),
				sdk.WithSearchType(sdk.SearchTypeFolder),
				sdk.WithSearchType(sdk.SearchTypeDashboard),
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
		client := sdk.NewClient(ts.URL, "", ts.Client())
		ctx := context.Background()
		_, err := client.SearchWithParams(ctx, tc.In...)
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

func TestWithSearchQuery(t *testing.T) {
	testStringSearchParam(t, sdk.WithSearchQuery, "query", []string{"foo", "bar"})
}

func TestWithSearchTag(t *testing.T) {
	testRepeatableStringSearchParam(t, sdk.WithSearchTag, "tag", []string{"foo", "bar"})
}

func TestWithSearchDashboardID(t *testing.T) {
	testRepeatableIntSearchParam(t, sdk.WithSearchDashboardID, "dashboardIds", []int{100, 200})
}

func TestWithSearchFolderID(t *testing.T) {
	testRepeatableIntSearchParam(t, sdk.WithSearchFolderID, "folderIds", []int{100, 200})
}

func TestWithSearchPage(t *testing.T) {
	testNonZeroUIntSearchParam(t, sdk.WithSearchPage, "page", []uint{100, 200})
}

func TestWithSearchLimit(t *testing.T) {
	testNonZeroUIntSearchParam(t, sdk.WithSearchLimit, "limit", []uint{100, 200})
}

func TestWithSearchStarred(t *testing.T) {
	testBoolSearchParam(t, sdk.WithSearchStarred, "starred", []bool{true, false})
}

func TestWithSearchType(t *testing.T) {
	var (
		sp         = sdk.WithSearchType
		key        = "type"
		testValues = []sdk.SearchType{sdk.SearchTypeFolder, sdk.SearchTypeDashboard}
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
