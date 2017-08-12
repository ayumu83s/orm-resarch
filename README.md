# ORMの調査

チェックすること
1件select
複数件select
joinしてselect
トランザクション
CRUD

## xorm
	条件を設定した構造体をGet、Countなどに渡してあげる使い方をする
	でも基本はwhere設定した方が良い気がする

	テーブルの生成
	xorm reverse [-s] mysql "root:@tcp(localhost:3306)/sakila?charset=utf8" $GOPATH/go-xorm/cmd/xorm/templates/goxorm
	-s をつければ、struct.goが1個だけ生成される

	configとテンプレを格納したパスは自分で指定できる。出力先も。
	xorm reverse -s mysql "root:@tcp(localhost:3306)/sakila?charset=utf8" ./goxorm_tmpl ./xorm

## sqlboiler
	tomlを用意して以下。-oでoutputディレクトリ指定
	sqlboiler mysql -o sqlboiler -p sqlboiler_models

	parseTime=trueで接続しないとエラーでる。
	sqlboiler: unable to select from actor: failed to bind pointers to obj: sql: Scan error on column index 3: unsupported Scan, storing driver.Value type []uint8 into type *time.Time

	boil.DebugMode = true
	発行したクエリを確認できる
## gorp
## gorm
## genmai
