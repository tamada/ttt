package ttt

import "testing"

func TestReadDataFromJsonDataStore(t *testing.T) {
	ds := JSONDataStore{}
	if err := ds.Init(); err != nil {
		t.Error(err)
	}
	if len(ds.Courses()) != 11 {
		t.Errorf("courses size did not match, wont %d, got %d", 11, len(ds.Courses()))
	}
	if len(ds.Lectures()) != 80 {
		t.Errorf("lectures size did not match, wont %d, got %d", 81, len(ds.Lectures()))
	}
}
