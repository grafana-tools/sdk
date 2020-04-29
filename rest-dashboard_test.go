package sdk_test

import (
	"net/url"
	"strconv"
	"testing"

	"github.com/grafana-tools/sdk"
)

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
		testValues = []sdk.SearchType{"foo", "bar"}
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
