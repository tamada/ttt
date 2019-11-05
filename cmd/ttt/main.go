package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	flag "github.com/spf13/pflag"
	"github.com/tamada/ttt"
)

/*
VERSION represents the version of ttt.
*/
const VERSION = "1.0.0"

func printResult(result ttt.CourseDiplomaResult, opts *options) {
	fmt.Printf("    %s（必修修得状況　%2d/%2d, %2d/%3d単位）\n", result.Name,
		len(result.GotRequirements), len(result.Requirements), result.GotCredit, result.DiplomaCredit)
	if opts.verboseFlag {
		fmt.Printf("        修得済の必修: %s\n", strings.Join(result.GotRequirements, ", "))
		fmt.Printf("        未修得の必修: %s\n", strings.Join(result.RestRequirements, ", "))
	}
}

func printResults(results []ttt.CourseDiplomaResult, fileName string, opts *options) error {
	fmt.Printf("ファイル名 %s\n", fileName)
	for _, result := range results {
		printResult(result, opts)
	}
	return nil
}

func (opts *options) checkCredits(credits []string, fileName string, z *ttt.Checker) error {
	courses := z.FindCourses(opts.course)
	results := []ttt.CourseDiplomaResult{}
	for _, course := range courses {
		result := z.Check(credits, course)
		results = append(results, result)
	}
	return printResults(results, fileName, opts)
}

func findLectureNames(lectures []ttt.Lecture) string {
	list := []string{}
	for _, lec := range lectures {
		list = append(list, lec.Name)
	}
	return strings.Join(list, ", ")
}

func validateCredits(credits []string, z *ttt.Checker) []string {
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

func (opts *options) performEach(fileName string, z *ttt.Checker) error {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	credits := []string{}
	if err := json.Unmarshal(bytes, &credits); err != nil {
		return err
	}
	credits = validateCredits(credits, z)
	return opts.checkCredits(credits, fileName, z)
}

func (opts *options) showError(err error) {
	if opts.onError == WARN || opts.onError == QUIT {
		fmt.Println(err.Error())
	}
}

func (opts *options) perform() int {
	ds := ttt.NewJSONDataStore()
	z := ttt.NewChecker(ds)
	for _, credits := range opts.args {
		err := opts.performEach(credits, z)
		if err != nil {
			opts.showError(err)
			if opts.onError == QUIT {
				return 1
			}
		}
	}
	return 0
}

func getHelpMessage(prog string) string {
	return fmt.Sprintf(`%s バージョン %s
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
    CREDITS.JSON             単位を取得した講義を列挙したJSONファイル．複数指定可能．`, prog, VERSION, prog)
}

/*
ErrorType shows error types.
*/
type ErrorType int

/* IGNORE is the one of ErrorType, it ignore errors. */
const (
	IGNORE ErrorType = iota
	/* WARN is the one of ErrorType, it warn errors and perform the process. */
	WARN
	/* QUIT is the one of ErrorType, it warn errors and exit process. */
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
	flags := flag.NewFlagSet("ttt", flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(getHelpMessage("ttt")) }
	flags.StringVarP(&opts.course, "course", "c", "", "specifies course name (partial match)")
	flags.StringVarP(&opts.onErrorString, "on-error", "e", "WARN", "specifies the behavior on error (default: WARN)")
	flags.IntVarP(&opts.year, "year", "y", 2018, "specifies admission year (default: 2018)")
	flags.BoolVarP(&opts.verboseFlag, "verbose", "v", false, "verbose mode")
	flags.BoolVarP(&opts.helpFlag, "help", "h", false, "print this message")
	return flags, opts
}

func parseOnErrorString(opts *options) error {
	one := strings.ToLower(opts.onErrorString)
	switch one {
	case "warn":
		opts.onError = WARN
	case "ignore":
		opts.onError = IGNORE
	case "quit":
		opts.onError = QUIT
	default:
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
	if len(opts.args) == 0 {
		return opts, fmt.Errorf("CREDITS.JSON が指定されていません．")
	}
	return opts, nil
}

func goMain(args []string) int {
	opts, err := parseArgs(args)
	if opts != nil && opts.helpFlag {
		fmt.Println(getHelpMessage(`ttt`))
		return 0
	}
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return opts.perform()
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
