package ziraffe

import (
	"encoding/json"
	"io/ioutil"
)

type JsonDataStore struct {
	initialized bool
	courses     []Course
	lectures    []Lecture
}

func NewJsonDataStore() DataStore {
	ds := JsonDataStore{initialized: false, courses: []Course{}, lectures: []Lecture{}}
	ds.Init()
	return &ds
}

func (ds *JsonDataStore) Init() error {
	if !ds.initialized {
		if err := readJSON(&ds.courses, "data/courses.json"); err != nil {
			return err
		}
		if err := readJSON(&ds.lectures, "data/lectures.json"); err != nil {
			return err
		}
		ds.initialized = true
	}
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

func (ds *JsonDataStore) Courses() []Course {
	return ds.courses
}

func (ds *JsonDataStore) Lectures() []Lecture {
	return ds.lectures
}
