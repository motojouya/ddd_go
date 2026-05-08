
# Development Guide

## Overview
このドキュメントは、開発に関するガイドラインの解説。  
作業順序についても言及している。  

## Design
モデリングを中心とした設計を行った後に実装する。  
設計ドキュメントは、docsディレクトリに配置する。  

設計ドキュメントについては[docs/guildeline/design.md](docs/guildeline/design.md)を参照。  

## Coding
実装は、特定の集約に絞って、基本的に以下の順序でpackageを実装していく。
1. core
2. entry
3. exit
4. schema
5. store
6. behavior
7. controller

1.coreは集約の中心で、他の集約のcoreに依存することはあれど、coreが集約配下の他のpackageに依存することはない。
また、2-4までは項目の定義を実装する傾向が強く、coreの情報に依存する。  
5-7は処理の流れを実装する傾向が強く、1-4の情報に依存する。  

一つの集約の実装をやりきることが、一つのタスクの基本的な目標となる。  

どの集約から実装していくかは、依存の少ないマスタ寄りの集約から着手していく。集約同士の依存関係は[docs/aggregate.md](docs/aggregate.md)を参照。  
ただし、controller,storeにおいては被依存の集約を参照することもあり得るので、それらについては、実装を保留し、被依存の集約ができたら実装する。  

実装時にどこに何を実装するか、全体概要をつかむ場合は[docs/guildeline/package.md](docs/guildeline/package.md)を参照。  
実装時の注意点や、実装ガイドラインは[docs/guildeline/coding/index.md](docs/guildeline/coding/index.md)を参照。  

## Testing
テストは、集約配下のcore,entry,exitについては、パッケージのファイルがそろったら一気にユニットテストを実装する。  
store,behavior,controllerについては、一つの関数を作成したら、対応するユニットテストを実装する。  
全ての集約の全てのパッケージが実装されたら統合テストを実装する。  

テストの実装方法、注意事項については[docs/guildeline/test.md](docs/guildeline/test.md)を参照。  

