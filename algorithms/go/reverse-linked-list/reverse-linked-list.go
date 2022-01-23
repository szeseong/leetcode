/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
   
    for {
        
        if head == nil {
            break
        }

        prev =  &ListNode{
           Val: head.Val,
           Next: prev,
        }
       
        
        head = head.Next
    }
    return prev
}