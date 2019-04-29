# EVIL EYE

## Required
- go 1.11.x

## How2Run
``` sh
$ make buildi && make docker-compose-without-pull
```

## EvilEyeとは!
![](images/anokuro.png)
- 黒歴史保存マシーンだぞ!
- 内部にブロックチェーンを埋め込んでいるのだ!

## EvilEyeのここがすごい！
### 黒歴史を改ざん不可能にして保存できるぞ!
![](images/chu2male.png)
- 黒歴史は誰もが消し去りたいもの...。そんな黒歴史をブロックチェーンに刻む事によって改ざん不可能にできるぞ！
- 複数のノードがデータを保存している(デフォルト3台)ので、誰かが抜け駆けして消したとしても多数決で正しいブロックチェーンを見つけれるぞ！
    - これでみんなの黒歴史を永遠に保存できるぞ！
- みんな誰もが第三者に適当な黒歴史を刻まれないように投票システムもあるぞ！これで誰もが安全な黒歴史ライフを送れるね！
### 独自合意形成アルゴリズム `HEISEI`
![](images/chu2female.png)
- 独自の合意形成アルゴリズム、`HEISEI` を使っているぞ！
    - ちなみにベースはPoWだぞ！
- この合意形成アルゴリズムは `SHA256(prevHash, data, hash)` で計算した結果のbit列に、特定のbitパターンが存在するかどうかで合意形成をするぞ!
    - ちなみにその特定のbitパターンは今は懐かしの平成初頭、インターネット黎明期によく使われていた文字エンコーディング、EUC_JPでの `"平"` `"成"` のいづれかだ！！
- アルゴリズムの詳細は[こちら](https://github.com/NoahOrberg/evileye/blob/feature/fix-readme/doc/consensus.md)
