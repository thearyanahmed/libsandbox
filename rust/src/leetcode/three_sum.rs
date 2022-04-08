pub fn solution(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
    let mut res = Vec::new();
    nums.sort();

    for i in 0..nums.len() {
        if i > 0 && nums[i] == nums[i - 1] {
            continue;
        }

        let mut left = i + 1;
        let mut right = nums.len() - 1;

        while left < right {
            let sum = nums[i] + nums[left] + nums[right];
            if sum > 0 {
                right -= 1;
            } else if sum < 0 {
                left += 1
            } else {
                res.push(vec![nums[i], nums[left], nums[right]]);
                left += 1;

                while nums[left] == nums[left - 1] && left < right {
                    left += 1;
                }
            }
        }
    }

    return res;
}