package ttt

import (
	"testing"
)

var ds = NewJSONDataStore()
var z = NewVerifier(ds)

func TestSimilarLectures(t *testing.T) {
	testdata := []struct {
		name      string
		wontCount int
	}{
		{"微分積分IV", 2},
		{"ソフトウェア工学I", 0},
		{"コンピュータのための音楽", 1},
	}

	for _, td := range testdata {
		lectures := z.FindSimilarLectures(td.name)
		if len(lectures) != td.wontCount {
			t.Errorf("%s similar count did not match, wont %d, got %d", td.name, td.wontCount, len(lectures))
		}
	}
}

func TestFindCourse(t *testing.T) {
	testdata := []struct {
		name      string
		wontCount int
	}{
		{"全", 1},
		{"脳", 1},
		{"システム", 3},
	}

	for _, td := range testdata {
		courses := z.FindCourses(td.name)
		if len(courses) != td.wontCount {
			t.Errorf("%s: course count did not match, wont %d, got %d", td.name, td.wontCount, len(courses))
		}
	}
}

func TestFindLecture(t *testing.T) {
	ds.Init()
	testdata := []struct {
		name   string
		found  bool
		grade  Grade
		credit CreditCount
	}{
		{"基礎プログラミング演習I", true, 1, 2},
		{"基礎プログラミング演習", false, -1, -1},
	}

	for _, td := range testdata {
		lecture := z.FindLecture(td.name)
		if (lecture != nil) != td.found {
			t.Errorf("%s not found", td.name)
		}
		if td.found {
			if lecture.Credit != td.credit {
				t.Errorf("%s credit did not match, wont %d, got %d", td.name, td.credit, lecture.Credit)
			}
			if lecture.Grade != td.grade {
				t.Errorf("%s grade did not match, wont %d, got %d", td.name, td.grade, lecture.Grade)
			}
		}
	}
}

func TestDiplomaCheck(t *testing.T) {
	ds.Init()

	testdata := []struct {
		course               string
		gotNames             []string
		wontRequirementCount int
		wontGotCredit        CreditCount
		wontRest             int
	}{
		{"ネットワークシステム", []string{"離散数学", "情報理論"}, 8, 4, 6},
	}

	for _, td := range testdata {
		course := z.FindCourses(td.course)
		if len(course) != 1 {
			t.Errorf("%s course not found", td.course)
		}
		result := z.Verify(td.gotNames, course[0])
		if result.Name != td.course {
			t.Errorf("course name did not match, wont %s, got %s", td.course, result.Name)
		}
		if len(result.Requirements) != td.wontRequirementCount {
			t.Errorf("%s requirement count did not match, wont %d, got %d", td.course, td.wontRequirementCount, len(result.Requirements))
		}
		if result.GotCredit != td.wontGotCredit {
			t.Errorf("%s got credit count did not match, wont %d, got %d", td.course, td.wontGotCredit, result.GotCredit)
		}
		if len(result.RestRequirements) != td.wontRest {
			t.Errorf("%s rest requirements did not match, wont %d, got %d", td.course, td.wontRest, len(result.RestRequirements))
		}
	}
}
