package ziraffe

import (
	"sort"
	"strings"
)

/*
DataStore shows master of lectures and course data.
*/
type DataStore interface {
	Lectures() []Lecture
	Courses() []Course
	Init() error
}

/*
Grade means target grades of lectures.
*/
type Grade int

/*
Credit means credits of a lectures.
*/
type CreditCount int

/*
Lecture shows a lecture.
*/
type Lecture struct {
	Name   string      `json:"name"`
	Grade  Grade       `json:"grade"`
	Credit CreditCount `json:"credit"`
}

/*
Course shows requirements and recommended lectures of a course.
*/
type Course struct {
	Name          string      `json:"name"`
	DiplomaCredit CreditCount `json:"diploma-credit"`
	Requirements  []string    `json:"requirements"`
	Recommends    []string    `json:"recommends"`
}

type distance struct {
	distance int
	lecture  Lecture
}

type CourseDiplomaResult struct {
	Name             string
	Requirements     []string
	DiplomaCredit    CreditCount
	GotCredit        CreditCount
	GotRequirements  []string
	RestRequirements []string
}

type Ziraffe struct {
	Store DataStore
}

func NewZiraffe(ds DataStore) *Ziraffe {
	z := Ziraffe{Store: ds}
	return &z
}

func contains(slice []string, item string) bool {
	for _, element := range slice {
		if element == item {
			return true
		}
	}
	return false
}

func (z *Ziraffe) countNumberOfCredits(gotCredits []string, course Course) CreditCount {
	var sum CreditCount = 0
	for _, credit := range gotCredits {
		if contains(course.Requirements, credit) || contains(course.Recommends, credit) {
			lecture := z.FindLecture(credit)
			if lecture != nil {
				sum += lecture.Credit
			}
		}
	}
	return sum
}

func findRequirements(gotCredits []string, requirements []string, includeFunc func(flag bool) bool) []string {
	results := []string{}
	for _, r := range requirements {
		if includeFunc(contains(gotCredits, r)) {
			results = append(results, r)
		}
	}
	return results
}

func (z *Ziraffe) CheckCourse(gotCredits []string, course Course) CourseDiplomaResult {
	return CourseDiplomaResult{
		Name:             course.Name,
		Requirements:     course.Requirements,
		GotCredit:        z.countNumberOfCredits(gotCredits, course),
		DiplomaCredit:    course.DiplomaCredit,
		GotRequirements:  findRequirements(gotCredits, course.Requirements, func(flag bool) bool { return flag }),
		RestRequirements: findRequirements(gotCredits, course.Requirements, func(flag bool) bool { return !flag }),
	}
}

/*
FindCourses finds courses from name with partial matching.
*/
func (z *Ziraffe) FindCourses(name string) []Course {
	results := []Course{}
	if name == "" {
		return z.Store.Courses()
	}
	for _, c := range z.Store.Courses() {
		if strings.Contains(c.Name, name) {
			results = append(results, c)
		}
	}
	return results
}

func (z *Ziraffe) FindLecture(name string) *Lecture {
	for _, lecture := range z.Store.Lectures() {
		if lecture.Name == name {
			return &lecture
		}
	}
	return nil
}

func (z *Ziraffe) FindSimilarLectures(name string) []Lecture {
	distances := []distance{}
	for _, lecture := range z.Store.Lectures() {
		distances = append(distances, distance{distance: LevenshteinS(name, lecture.Name), lecture: lecture})
	}
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})
	min := distances[0].distance
	results := []Lecture{}
	if min == 0 {
		return results
	}
	for _, d := range distances {
		if d.distance == min {
			results = append(results, d.lecture)
		} else {
			break
		}
	}
	return results
}
