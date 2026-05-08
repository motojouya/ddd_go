
# Coding Guide Overview

## Overview
このドキュメントは、実装に関するガイドラインの解説。  
主に特定のpackageにおいて、どのような役割で、どのように実装するかについて解説する。  
packageの全体構成、概要については[docs/guildeline/package.md](docs/guildeline/package.md)を参照。  

## Packages
特定の集約配下のディレクトリ構成は以下となるが、詳細は`docs/guildeline/coding`配下のドキュメントを参照。  
companyは例であって、任意の集約名が入る。  

```
company/
  route.go
  command.go
  job.go
  core/
  entry/
  exit/
  record/
  store/
  behavior/
  controller/
  schema/
```

### (route|command|job).go
controllerで定義した機能を、外部から呼び出すためのエントリーポイントを定義する。
それぞのファイルは以下の役割を持つ。

- route.go
  web serverのURL Routingを定義する。
- command.go
  cli commandの定義をする。
- job.go
  非同期処理のエントリーポイントの定義をする。

### core package
いわゆるドメインモデルであり、副作用のないロジックやビジネスルールを定義する。
詳細は[docs/guildeline/coding/core.md](docs/guildeline/coding/core.md)を参照。

### entry package
システムへの入力を定義し、coreに変換されるデータを定義する。
詳細は[docs/guildeline/coding/entry.md](docs/guildeline/coding/entry.md)を参照。

### exit package
システムからの出力を定義し、coreから変換されるデータを定義する。
詳細は[docs/guildeline/coding/exit.md](docs/guildeline/coding/exit.md)を参照。

### record package
DBのレコードを定義する。coreに書き込んで足りる場合は、coreで完結させ、recordは用意しない。
詳細は[docs/guildeline/coding/record.md](docs/guildeline/coding/record.md)を参照。

### store package
DBへのアクセスを定義し、behaviorから呼ばれる。joinがない、条件が単純なクエリはbehaviorで完結させ、storeは用意しない。
詳細は[docs/guildeline/coding/store.md](docs/guildeline/coding/store.md)を参照。

### behavior package
集約としてのデータの更新、参照の処理を定義する。
詳細は[docs/guildeline/coding/behavior.md](docs/guildeline/coding/behavior.md)を参照。

### controller package
複数の集約のbehaviorを取りまとめて、外部から呼ばれる機能を定義する。複数の集約を取り扱うが、中心となる集約があるので、その集約に配置される。
詳細は[docs/guildeline/coding/controller.md](docs/guildeline/coding/controller.md)を参照。

### schema package
DBのスキーマを定義するマイグレーションファイルを配置する。  
詳細は[docs/guildeline/coding/schema.md](docs/guildeline/coding/schema.md)を参照。

