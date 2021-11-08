/*
 * @lc app=leetcode.cn id=1658 lang=golang
 *
 * [1658] 将 x 减到 0 的最小操作数
 */

// @lc code=start
func minOperations(nums []int, x int) int {
    left:=0
    right:=0
    sum:=0
    for i:=0;i<len(nums);i++{
        sum+=nums[i]
    }
    r:=0
    res:=9999999
    if  sum-x<0{
        return -1
    }
    for right<len(nums){
      
        
         r+=nums[right]
        right++
        if r==sum-x{
           res=min(res,left+len(nums)-right)
        }
        for left<=right{
            if r>sum-x{
             r-=nums[left]
             left++
            } else{
                break
            }
        }
           
         if r==sum-x{
           res=min(res,left+len(nums)-right)
        }
    }
        

    if res==9999999{
        return -1
    }
   return res 

}
func min(a,b int)int{
    if a<b{
        return a
    }
    return b
}
// @lc code=end

