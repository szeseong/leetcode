/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return root
    }

    
    if root.Left != nil && root.Right == nil {  
        root.Right = invertTree(root.Left) 
        root.Left = nil
    }  else if root.Left == nil && root.Right != nil {
        root.Left = invertTree(root.Right)  
        root.Right = nil
    }   else if root.Left != nil && root.Right != nil {

        l := invertTree(root.Left)        
        r := invertTree(root.Right)
        
        root.Left = r
        root.Right = l

    } 
    
    return root
    
}