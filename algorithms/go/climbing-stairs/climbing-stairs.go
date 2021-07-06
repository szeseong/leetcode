func climbStairs(n int) int {
    currentCount := 1
    previousCount := 0

    for i := 0; i < n; i ++{
       previousCount, currentCount = currentCount, previousCount + currentCount
    }
    
    return currentCount
}
