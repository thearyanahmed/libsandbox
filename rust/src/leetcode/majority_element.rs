pub fn solution(nums: Vec<i32>) -> i32 {

    let mut nums = nums;
    nums.sort();

    return nums[nums.len() / 2];
    //
    // let mut major = nums[0];
    // let mut count = 0;
    //
    // for (i) in nums.iter().enumerate() {
    //     if i == 0 {
    //         continue;
    //     }
    //
    //     if count == 0 {
    //         count += 1;
    //         major = nums[i];
    //     } else if major == nums[i] {
    //         count += 1;
    //     } else {
    //         count -= 1;
    //     }
    // }
    //
    // major
}