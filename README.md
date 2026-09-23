# practice-golang

『改訂新版 Go言語プログラミングエッセンス』（mattn 著、技術評論社）を読みながら、AI と一緒に Go を学ぶリポジトリ。

## 進め方

```
本を読む → practice/ に写経して動かす
  ├─ わからない   → 質問役（go-ask）に聞く
  ├─ 章の区切り   → 出題役（go-drill）に問題を作らせて、exercises/ で解く
  └─ 解いた・メモを書いた → レビュー役（go-review）に見てもらう
                     ↓
       つまずきを docs/progress.md に記録 → 次の出題に反映
```

## AI の役

herdr のペインごとに、役を指定して起動する。起動後は普通に話しかけるだけでよい。

| 役 | Claude Code | Codex |
|---|---|---|
| 質問 | `claude "/go-ask"` | `codex '$go-ask'` |
| 出題 | `claude "/go-drill ch03 スライス"` | `codex '$go-drill ch03 スライス'` |
| レビュー | `claude "/go-review exercises/ch03"` | `codex '$go-review exercises/ch03'` |

役の定義は `.agents/skills/` にあり、`.claude/skills/` はそこへのシンボリックリンク。共通ルールは `AGENTS.md`。

- 章を読み始めた・読み終えたら、どの役にでも「ch03 を読書中にして（スライスまで）」と頼む。`docs/progress.md` に反映され、出題範囲の判断に使われる。
- Codex は初回起動時に「このディレクトリを信頼する」を選ぶ。信頼しないと `.codex/config.toml`（Go のビルドキャッシュを /tmp に置く設定）が読まれず、サンドボックス内で `go test` などが失敗する。

## よく使うコマンド

[go-task](https://taskfile.dev/) で登録している（`go install github.com/go-task/task/v3/cmd/task@latest` で入る）。`--` の後ろはタスクへの引数。

| コマンド | 内容 |
|---|---|
| `task ask` | 質問役を起動 |
| `task drill -- ch03 スライス` | 出題役を起動 |
| `task review -- exercises/ch03` | レビュー役を起動 |
| `task run -- ch03/01_hello` | `go run ./practice/ch03/01_hello` |
| `task test` | 全問のテスト（`task test -- ch03/01_slice_append` で1問だけ） |
| `task check` | `go vet` と `gofmt`（gofmt の差分があれば失敗） |

役の起動は `task ask AI=codex` のように付けると Codex で起動する。

## ディレクトリ

| パス | 中身 |
|---|---|
| `practice/` | 本を読みながら写経・実行するコード |
| `exercises/` | 練習問題（TODO 付きのコードとテスト） |
| `docs/progress.md` | 進捗とつまずきメモ |
| `lesson_go/` | 以前に本を途中まで進めたときの記録 |
