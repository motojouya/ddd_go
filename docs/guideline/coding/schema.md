
# Schema Directory Guide

## Overview
DBのスキーマ定義を管理するディレクトリ。  
アプリケーションコードではなく、マイグレーションファイルを配置する。

## 資料
`docs/draft`配下の対応する集約のドキュメントを参照。  
対象のドキュメントのModel欄、Property欄、Note欄に記載がある。  
ドキュメントの読み方については、[docs/guildeline/design.md](docs/guildeline/design.md)を参照。  
集約の中のcore,recordパッケージはDBスキーマに対応する情報が定義されているので、それらも参照すること。  

## マイグレーションファイルの生成
`golang-migrate`を使用して以下のようにマイグレーションファイルを生成する。  

```bash
migrate create -ext sql -dir pkg/{aggregate}/schema {description}
```

上記でupとdownのSQLファイルが生成されるので、それぞれに適切なSQLを記述する。  

- up.sql（テーブル作成）
  テーブルの作成、カラムの追加、インデックスの作成などを記述。

- down.sql（ロールバック）
  up.sqlの逆の操作を記述。

## 命名規則

### テーブル名
- 単数形を使用（例: `item`, `company`, `warehouse`）
- スネークケースを使用

### カラム名
- スネークケースを使用
- ID系は `{entity}_id` の形式（例: `company_id`, `item_id`）
- フラグ系は真偽値を明確にする名前（例: `is_active`, `enabled`）

### インデックス名
- `idx_{table}_{column}` の形式（例: `idx_item_company_id`）
- 複合インデックスは `idx_{table}_{col1}_{col2}` の形式

### 外部キー制約名
- `fk_{table}_{ref_table}` の形式（例: `fk_item_company`）

### NOT NULL
Goの構造体上でポインタ型として定義されているフィールドは、DB上ではNULLを許容するカラムとなる。  
それ以外は、基本的にNOT NULL 制約を設定する。

