package envdef

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
				UpdateSlice:   UpdateSlice{},
				DeleteSlice:   DeleteSlice{"DELETE=delete"},
				NoChangeSlice: NoChangeSlice{"UPDATE=default", "NOCHANGE=nochange"},
			},
		},
	}

	for _, d := range testdata {
		res, _ := Diff(d.source, d.dist, d.overwrite)

		assert.ElementsMatch(t, d.expected.InsertSlice, res.InsertSlice)
		assert.ElementsMatch(t, d.expected.UpdateSlice, res.UpdateSlice)
		assert.ElementsMatch(t, d.expected.DeleteSlice, res.DeleteSlice)
		assert.ElementsMatch(t, d.expected.NoChangeSlice, res.NoChangeSlice)
	}
}
