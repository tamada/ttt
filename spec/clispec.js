const Cli = require('../cli')

describe('コマンドライン引数の解釈をテストする．', () => {
    const cli = new Cli()
    it('credits.json を指定せずに実行しようとして失敗する．', () => {
        cli.perform([], (message) => {
            expect('CREDITS.JSON が指定されていません').toBe(message)
        })
    })
    it('log type に不正なものを指定する（info）．', () => {
        cli.perform(["-l", "info", "credits.json"], (message) => {
            expect('info: 未知の挙動設定です').toBe(message)
        })
    })
    it('log type に不正なものを指定する（Warning）．', () => {
        cli.perform(["--log=Warning", "credits.json"], (message) => {
            expect('Warning: 未知の挙動設定です').toBe(message)
        })
    })
    it('year に不正なものを指定する（負数）．', () => {
        cli.perform(["--year=-100", "credits.json"], (message) => {
            expect('-100: 入学年が負数です').toBe(message)
        })
    })
    it('year に不正なものを指定する（数値ではない値）．', () => {
        cli.perform(["-y", "abc", "credits.json"], (message) => {
            expect('abc: 入学年が解釈できません').toBe(message)
        })
    })
})
