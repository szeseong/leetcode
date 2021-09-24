/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    
    if len(nums) == 0 {
        return nil
    }
    
    mid := len(nums)/2
    
    
    ans := &TreeNode{
        Val: nums[mid],
    }
    
    if mid > 0 {
        left := nums[:mid]
        ans.Left = sortedArrayToBST(left)      
    }else{
        ans.Left = nil
        return ans
    }
    
    
    if mid < len(nums) {
        right := nums[mid+1:]
        ans.Right = sortedArrayToBST(right)
    }else{
        ans.Right = nil
        return ans
    }
   
 
    return ans
}

