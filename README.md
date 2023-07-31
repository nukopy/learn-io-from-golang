# learn-io-from-golang

Zenn の本「[Go から学ぶ I/O](https://zenn.dev/hsaki/books/golang-io-package)」の写経リポジトリ

## 開発環境

- macOS 14.0 Sonoma
- Go 1.20.6

## 書籍の概要

引用元：[Go から学ぶ I/O：Chapter 01 - はじめに](https://zenn.dev/hsaki/books/golang-io-package/viewer/intro)

> この本では、Go 言語で扱える I/O についてまとめています。
>
> Go 言語 I/O を扱うためのパッケージとしては、ドンピシャのものとしては io パッケージがあります。
> しかし、例えば実際にファイルを読み書きしようとするときに使うのは、os パッケージの os.File 型まわりのメソッドです。
> 標準入力・出力を扱おうとすると fmt パッケージが手っ取り早いですし、また速さを求める場面では bufio パッケージのスキャナを使うということもあるでしょう。
> このように、「I/O」といっても Go でそれに関わるパッケージは幅広いのが現状です。
>
> また、ファイルオブジェクト f に対して `f.Read()` とかいう「おまじない」と唱えるだけで、なんでファイルの中身が取得できるの？一体裏で何が起こっているの？という疑問を感じている方もいるかと思います。
>
> ここでは
>
> - os や io とかいっぱいあるけど、それぞれどういう関係なの？
> - 標準入力・出力を扱うときに fmt と bufio はどっちがいいの？
> - そもそも bufio パッケージって何者？
> - Go でやった I/O まわりの操作は、実現のために裏で何が起こっているの？
>
> こういったことを一から解説していきます。

## 目次

- [x] Chapter 01：はじめに
- [x] Chapter 02：ファイルの読み書き
- [x] Chapter 03：ネットワーク
- [ ] Chapter 04：io パッケージによる抽象化
- [ ] Chapter 05：bufio パッケージによる buffered I/O
- [ ] Chapter 06：fmt で学ぶ標準入力・出力
- [ ] Chapter 07：bytes パッケージと strings パッケージ
- [ ] Chapter 08：おわりに

## 実行

いくつか実装したサンプルコードの実行方法

### ネットワーク I/O

TCP サーバ / クライアントの実行方法

```bash
# TCP サーバの起動
cd learn-io-from-golang/network_io/server
go run .

# 別のターミナルから TCP クライアントを起動
cd learn-io-from-golang/network_io/client
go run .
```
