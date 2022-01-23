func majorityElement(nums []int) int {
    var maps = make(map[int]int)
    
    
    n := len(nums)
    
    for i := 0; i < n; i++{
        //c := 0
        num := nums[i]
        maps[num]++
        
        if c, ok := maps[num]; ok {            
            if c > (n/2) {
                return num
            }
        }   
    }
    
    return 0
}