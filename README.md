[![Build Status](https://travis-ci.com/tamada/ttt.svg?branch=master)](https://travis-ci.com/tamada/ttt)
[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/ttt/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-yellowgreen.svg)](https://github.com/tamada/ttt/releases/tag/v1.0.0)

# ttt

このツールは，KSU ISEにおける各コースの終了要件を単位取得科目から確認するためのツールです．
ttt は「たくさん単位を取ろう」の頭文字です．

## 使い方

`credits.json` を準備し，次のようにコマンドを実行してください．

```sh
$ ttt credits.json
```

### `credits.json`

今まで単位を取得した科目を列挙したJSONファイルです．
文字列の配列として科目を列挙してください．

具体的には，次のように単位取得科目を書いていってください．

```json
[
    "ソフトウェア工学I",
    "基礎プログラミング演習I",
    "基礎プログラミング演習II",
    "発展プログラミング演習",
]
```

### Help

```sh
$ ttt --help
ttt バージョン 1.0.0
ttt [OPTIONS] <CREDITS.JSON...>
OPTIONS
    -c, --course=<COURSE>    特定のコースのみの判定を行う．部分一致．
                             指定されない場合は全コースで判定を行う．
    -e, --on-error=<TYPE>    エラー時の挙動を設定する．デフォルトは WARN（エラーを表示して続行）．
                             有効値は IGNORE（エラーを無視），WARN，QUIT（エラーを表示して終了）．
    -y, --year=<YEAR>        入学年を指定する．デフォルトは 2018．
    -v, --verbose            冗長出力モード．デフォルトはOFF．
    -h, --help               このメッセージを表示する．
ARGUMENTS
    CREDITS.JSON             単位を取得した講義を列挙したJSONファイル．複数指定可能．
```

### Web上での実行

[Webページ](https://tamada.github.io/ttt/verify.html)から，Web上で判定できます．

## 判定方法

`data`ディレクトリに，`lectures.json`と`courses.json`が含まれています．
`lectures.json`が科目一覧で，科目ごとの科目名，配当学年，単位数が記録されています．
`courses.json`はコース一覧を表しています．
詳細は，`data` フォルダをご覧ください．
