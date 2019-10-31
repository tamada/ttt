package ziraffe

import "testing"

func TestLevenshtein(t *testing.T) {
	testdata := []struct {
		s1       string
		s2       string
		distance int
	}{
		{"distance", "similarity", 8},
		{"android", "ipodtouch", 7},
		{"dog", "cat", 3},
		{"mother", "other", 1},
	}

	for _, td := range testdata {
		distance := LevenshteinS(td.s1, td.s2)
		if distance != td.distance {
			t.Errorf("distance of \"%s\" and \"%s\" wont %d, but got %d", td.s1, td.s2, td.distance, distance)
		}
	}
}
