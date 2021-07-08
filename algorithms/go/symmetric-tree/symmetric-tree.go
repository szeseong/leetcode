/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
 
    return areTwoNodesSymmetric(root.Left, root.Right)
}

func areTwoNodesSymmetric(left *TreeNode, right *TreeNode) bool {
    if left == nil && right == nil {
        return true
    }
    
    if left == nil || right == nil {
        return false
    }
    
    if left.Val != right.Val {
        return false
    } 
    
    if areTwoNodesSymmetric(left.Left, right.Right) == false {
        return false
    }
    
    if areTwoNodesSymmetric(left.Right, right.Left) == false {
        return false
    }
    
    // go back one level
    return true
    
}
