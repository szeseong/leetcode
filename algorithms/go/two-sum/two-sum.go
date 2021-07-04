package twoSum

func twoSum(nums []int, target int) []int {

    //mapNums := map[int]int{}
    mapNums := make(map[int]int)

    for firstNumIndex, firstNum := range nums {
        secondNum := target - firstNum

        if secondNumIndex, ok := mapNums[secondNum]; ok {
            return []int{firstNumIndex, secondNumIndex}
        }

        mapNums[firstNum] = firstNumIndex
    }

    return nil
}
