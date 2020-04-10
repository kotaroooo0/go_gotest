package tddbc

func Say(greeting string) string {
	return greeting + " TDD BootCamp!!"
}

// 問題文
// 整数閉区間を示すクラス（あるいは構造体）をつくりたい。
// 整数閉区間オブジェクトは下端点と上端点を持ち、文字列表現も返せる（例: 下端点 3, 上端点 8 の整数閉区間の文字列表記は "[3,8]"）。
// ただし、上端点より下端点が大きい閉区間を作ることはできない。
// 整数の閉区間は指定した整数を含むかどうかを判定できる。
// また、別の閉区間と等価かどうかや、完全に含まれるかどうかも判定できる。

// TODOLIST
// - 構造体の作成する関数の実装
//   - 上限
//   - 下限
// - 文字列表現を返すメソッドを実装
// - 上端点より下端点が大きい閉区間を作ることはできないようにする
// - 整数の閉区間は指定した整数を含むかどうかを判定する関数の実装
// - 別の閉区間と等価かどうかや、完全に含まれるかどうかも判定できる関数の実装

type IntClosedRange struct {
	upper int
	lower int
}

func NewIntClosedRange(upper, lower int) IntClosedRange {
	return IntClosedRange{
		upper: upper,
		lower: lower,
	}
}
