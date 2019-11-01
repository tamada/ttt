package main

import (
	"encoding/json"
	"fmt"

	"github.com/tamada/ziraffe"
)

type WasmDataStore struct {
	initialized bool
	courses     []ziraffe.Course
	lectures    []ziraffe.Lecture
}

func NewWasmDataStore() *WasmDataStore {
	ds := WasmDataStore{initialized: false, courses: []ziraffe.Course{}, lectures: []ziraffe.Lecture{}}
	ds.Init()
	return &ds
}

func initImpl(ds *WasmDataStore) {
	loadJson(&courses, COURSES_JSON)
	loadJson(&lectures, LECTURES_JSON)
}

func loadJson(target interface{}, jsonString string) {
	if err := json.Unmarshal(target, jsonString); err != nil {
		fmt.Println(err.Error())
	}
}

func (ds *WasmDataStore) Init() {
	if !ds.initialized {
		initImpl(ds)
		ds.initialized = true
	}
}

func (ds *WasmDataStore) Courses() []ziraffe.Course {
	return ds.courses
}

func (ds *WasmDataStore) Lectures() []ziraffe.Lecture {
	return ds.lectures
}

const LECTURES_JSON = ""
const COURSES_JSON = ""
