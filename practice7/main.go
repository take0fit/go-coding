package main

import (
	"fmt"
)

// 問題: 長方形の幅と高さを受け取り、その面積を計算する関数を作成してください。
func rectangleArea(width, height float64) (float64, error) {
	if width <= 0 || height <= 0 {
		return 0, fmt.Errorf("幅と高さは正の数でなければなりません")
	}
	return width * height, nil
}

func main() {
	area, err := rectangleArea(5, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("面積は", area, "です")
}
