# helm

## やったこと

* テンプレートの生成
  > $ helm create [dirname]
* テンプレートの書式検証
  > $ helm lint [dirname] -f [dirname]/values.yaml --strict --debug
* マニフェストの出力
  > $ helm template [dirname] -f [dirname]/values.yaml
* コンテナのデプロイ
  > $ helm install [name] -f [dirname]/configs/[confname] [dirname]
* コンテナの更新
  > $ helm upgrade [name] -f [dirname]/configs/[confname] [dirname]
* コンテナの削除
  > $ helm uninstall [name]

## 参考

* https://helm.sh/ja/docs/intro/using_helm/
* https://helm.sh/ja/docs/howto/charts_tips_and_tricks/
* https://golang.hateblo.jp/entry/golang-text-html-template
