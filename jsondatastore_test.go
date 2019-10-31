package ziraffe

import "testing"

func TestReadData(t *testing.T) {
	ds := JsonDataStore{}
	if err := ds.Init(); err != nil {
		t.Error(err)
	}
	if len(ds.Courses()) != 11 {
		t.Errorf("courses size did not match, wont %d, got %d", 11, len(ds.Courses()))
	}
	if len(ds.Lectures()) != 81 {
		t.Errorf("lectures size did not match, wont %d, got %d", 81, len(ds.Lectures()))
	}
	diploma := ds.Diploma()
	if diploma.CourseCredits != 16 {
		t.Errorf("course credits did not match, wont 16, got %d", diploma.CourseCredits)
	}
	if diploma.MajorCredits != 88 {
		t.Errorf("course credits did not match, wont 88, got %d", diploma.MajorCredits)
	}
	if diploma.RequiredCredits != 26 {
		t.Errorf("course credits did not match, wont 26, got %d", diploma.RequiredCredits)
	}
}
