# 独自合意形成アルゴリズム `HEISEI` について

## 独自合意形成アルゴリズム `HEISEI` とは
- 独自合意形成アルゴリズム `HEISEI` とは、PoW (https://en.wikipedia.org/wiki/Proof-of-work_system) をベースに作成したEvilEye独自の合意形成アルゴリズム
- PoWでは、ハッシュ値の先頭が `0` のN個の連番にならなければならないところを、HEISEIアルゴリズムでは、平成初頭の日本のインターネット界隈を牽引してきた素晴らしき文字コードであるEUC_JPでの `"平"`, `"成"` がビットパターンとして存在していればブロック作成の権利を得る

## 合意形成のフロー
大まかには以下のようなフローでブロック生成のための合意形成を行う

- 各ノードはバックグラウンドでひたすら計算をする
- 計算が成功したら、各ノードのgRPCメソッド、 `SuccessHashCalc` を呼び、合意形成がちゃんとできるかを他のノード全てに再確認(再計算)してもらう.
- 各ノードは再計算が終了したら、結果をリクエストの`is_ok` に込めて各ノードのgRPCメソッド、 `SendCheckResult` を呼び、他のノード全てに結果を送信する
- 計算を成功させたノード数が、 `N(全ノード数)/2` を上回ったら、そのリクエストのnonceで各ノードでブロックを生成する

以下で図とともに詳細を解説していく  
![](images/consensus_1.jpg)
- 各ノードはバックグラウンドでひたすらハッシュ計算をし、ブロック作成可能なハッシュ値を生成できたら、各ノードの `SuccessHashCalc` (gRPCのメソッド)を叩き、各ノードに伝搬する
- 各ノードは `SuccessHashCalc` でNonce、RequestID、を受け取り、手元で再計算し、計算が終了したら `SendCheck` (gRPCのメソッド) を叩き、各ノードに成功か否かを伝える

![](images/consensus_2.jpg)
- ノードは `SuccessHashCalc` を受け取ったとき、リクエスト内部の `is_ok` をRequestIDごとにカウントし、ハッシュ計算の成功のカウントが予め用意してあるノード数に応じたしきい値を超えた場合に初めてブロックを生成する

