
# Project Documentation

## Overview
プロジェクト全体でのシステム構成については、[docs/system.md](docs/system.md)を参照。  
このリポジトリは、上記システム構成の中の在庫管理システムにあたる。  

このリポジトリでは、集約ごとにディレクトリを切り、各集約の中でモジュールごとのディレクトリを切る構成をとる。  
各集約の中の構成は[docs/guideline/package.md](docs/guideline/package.md)を参照。  
集約と集約の関連については、[docs/aggregate.md](docs/aggregate.md)を参照。  
モデリングにおけるキー設計については[docs/key.md](docs/key.md)を参照。  

## Directory

トップレベルでのディレクトリ構成は、基本はGoの標準に倣いつつ、少し変更している。
pkg配下には各集約ごとにディレクトリに切るが、その詳細については、[docs/guideline/package.md](docs/guideline/package.md)に記載する。

- build/  
  ビルド設定、dockerなどのコンテナ設定など  
- cmd/  
  アプリケーションのエントリポイント  
  - server/main.go  
    webサーバのエントリーポイント  
  - tool/main.go  
    cliコマンドのエントリーポイント  
- docs/  
  ドキュメント全般。LLMも人も読める形で書く。  
- pkg/  
  - route.go  
    webサーバのurl route定義  
  - command.go  
    cliコマンドのエントリーポイント定義  
  - job.go  
    他処理から呼ばれる非同期処理のエントリーポイント定義  
  - ...each aggregate  
    各集約ごとの実装。詳細は後述  
- test/  
  単体テストではなく、統合テストの実装。runnで実行するyamlファイルを配置する。  
- tmp/  
  一時的なファイル。migrationファイルは各集約のschemaディレクトリに配置されるが、実行時は集める必要があるため、ここにコピーされる。gitignoreされているため、ローカルでの開発でのみ使用される。  

## Development
開発の進め方については、[docs/guideline/index.md](docs/guideline/index.md)を参照

