
# Controller Package Guide

## Overview
複数の集約のrepositoryを取りまとめて、外部から呼ばれる機能を提供するパッケージ。  
アプリケーションのインターフェースとしての役割を担う。  
複数の集約を取り扱うが、中心となる集約があるので、その集約に配置される。  

## 資料
`docs/draft`配下の対応する集約のドキュメントを参照。  
controller欄には、実装のイメージが記載されているが、必ずしも引数やエラー処理が完全に記載されているわけではないので、あくまで参考。  
ドキュメントの読み方については、[docs/guildeline/design.md](docs/guildeline/design.md)を参照。  

## 実装
関数ごとに構造体を1つのファイルに定義し、実装する。  

repositoryと違い、**1関数1構造体**という形をとる。  
controllerの関数ごとに、依存しているpkgが違うため。  

ファイル名は、`<FunctionName>.go`とする。構造体名は、`<FunctionName>Control`とする。  

### 構造体とファクトリ関数

また、各controllerは2種類のファクトリ関数を持つ

- `New<処理名>Control`
  依存を受け取ってセットするだけのコンストラクタ。テスト時にはこっちを呼び出す  
- `Create<処理名>Control`
  DB接続やLocalerの初期化を含み、利用する集約のrepositoryを初期化する実際のコンストラクタ。構造体は作らず、`NewXxxControl`を呼び出す

構造体のフィールドは、基本的にinterface型で、実装に依存しない形とする。  

### 関数実装
以下、注意点。

更新系処理では`databaseController.Transact`でトランザクションを開始する。
`CompanyContext`や`WarehouseBaseContext`は`Transact`の内側で使用する

controllerの返り値はcoreではなく、exitパッケージの型を使用するため、exitの変換関数を利用してcoreの構造体を変換する。  
exitは他の集約のexit構造体を持つこともあるので、`pkg/basic/core`のRelate,RelateUnique関数に、exitのRelate関数を渡す形で紐づけする。

### その他
Webであれば、controllerは`pkg/basic/controller/handle.go`の`Hand`関数に渡して使用する。  
routingに紐づけられる。

