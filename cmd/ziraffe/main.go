package main

import (
  "os"
  "fmt"
  "strings"
  flag "github.com/spf13/pflag"
)

const VERSION = "1.0.0"

type Credits {
  credits []string `json:"credits"`
}

func performEach(fileName string, opts *options) {
  
}

func perform(opts *options) {
  for _, credits := range opts.args {
    performEach(credits, opts)
  }
}

func printHelp(prog string){
  fmt.Printf(`%s vresion %s
%s [OPTIONS] <CREDITS.JSON...>
OPTIONS
    -c, --course=<COURSE>    特定のコースのみの判定を行う．部分一致．
                             指定されない場合は全コースで判定を行う．
    -e, --on-error=<TYPE>    エラー時の挙動を設定する．デフォルトは WARN（エラーを表示して続行）．
                             有効値は IGNORE（エラーを無視），WARN，QUIT（エラーを表示して終了）．
    -y, --year=<YEAR>        入学年を指定する．デフォルトは 2018．
    -h, --help               このメッセージを表示する．
ARGUMENTS
    CREDITS.JSON      `, prog, VERSION, prog);
}

type ErrorType int
const (
  IGNORE ErrorType = iota,
  WARN,
  QUIT
)

func (e ErrorType) String() string {
  switch(e) {
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
  course string
  onErrorString string
  year int
  helpFlag bool
  onError onError
  args []string
}

func buildFlagSet() (*flag.FlagSet, *options) {
  opts := new(options)
  flags := NewFlagSet("ziraffe", flag.ContinueOnError)
  flags.Usage = func() { printHelp("ziraffe") }
  flags.StringVarP(&opts.course, "course", "c", "", "specifies course name (partial match)")
  flags.StringVarP(&opts.onErrorString, "on-error", "e", "WARN", "specifies the behavior on error (default: WARN)")
  flags.IntVarP(&opts.year, "year", "y", 2018, "specifies admission year (default: 2018)")
  flags.BoolVarP(&opts.help, "help", "h", false, "print this message")
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
    return fmt.Errorf("%s: 未知のエラー時の挙動です")
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
  opts.args = flags.Args()
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
