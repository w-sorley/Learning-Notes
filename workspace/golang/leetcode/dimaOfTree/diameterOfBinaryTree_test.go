package dimaoftree

import (
	"fmt"
	"strconv"
	"testing"
)

func TestDimaOfTree(t *testing.T) {
	root := TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	// left := TreeNode{
	// 	Val: 0,
	// }
	// root.Left = &left
	// right := TreeNode{
	// 	Val: 0,
	// }
	// root.Right = &right

	ret := diameterOfBinaryTree(&root)
	expected := 0
	fmt.Println("ret:" + strconv.Itoa(ret))
	if ret != expected {
		t.Error("failed!")
	}

}
