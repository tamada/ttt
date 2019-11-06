[![Build Status](https://travis-ci.com/tamada/ttt.svg?branch=master)](https://travis-ci.com/tamada/ttt)
[![codebeat badge](https://codebeat.co/badges/f83e54cf-f7fb-4c52-839f-2f51c9c3e363)](https://codebeat.co/projects/github-com-tamada-ttt-master)
[![Coverage Status](https://coveralls.io/repos/github/tamada/ttt/badge.svg?branch=implements_by_go)](https://coveralls.io/github/tamada/ttt?branch=implements_by_go)
[![License](https://img.shields.io/badge/License-WTFPL-blue.svg)](https://github.com/tamada/ttt/blob/master/LICENSE)
[![Version](https://img.shields.io/badge/Version-1.0.0-yellowgreen.svg)](https://github.com/tamada/ttt/releases/tag/v1.0.0)

# ttt

このツールは，KSU ISEにおける各コースの終了要件を単位取得科目から確認するためのツールです．
ttt は「単位をたくさん取ろう」の頭文字です．

なお，このツールを使った事による不利益は，いかなる場合であっても一切保証しません．

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

[Webページ](https://tamada.github.io/ttt/checker.html)から，Web上で判定できます．

### Docker を使った実行

```sh
$ docker run --rm -v "$PWD"/home/ttt tamada/ttt:1.0.0 credits.json
```

`credits.json` のあるディレクトリで実行してください．
なお，docker のオプション，引数の意味は次の通りです．

* `--rm`
    * Docker終了後，コンテナは自動的に削除される．
* `-v "$PWD":/home/ttt`
    * ホストOSの`$PWD`（現在のディレクトリ）を，コンテナOSの`/home/ttt`に割り当てる．
    * コンテナOSでは，`/home/ttt`で`ttt`が実行される（`ttt`ユーザにて実行される）．
* `tamada/ttt:1.0.0`
    * [Docker Hub](https://hub.docker.com/repository/docker/tamada/ttt)で配布されているDockerfileを特定するためのID．
    * `tamada`が`ttt`という名前でバージョン`1.0.0`として配布しているDockerfileを利用する．
* `credits.json`
    * ホストOSにある `$PWD/credits.json` を `ttt` の引数に与える．

## 判定方法

`data`ディレクトリに，`lectures.json`と`courses.json`が含まれています．
`lectures.json`が科目一覧で，科目ごとの科目名，配当学年，単位数が記録されています．
`courses.json`はコース一覧を表しています．
詳細は，`data` フォルダをご覧ください．

## インストール方法

### Homebrew

```sh
$ brew tap tamada/brew
$ brew install ttt
```

### Docker

[Dockerを使った実行](#docker-を使った実行)を参照のこと．

### Go lang

```sh
$ go get github.com/tamada/ttt
```

### 手作業でインストール（バイナリインストール）

* [Releaseページ](https://github.com/tamada/ttt/releases)から適切なバージョン，OS，アーキテクチャのファイルをダウンロードする．
* ダウンロードした tar.gz ファイルを慎重する．
* 適切な場所にインストールする．
    * `data`ディレクトリは`/usr/local/share/ttt/data` 以下にあるか，カレントディレクトリ以下にあることを前提としています．

### 手作業でインストール（ソースからビルド）

* [GitHubのページ](https://github.com/tamada/ttt)からリポジトリをクローンしてください．
    * `git clone https://github.com/tamada/ttt`
* 作成されたディレクトリに移動してください．
    * `cd ttt`
* `make`を実行してください．
    * `make`
* カレントディレクトリに `ttt` というバイナリが作成されます．
