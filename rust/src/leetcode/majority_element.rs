pub fn solution(nums: Vec<i32>) -> i32 {
    let mut nums = nums;
    nums.sort();

    nums[nums.len() / 2]
}