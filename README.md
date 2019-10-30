# ziraffe

このツールは，KSU ISEにおける各コースの終了要件を単位取得科目から確認するためのツールです．

## 使い方

`credits.json` を準備し，次のようにコマンドを実行してください．

```sh
$ node ziraffe/index.js credits.json
```

### credits.json

次のように単位取得科目を書いていってください．

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
$ node ziraffe/index.js --help
ziraffe version 1.0.0
OPTIONS
    -y, --year=<YEAR>     入学年を西暦4桁で入力する．デフォルトは2018．
    -l, --log=<TYPE>      エラーが起こったときの挙動を設定する．デフォルトは WARN（警告して実行）．
                          有効値は IGNORE（無視して実行）, WARN, SEVERE（エラー報告して終了）．
    -h, --help            このメッセージを表示して終了する．
```
