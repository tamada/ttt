package ttt

import "testing"

func TestReadDataOfStandaloneDataStore(t *testing.T) {
	ds := NewStandaloneDataStore()
	if ds == nil {
		t.Error("initialize error")
	}
	if len(ds.Courses()) != 11 {
		t.Errorf("courses size did not match, wont %d, got %d", 11, len(ds.Courses()))
	}
	if len(ds.Lectures()) != 80 {
		t.Errorf("lectures size did not match, wont %d, got %d", 81, len(ds.Lectures()))
	}
}
