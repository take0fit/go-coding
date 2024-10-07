package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serialize(root *TreeNode) string {
	var sb strings.Builder
	var build func(node *TreeNode)
	build = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val) + ",")
		build(node.Left)
		build(node.Right)
	}
	build(root)
	return sb.String()
}

func deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if vals[0] == "null" {
			vals = vals[1:]
			return nil
		}
		val, _ := strconv.Atoi(vals[0])
		vals = vals[1:]
		return &TreeNode{
			Val:   val,
			Left:  build(),
			Right: build(),
		}
	}
	return build()
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}

	data := serialize(root)
	fmt.Println("シリアライズ結果:", data)

	deserializedRoot := deserialize(data)
	fmt.Println("デシリアライズ結果:", serialize(deserializedRoot))
}
