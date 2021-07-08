/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func maxDepth(root *TreeNode) int {
    
    if root == nil {
      return 0        
    }  
    
    return walk(root, 1)
}

func walk(root *TreeNode,currentDepth int) int{
    if root == nil {
        return currentDepth
    }
    
    leftDepth := currentDepth
    rightDepth :=  currentDepth
    
    if root.Left != nil {
        leftDepth = walk(root.Left, currentDepth) + 1
    }  
    
    if root.Right != nil {
        rightDepth = walk(root.Right, currentDepth) + 1
    }  
    
    if leftDepth > rightDepth {
        return leftDepth
    } else{
        return rightDepth
    }
    
}

