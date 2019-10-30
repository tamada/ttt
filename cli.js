'use strict'

const VERSION = '1.0.0'

const Ziraffe = require('./ziraffe')

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
    if (args["l"]) {
        return { "type": forceUpper(args["l"]), "origin": args["l"] }
    } else if (args["log"]) {
        return { "type": forceUpper(args["log"]), "origin": args["log"] }
    }
    return { "type": "WARN" }
}

const findYear = (args) => {
    if (args["y"]) {
        return { "number": parseInt(args["y"]), "origin": args["y"] }
    } else if (args["year"]) {
        return { "number": parseInt(args["year"]), "origin": args["year"] }
    }
    return { "number": 2018 }
}

const validateArgs = (args, success, failure) => {
    if (args["_"].length == 0) {
        failure("CREDITS.JSON が指定されていません")
        return
    }
    const logType = findLogType(args)
    if (logType.type !== "IGNORE" && logType.type !== "WARN" && logType.type !== "SEVERE") {
        failure(`${logType.origin}: 未知の挙動設定です`)
        return
    }
    const year = findYear(args)
    if (isNaN(year.number)) {
        failure(`${year.origin}: 入学年が解釈できません`)
        return
    }
    else if (year.number <= 0) {
        failure(`${year.origin}: 入学年が負数です`)
        return
    }
    success(args["_"], logType.type, year.number)
}

const findCredits = (file) => {
    if (!file.startsWith("./") && !file.startsWith("../") && !file.startsWith("/")) {
        file = "./" + file
    }
    return require(file).credits
}

const printResults = (results) => {
    results.forEach(course => {
        console.log(`コース: ${course.name}`)
        console.log(`    取得単位数:   ${course.got_credits}`)
        console.log(`    取得必修科目: ${course.got_requirements.length}科目/${course.requirements.length}科目`)
        console.log(`        ${course.got_requirements.join("\n        ")}`)
        console.log(`    残り必修科目: ${course.rest_requirements.length}科目`)
        console.log(`        ${course.rest_requirements.join("\n        ")}`)
    })
}

const Cli = class Cli {
    exitOnError(ziraffe, credits, logType) {
        let array = credits.filter(name => !ziraffe.isFindLecture(name))
        if (logType !== "IGNORE") {
            array.forEach(name => {
                const maybes = ziraffe.similarLectures(name).map(item => item.name1)
                console.log(`${name}: 講義科目が見つかりません．もしかして，${maybes.join(", ")}`)
            })
        }
        if (logType == "SEVERE" && array.length > 0) {
            return true
        }
        return false
    }

    performEach(ziraffe, file, logType, year) {
        const credits = findCredits(file)
        if (this.exitOnError(ziraffe, credits, logType)) {
            return
        }
        const results = ziraffe.graduationCheck(credits)
        printResults(results)
    }

    perform(argv, failure) {
        var args = require('minimist')(argv)
        if (args["h"] || args["help"]) {
            printHelp()
            return
        }
        validateArgs(args, (args2, logType, year) => {
            const ziraffe = new Ziraffe()
            args2.forEach(file => {
                this.performEach(ziraffe, file, logType, year)
            })
        }, (message) => {
            failure(message)
        })
    }
}

// new Cli().perform(process.argv.slice(2), (message) => {
//     console.log(message)
//     printHelp()
// })

module.exports = Cli
