package main

import (
	"fmt"
	"syscall/js"

	"github.com/tamada/ziraffe"
)

func lectures(this js.Value, args []js.Value) interface{} {
	return ds.Lectures()
}

func courses(this js.Value, args []js.Value) interface{} {
	return ds.Courses()
}

var ds ziraffe.DataStore
var z *ziraffe.Ziraffe

func gotStringArrayFromJsValue(value js.Value) []string {
	results := []string{}
	length := value.Get("length").Int()
	for i := 0; i < length; i++ {
		results = append(results, value.Index(i).String())
	}
	return results
}

func convertResultToHTML(r ziraffe.CourseDiplomaResult) string {
	return fmt.Sprintf(`<li>%s （必修修得状況　%d/%d, %d/%d単位）</li>`,
		r.Name, len(r.GotRequirements), len(r.Requirements), r.GotCredit, r.DiplomaCredit)
}

func checkDiplomaOfCourses(this js.Value, args []js.Value) interface{} {
	credits := gotStringArrayFromJsValue(args[0])
	html := ""
	for _, course := range ds.Courses() {
		r := z.CheckCourse(credits, course)
		html = html + convertResultToHTML(r)
	}
	doc := js.Global().Get("document")
	element := doc.Call("getElementById", "result-list")
	element.Set("innerHTML", html)
	return ""
}

func initDataStore(this js.Value, args []js.Value) interface{} {
	ds = ziraffe.NewStandaloneDataStore()
	z = ziraffe.NewZiraffe(ds)
	return ""
}

func buildHTMLOfLectures(targetGrade ziraffe.Grade) string {
	resultString := ""
	for _, lecture := range ds.Lectures() {
		if lecture.Grade == targetGrade {
			resultString = fmt.Sprintf("%s<li><input type=\"checkbox\" value=\"%s\">%s (%d)</input></li>",
				resultString, lecture.Name, lecture.Name, lecture.Credit)
		}
	}
	return resultString
}

func buildHTML(this js.Value, args []js.Value) interface{} {
	lecturesOfGrades := [4]string{}
	for i := 0; i < 4; i++ {
		lecturesOfGrades[i] = buildHTMLOfLectures(ziraffe.Grade(i + 1))
	}
	doc := js.Global().Get("document")
	target := doc.Call("getElementById", "lectures-list")
	innerHTML := ""
	for i, lecturesOfGrade := range lecturesOfGrades {
		innerHTML = fmt.Sprintf("%s<li>%d年次<ul>%s</ul></li>", innerHTML, i+1, lecturesOfGrade)
	}
	target.Set("innerHTML", innerHTML)
	return ""
}

func registerCallbacks() {
	js.Global().Set("initDataStore", js.FuncOf(initDataStore))
	js.Global().Set("buildHTML", js.FuncOf(buildHTML))
	js.Global().Set("checkDiplomaOfCourses", js.FuncOf(checkDiplomaOfCourses))
}

func main() {
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}
