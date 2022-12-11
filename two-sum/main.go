func twoSum(nums []int, target int) []int {
    var u []int
    for i:=0;i<len(nums);i++{
        for j:=i+1;j<len(nums);j++{
            if nums[i]+nums[j]==target{
                u=append(u,i,j)
                return u
            }
            
        }
    }
    return u
}