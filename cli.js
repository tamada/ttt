'use strict'

const VERSION = '1.0.0'

const Ziraffe = require('ziraffe')

const printHelp = () => {
  console.log(`ziraffe version ${VERSION}
$ node ziraffe/cli.js [OPTIONS] <CREDITS.JSON>
OPTIONS
    -y, --year=<YEAR>    入学年を西暦4桁で入力する．デフォルトは2018．
    -l, --log=<TYPE>     エラーが起こったときの挙動を設定する．デフォルトは WARN（警告して実行）．
                         有効値は IGNORE（無視して実行）, WARN, SEVERE（エラー報告して終了）．
    -h, --help           このメッセージを表示して終了する．`)
}

const forceUpper = (s) => {
    return s.replace(/[a-z]/g, function(ch) {
        return String.fromCharCode(ch.charCodeAt(0) & ~32)
    })
}

const findLogType = (args) => {
    if(args["l"]){
        return { "type": forceUpper(args["l"]), "origin": args["l"] }
    } else if(args["log"]){
        return { "type": forceUpper(args["log"]), "origin": args["log"] }
    }
    return { "type": "WARN" }
}

const findYear = (args) => {
    if(args["y"]){
        return { "number": parseInt(args["y"]), "origin": args["y"]}
    } else if(args["year"]){
        return { "number": parseInt(args["year"]), "origin": args["year"] }
    }
    return { "number": 2018 }
}

const validateArgs = (args, success, failure) => {
    if(args["_"].length == 0){
        failure("CREDITS.JSON が指定されていません")
        return
    }
    const logType = findLogType(args)
    if(logType.type !== "IGNORE" && logType.type !== "WARN" && logType.type !== "SEVERE") {
        failure(`${logType.origin}: 未知の挙動設定です`)
        return
    }
    const year = findYear(args)
    if(isNaN(year.number)){
        failure(`${year.origin}: 入学年が解釈できません`)
        return
    }
    else if(year.number <= 0){
        failure(`${year.origin}: 入学年が負数です`)
        return
    }
    success(args["_"], logType.type, year.number)
}

const Cli = class Cli {
    performEach = (ziraffe, file, logType, year) => {
        const credits = require(file)

    }
    perform = (argv, failure) => {
        var args = require('minimist')(argv)
        if(args["h"] || args["help"]) {
            printHelp()
            return
        }
        validateArgs(args, (args2, logType, year) => {
            const ziraffe = new Ziraffe()
            args2.forEach(file => {
                execEach(ziraffe, file, logType, year)
            })
        }, (message) => {
            failure(message)
        })
    }
}

new Cli().perform(process.argv.slice(2), (message) => {
    console.log(message)
    printHelp()
})

module.exports = Cli
