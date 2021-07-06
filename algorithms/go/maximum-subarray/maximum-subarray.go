func maxSubArray(nums []int) int {
    UintSize := 32 << (^uint(0) >> 32 & 1)
    MaxInt  := 1<<((32 << (^uint(0) >> 32 & 1))-1) - 1 // 1<<31 - 1 or 1<<63 - 1
    MinInt  := -MaxInt - 1         // -1 << 31 or -1 << 63
    
    sum := 0
    max := MinInt
    
    for start := 0; start < len(nums); start++{
        sum = 0
        for i := start; i < len(nums); i++ {
            sum += nums [i]

            if max < sum {
                max = sum
            }
        } 
    }
     
    return max
}
