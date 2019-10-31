package ziraffe

import (
	"encoding/json"
	"io/ioutil"
)

type JsonDataStore struct {
}

var initialized = false
var courses = []Course{}
var lectures = []Lecture{}
var diploma = Diploma{}

func NewJsonDataStore() DataStore {
	ds := JsonDataStore{}
	ds.Init()
	return &ds
}

func (ds JsonDataStore) Init() error {
	if !initialized {
		if err := readJSON(&courses, "data/courses.json"); err != nil {
			return err
		}
		if err := readJSON(&lectures, "data/lectures.json"); err != nil {
			return err
		}
		if err := readJSON(&diploma, "data/diploma_rules.json"); err != nil {
			return err
		}
	}
	initialized = true
	return nil
}

func readJSON(target interface{}, path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, target); err != nil {
		return err
	}
	return nil
}

func (ds JsonDataStore) Diploma() Diploma {
	return diploma
}

func (ds JsonDataStore) Courses() []Course {
	return courses
}

func (ds JsonDataStore) Lectures() []Lecture {
	return lectures
}
