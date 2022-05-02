# Context

## 役割
- 処理の締め切り
- キャンセル信号伝播
- リクエストスコープ値の伝播

## 意義
一つの処理が複数の後ルーチンを跨いで行われる場合
### `./about/main.go`
- main関数が動いているメインゴルーチンは「リクエストがくるごとに、新しいゴルーチンを`go`文で立てる」
  - レスポンスを返すハンドラの処理は、別のゴルーチン上で行われる

例：
ハンドラで行う処理の中で、DB接続してデータを取得したい。
データ取得処理のために、別のゴルーチンを立てる。
https://go.dev/blog/concurrency-timeouts

## 複数個ゴルーチンが絡むとどうなるか
> 基本的に、Goでは「異なるゴールーチン間での情報共有は、ロックを使ってメモリを共有するよりも、チャネルを使った伝達を使うべし」という考え方を取っています。
> 並行に動いている複数のゴールーチン上から、メモリ上に存在する一つのデータにそれぞれが「安全に」アクセスできることを担保するのはとても難しいからです。

### 暗黙的なゴルーチンへの情報伝達

```go:main.go
type MyInfo int

func myFunc(ch chan MyInfo) {}

func main() {
	info := make(chan MyInfo)
	go myFunc(info)
}
```

### 拡張性の乏しさ
> また、上記のコードでは伝達する情報はMyInfo型と事前に決まっています。
> しかし、追加開発で、MyInfo型以外にもMyInfo2型という新しい情報も伝達する必要が出てきた」という場合にはどうしたらいいでしょうか。
> MyInfo型の定義をinterface{}型等、様々な型に対応できるようにする
> MyFunc関数の引数に、chan MyInfo2型のチャネルを追加するなどの方法が考えられますが、前者は静的型付けの良さを完全に捨ててしまっている・受信側で元の型を判別する手段がないこと、後者は可変長に対応できないことが大きな弱点です。
> このように、チャネルを使うことで伝達情報の型制約・数制約が入ってしまうことが、拡張を困難にしてしまっています。

### 伝達制御の難しさ
> main関数内にて、myFunc関数に渡したチャネルinfoをクローズすることで、myFuncが動いているゴールーチンにキャンセル信号を送信しています。
> この場合、MyFunc関数の中から起動されている3つのゴールーチンmyFunc2の処理はどうなってしまうでしょうか。
> これらも中断されるのか、それとも起動させたままにさせたいのか、3つとも同じ挙動をするのか、というところを正確にコントロールするには、引数として渡すチャネルを慎重に設計する必要があります。

```go:main.go
func myFunc2(ch chan MyInfo) {}

func myFunc(ch chan MyInfo) {
	go myFunc2(info1)
	go myFunc2(info2)
	go myFunc2(info3)
}

func main() {
	info := make(chan MyInfo)
	go myFunc(info)

	// 別のゴルーチンで実行しているmyFuncを中断させる
	close(info)
}
```

## contextの定義
context.Context型を確認してみる。
```
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```

```
func myFunc(ctx context.Context) {
	// ctxから、メインゴールーチン側の情報を得られる
	// (例)
	// ctx.Doneからキャンセル有無の確認
	// ctx.Deadlineで締め切り時間・締め切り有無の確認
	// ctx.Errでキャンセル理由の確認
	// ctx.Valueで値の共有
}

func main() {	
	var ctx context.Context
	ctx = (ユースケースに合わせたcontextの作成)
	go myFunc(ctx) // myFunc側に情報をシェア
}
```

