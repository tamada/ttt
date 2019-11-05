package ttt

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
CreditCount means credits of a lectures.
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

/*
CourseDiplomaResult shows the verification result of course diploma.
*/
type CourseDiplomaResult struct {
	Name             string
	Requirements     []string
	DiplomaCredit    CreditCount
	GotCredit        CreditCount
	GotRequirements  []string
	RestRequirements []string
}

/*
Checker is for checking course diploma.
*/
type Checker struct {
	Store DataStore
}

/*
NewChecker creates an object of Verifier.
*/
func NewChecker(ds DataStore) *Checker {
	z := Checker{Store: ds}
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

func (z *Checker) findCreditOfLecture(name string) CreditCount {
	lecture := z.FindLecture(name)
	if lecture == nil {
		return 0
	}
	return lecture.Credit
}

func (z *Checker) countNumberOfCredits(gotCredits []string, course Course) CreditCount {
	var sum CreditCount
	for _, credit := range gotCredits {
		if contains(course.Requirements, credit) || contains(course.Recommends, credit) {
			sum += z.findCreditOfLecture(credit)
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

/*
Check verifies the course diploma.
*/
func (z *Checker) Check(gotCredits []string, course Course) CourseDiplomaResult {
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
func (z *Checker) FindCourses(name string) []Course {
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

/*
FindLecture finds a lecture from the given name.
*/
func (z *Checker) FindLecture(name string) *Lecture {
	for _, lecture := range z.Store.Lectures() {
		if lecture.Name == name {
			return &lecture
		}
	}
	return nil
}

func createDistances(name string, z *Checker) []distance {
	distances := []distance{}
	for _, lecture := range z.Store.Lectures() {
		distances = append(distances, distance{distance: LevenshteinS(name, lecture.Name), lecture: lecture})
	}
	return distances
}

func sortDistances(distances []distance) {
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})
}

/*
FindSimilarLectures finds lectures which have similar name with the given name.
If exact matched name of lecture is exist, this function returns the 0-sized array.
*/
func (z *Checker) FindSimilarLectures(name string) []Lecture {
	distances := createDistances(name, z)
	sortDistances(distances)
	min := distances[0].distance
	if min == 0 {
		return []Lecture{}
	}
	return findLecturesWithMinimumDistance(min, distances)
}

func findLecturesWithMinimumDistance(min int, distances []distance) []Lecture {
	results := []Lecture{}
	for _, d := range distances {
		if d.distance == min {
			results = append(results, d.lecture)
		} else {
			break
		}
	}
	return results
}
