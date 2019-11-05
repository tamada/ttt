package ttt

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

/*
JSONDataStore is an instance of DataStore.
*/
type JSONDataStore struct {
	initialized bool
	courses     []Course
	lectures    []Lecture
}

/*
NewJSONDataStore creates JsonDataStore object.
*/
func NewJSONDataStore() DataStore {
	ds := JSONDataStore{initialized: false, courses: []Course{}, lectures: []Lecture{}}
	ds.Init()
	return &ds
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func pathOfData(base string) string {
	path := filepath.Join("/usr/local/share/ttt", base)
	if exists(path) {
		return path
	}
	return base
}

/*
Init conducts initialization process for DataStore.
*/
func (ds *JSONDataStore) Init() error {
	if !ds.initialized {
		if err := readJSON(&ds.courses, pathOfData("data/courses.json")); err != nil {
			return err
		}
		if err := readJSON(&ds.lectures, pathOfData("data/lectures.json")); err != nil {
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

/*
Courses returns the couse list.
*/
func (ds *JSONDataStore) Courses() []Course {
	return ds.courses
}

/*
Lectures returns the lecture list.
*/
func (ds *JSONDataStore) Lectures() []Lecture {
	return ds.lectures
}
