# system_programming_go

## VSCode Shortcuts

* step-over =>	F8
* step-in	=> F7
* step-in	=> Shift + F8
* definition => F12

## godoc

```bash
# `-analysis type` をつけるとインタフェース分析をしてくれる => TopのPackage => implementsと言うカテゴリが追加される
godoc -http ":6060" -analysis type
# 依存パッケージが多い時用
GOPATH=/ godoc -http ":6060" -analysis type
```