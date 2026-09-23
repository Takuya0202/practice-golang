# AGENTS.md

書籍『改訂新版 Go言語プログラミングエッセンス』（mattn 著、技術評論社、2025年）を読みながら Go を学ぶためのリポジトリ。
Claude Code と Codex の両方がこのファイルを読む。役ごとの振る舞いはスキル（後述）に書く。

## 大原則

- 目的はユーザー自身が書いて理解すること。AI は答えを先回りして書かない。
- 回答は日本語で行う。
- Go の挙動は go.mod の Go バージョン（1.25）を基準に説明する。ユーザーが示した本の記述や一般的な知識が現行の挙動と違う場合は、そのことを明示する。AI は本の本文を持っていないので、本に何が書いてあるかを推測で語らない。
- 挙動について自信がないときは、推測で断定せず `go doc` や小さなコードを実際に動かして確かめる。
  - 確かめるためのコードは `.scratch/<名前>/main.go` に書き、`go run ./.scratch/<名前>` で動かす。`.scratch/` は git 管理外で、`./...` の対象にもならない。後片付けは不要で、報告にも書かない。
  - 組み込み関数のドキュメントは `go doc -u builtin.append` のように `-u` を付ける。
- 標準ライブラリだけを使う。外部パッケージは、本がそれを扱う章に入ってから、ユーザーの了承を得て追加する。

## ディレクトリ構成

```
practice-golang/
├── go.mod              # モジュールはリポジトリ直下に1つ
├── AGENTS.md           # このファイル（共通ルール）
├── CLAUDE.md           # AGENTS.md を読み込むだけ
├── docs/progress.md    # 進捗とつまずきメモ（役どうしの共有メモ）
├── .agents/skills/     # 役の定義（正本）
├── .claude/skills/     # .agents/skills へのシンボリックリンク
├── .codex/config.toml  # Codex 用設定（Go のビルドキャッシュを /tmp に置く）
├── .scratch/           # AI が挙動を確かめる使い捨てコード（git 管理外）
├── practice/           # 本を読みながらユーザーが写経・実行するコード
├── exercises/          # go-drill が作る練習問題
└── lesson_go/          # 以前に本を途中まで進めたときの記録（別モジュール）
```

- `practice/chNN/MM_topic/`：本のサンプルを写経・改造する場所。ユーザーが書く。AI は頼まれない限り編集しない。
- `exercises/chNN/MM_topic/`：練習問題。go-drill が作り、ユーザーが `// TODO:` を埋める。
- `lesson_go/`：参照のみ。編集しない。
- 命名：`NN` は章番号の2桁（`ch03`）、`MM` は章内の連番の2桁、`topic` は英小文字の snake_case（例：`exercises/ch03/01_slice_append/`）。
- Go は1ディレクトリ＝1パッケージ。サンプルや問題は1つずつディレクトリを分ける。
- package 名は、ディレクトリ名から先頭の `MM_` を除き、残りのアンダースコアを詰めたもの（例：`01_slice_append` → `sliceappend`、`03_utf8_decode` → `utf8decode`）。`practice/` の実行用コードは `package main`。
- topic を Go のキーワード（`map`, `select`, `defer`, `interface`, `switch`, `range`, `chan`, `func`, `go`, `type` など）1語だけにしない。`02_map_basics` のように2語以上にする。
- ファイル名の末尾（`.go` の直前）を GOOS/GOARCH 名（`_linux`, `_windows`, `_js`, `_wasm`, `_amd64`, `_arm64` など）にしない。Go がビルド制約とみなしてファイルを除外する。

## よく使うコマンド

```bash
go run ./practice/ch03/01_hello        # 写経したサンプルを実行
go test ./exercises/ch03/01_slice_append  # 1問だけ答え合わせ
go test ./exercises/...                # 全問（問題が1つもない間は "matched no packages" と出るが異常ではない）
go vet ./...                           # 静的チェック（lesson_go は別モジュールなので対象外）
gofmt -l practice exercises            # 何も表示されなければOK（差分があってもエラー終了はしない）
go test -race ./exercises/ch06/...     # 並行処理の問題は -race も
```

## 役（スキル）

役ごとに別セッションで動かす。セッションの最初に1回呼べば、そのセッションの間は同じ役を続ける。

| 役 | 内容 | Claude Code | Codex |
|---|---|---|---|
| go-ask | 質問役。聞かれたことだけに短く答える | `claude "/go-ask"` | `codex '$go-ask'` |
| go-drill | 出題役。`exercises/` に TODO 付きの問題とテストを作る | `claude "/go-drill ch03 スライス"` | `codex '$go-drill ch03 スライス'` |
| go-review | レビュー役。実際に動かして確認し、どこが・なぜを伝える | `claude "/go-review exercises/ch03"` | `codex '$go-review exercises/ch03'` |

役が指定されていないセッションでも、大原則は守る（練習問題の答えを書かない、など）。

各スキルは明示的に呼んだときだけ動く。Claude Code 向けに SKILL.md の `disable-model-invocation: true`、Codex 向けに `agents/openai.yaml` の `policy.allow_implicit_invocation: false` で指定している。前者は Agent Skills 仕様の外のキーだが、意図して置いているので消さない。

## docs/progress.md

役どうしは別セッションなので、情報の受け渡しはこのファイルで行う。

- 章の進捗：ユーザーが更新する。どの役でも「ch03 を読書中にして（スライスまで）」のように頼まれたら更新してよい。
- 練習問題：go-drill が問題を作ったら行を足し、go-review が結果を書き込む。
- つまずきメモ：go-review が理解のズレを見つけたら追記する。go-ask は「メモして」と頼まれたときだけ追記する。
- go-drill は、章やトピックの指定がないとき、つまずきメモのうち解消していないものから出題する。

## 本の章立て

出題範囲の判断に使う。ユーザーがまだ読んでいない章の機能は、練習問題で使わせない。

| 章 | タイトル | 主な内容 |
|---|---|---|
| 1 | プログラミング言語Goとは | 歴史、立ち位置、利用される場面 |
| 2 | 開発環境の準備 | インストール、セットアップ |
| 3 | 基本的な文法 | 静的な型、基本的な構文、goroutine、Go モジュール、プロジェクトレイアウト、lint、go fmt |
| 4 | 基本テクニックとベストプラクティス | ビルトイン関数、パッケージ、build constraints、cgo、go:embed、Functional Options Pattern、Builder Pattern、internal パッケージ、Embedded struct |
| 5 | Webアプリケーション開発に必要な要素 | net/http、html/template、log/slog、net/smtp、WebAssembly |
| 6 | 速いプログラムのためのテクニック | 並行と並列、goroutine、channel、非同期パターン |
| 7 | テストにおけるテクニック | テストの基本、便利なテクニック、Fuzzing |
| 8 | ベンチマークにおけるテクニック | ベンチマークの基本と比較、プロファイリング |
| 9 | GoによるCLIアプリケーション開発 | DB 登録・照会プログラム、テスト、フラグ・端末制御ライブラリ |
| 10 | GoによるWebアプリケーション開発 | TODO アプリ、リマインダメール、フレームワーク |
| 11 | GitHubでの開発における勘所 | パッケージ名、バージョニング、ドキュメント、自動テスト・自動リリース |
| 12 | データベースの扱い方 | database/sql、ent |
| 13 | Goとクラウドサービス | Google Cloud、AWS Lambda、Oracle Cloud |
