package envdef

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	testdata := []struct {
		path     string
		expected map[string]string
	}{
		{
			path: "testdata/.env",
			expected: map[string]string{
				"UPDATE":   "update",
				"DELETE":   "delete",
				"NOCHANGE": "nochange",
			},
		},
		{
			path:     "testdata/notfound",
			expected: nil,
		},
	}

	for idx := range testdata {
		res, _ := Read(testdata[idx].path)
		if len(testdata[idx].expected) != len(res) {
			t.Errorf("expected len(%v), got len(%v)", len(testdata[idx].expected), len(res))
		}

		for k, v := range testdata[idx].expected {
			if v != res[k] {
				t.Errorf("expected %s to be %s, got %s", k, v, res[k])
			}
		}
	}
}

func TestDiff(t *testing.T) {
	testdata := []struct {
		source    string
		dist      string
		overwrite bool
		expected  Result
	}{
		{
			source:    "testdata/.env.sample",
			dist:      "testdata/.env",
			overwrite: true,
			expected: Result{
				InsertSlice:   InsertSlice{"INSERT=insert"},
				UpdateSlice:   UpdateSlice{"UPDATE=default"},
				DeleteSlice:   DeleteSlice{"DELETE=delete"},
				NoChangeSlice: NoChangeSlice{"NOCHANGE=nochange"},
			},
		},
		{
			source:    "testdata/.env.sample",
			dist:      "testdata/.env",
			overwrite: false,
			expected: Result{
				InsertSlice:   InsertSlice{"INSERT=insert"},
				UpdateSlice:   UpdateSlice{"UPDATE=update"},
				DeleteSlice:   DeleteSlice{"DELETE=delete"},
				NoChangeSlice: NoChangeSlice{"NOCHANGE=nochange"},
			},
		},
	}

	for _, d := range testdata {
		res, _ := Diff(d.source, d.dist, d.overwrite)
		// check len

		if len(d.expected.InsertSlice) != len(res.InsertSlice) {
			t.Errorf("expected len(%v), got len(%v)", len(d.expected.InsertSlice), len(res.InsertSlice))
		}

		if len(d.expected.UpdateSlice) != len(res.UpdateSlice) {
			t.Errorf("expected len(%v), got len(%v)", len(d.expected.UpdateSlice), len(res.UpdateSlice))
		}

		if len(d.expected.DeleteSlice) != len(res.DeleteSlice) {
			t.Errorf("expected len(%v), got len(%v)", len(d.expected.DeleteSlice), len(res.DeleteSlice))
		}

		if len(d.expected.NoChangeSlice) != len(res.NoChangeSlice) {
			t.Errorf("expected len(%v), got len(%v)", len(d.expected.NoChangeSlice), len(res.NoChangeSlice))
		}

		if !reflect.DeepEqual(d.expected.InsertSlice, res.InsertSlice) {
			t.Errorf("expected %s, got %s", d.expected.InsertSlice, res.InsertSlice)
		}

		if !reflect.DeepEqual(d.expected.UpdateSlice, res.UpdateSlice) {
			t.Errorf("expected %s, got %s", d.expected.UpdateSlice, res.UpdateSlice)
		}

		if !reflect.DeepEqual(d.expected.DeleteSlice, res.DeleteSlice) {
			t.Errorf("expected %s, got %s", d.expected.DeleteSlice, res.DeleteSlice)
		}

		if !reflect.DeepEqual(d.expected.NoChangeSlice, res.NoChangeSlice) {
			t.Errorf("expected %s, got %s", d.expected.NoChangeSlice, res.NoChangeSlice)
		}
	}
}
