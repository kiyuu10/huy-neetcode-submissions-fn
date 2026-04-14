/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := dfs(root.Left)
    if left == -1 {
        return -1
    }

    right := dfs(root.Right)
    if right == -1 {
        return -1
    }

    if left-right > 1 || right-left > 1 {
        return -1
    }

    if left > right {
        return left + 1
    }

    return right + 1
}

func isBalanced(root *TreeNode) bool {
     return dfs(root) != -1
}
 
