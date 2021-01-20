# Install

Blog: [Go言語を練習しようぜ☆（＾～＾）？](https://crieit.net/drafts/5ffc46af9214c)  

例えば `C:\go` に Go言語（本体）をインストールしてください。  
例えば 以下のように環境変数を設定してください。  

システム環境変数:  

| Name | Value                    | 別の書き方  |
| ---- | ------------------------ | ----------- |
| Path | C:\go\bin                |             |
| Path | C:\Users\むずでょ\go\bin | %GOPATH%\go |

ユーザー環境変数:  

| Name   | Value                | 別の書き方       |
| ------ | -------------------- | ---------------- |
| GOPATH | C:\Users\むずでょ\go | %USERPROFILE%\go |

以下のコマンドを叩いてください。  

```shell
go version
go version go1.15.6 windows/amd64

# go言語は、インストール時に、インストール先ディレクトリ（GOROOT）を覚えています。
go env GOROOT
c:\go
```

## Modules の使い方

```shell
go mod init github.com/muzudho/telnet-echo-server
```

## Packages

```shell
go get -v -u github.com/reiver/go-telnet
```
