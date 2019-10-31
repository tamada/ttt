package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
	"github.com/tamada/ziraffe"
)

/* VERSION represents the version of ziraffe. */
const VERSION = "1.0.0"

func printResult(result ziraffe.CourseDiplomaResult, opts *options) {
	fmt.Printf("コース: %s （必修修得状況　%d/%d, %d/%d単位）\n", result.Name,
		len(result.GotRequirements), len(result.Requirements), result.GotCredit, result.DiplomaCredit)
	if opts.verboseFlag {
		fmt.Printf("    修得済の必修: %s\n", strings.Join(result.GotRequirements, ", "))
		fmt.Printf("    未修得の必修: %s\n", strings.Join(result.RestRequirements, ", "))
	}
}

func printResults(results []ziraffe.CourseDiplomaResult, opts *options) error {
	for _, result := range results {
		printResult(result, opts)
	}
	return nil
}

func checkCredits(credits []string, opts *options, z *ziraffe.Ziraffe) error {
	courses := z.FindCourses(opts.course)
	results := []ziraffe.CourseDiplomaResult{}
	for _, course := range courses {
		result := z.CheckCourse(credits, course)
		results = append(results, result)
	}
	return printResults(results, opts)
}

func findLectureNames(lectures []ziraffe.Lecture) string {
	list := []string{}
	for _, lec := range lectures {
		list = append(list, lec.Name)
	}
	return strings.Join(list, ", ")
}

func validateCredits(credits []string, z *ziraffe.Ziraffe) []string {
	result := []string{}
	for _, credit := range credits {
		lectures := z.FindSimilarLectures(credit)
		if len(lectures) != 0 {
			fmt.Printf("%s: 科目名が不正です．もしかして，%s\n", credit, findLectureNames(lectures))
		} else {
			result = append(result, credit)
		}
	}
	return result
}

func performEach(fileName string, opts *options, z *ziraffe.Ziraffe) error {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	credits := []string{}
	if err := json.Unmarshal(bytes, &credits); err != nil {
		return err
	}
	credits = validateCredits(credits, z)
	return checkCredits(credits, opts, z)
}

func showError(err error, opts *options) {
	if opts.onError == WARN || opts.onError == QUIT {
		fmt.Println(err.Error())
	}
}

func perform(opts *options) int {
	ds := ziraffe.NewJsonDataStore()
	z := ziraffe.NewZiraffe(ds)
	for _, credits := range opts.args {
		err := performEach(credits, opts, z)
		if err != nil {
			showError(err, opts)
			if opts.onError == QUIT {
				return 1
			}
		}
	}
	return 0
}

func printHelp(prog string) {
	fmt.Printf(`%s vresion %s
%s [OPTIONS] <CREDITS.JSON...>
OPTIONS
    -c, --course=<COURSE>    特定のコースのみの判定を行う．部分一致．
                             指定されない場合は全コースで判定を行う．
    -e, --on-error=<TYPE>    エラー時の挙動を設定する．デフォルトは WARN（エラーを表示して続行）．
                             有効値は IGNORE（エラーを無視），WARN，QUIT（エラーを表示して終了）．
    -y, --year=<YEAR>        入学年を指定する．デフォルトは 2018．
	-v, --verbose            冗長出力モード．デフォルトはOFF．
    -h, --help               このメッセージを表示する．
ARGUMENTS
    CREDITS.JSON      `, prog, VERSION, prog)
}

type ErrorType int

const (
	IGNORE ErrorType = iota
	WARN
	QUIT
)

func (e ErrorType) String() string {
	switch e {
	case IGNORE:
		return "IGNORE"
	case WARN:
		return "WARN"
	case QUIT:
		return "QUIT"
	default:
		return "UNKNOWN"
	}
}

type options struct {
	course        string
	onErrorString string
	year          int
	helpFlag      bool
	verboseFlag   bool
	onError       ErrorType
	args          []string
}

func buildFlagSet() (*flag.FlagSet, *options) {
	opts := new(options)
	flags := flag.NewFlagSet("ziraffe", flag.ContinueOnError)
	flags.Usage = func() { printHelp("ziraffe") }
	flags.StringVarP(&opts.course, "course", "c", "", "specifies course name (partial match)")
	flags.StringVarP(&opts.onErrorString, "on-error", "e", "WARN", "specifies the behavior on error (default: WARN)")
	flags.IntVarP(&opts.year, "year", "y", 2018, "specifies admission year (default: 2018)")
	flags.BoolVarP(&opts.verboseFlag, "verbose", "v", false, "verbose mode")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "print this message")
	return flags, opts
}

func parseOnErrorString(opts *options) error {
	one := strings.ToLower(opts.onErrorString)
	if one == "warn" {
		opts.onError = WARN
	} else if one == "ignore" {
		opts.onError = IGNORE
	} else if one == "quit" {
		opts.onError = QUIT
	} else {
		return fmt.Errorf("%s: 未知のエラー時の挙動です", opts.onErrorString)
	}
	return nil
}

func parseArgs(args []string) (*options, error) {
	flags, opts := buildFlagSet()
	if err := flags.Parse(args); err != nil {
		return nil, err
	}
	if err := parseOnErrorString(opts); err != nil {
		return nil, err
	}
	opts.args = flags.Args()[1:]
	return opts, nil
}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return perform(opts)
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
