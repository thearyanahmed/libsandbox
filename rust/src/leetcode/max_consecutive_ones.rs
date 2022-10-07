pub fn input() -> Vec<i32> {
    Vec::from([1,0,1,1,0,1])
}

pub fn solution(nums: Vec<i32>) -> i32 {
    let mut counter = 0;
    let mut max = 0;

    for i in nums {
        if i == 1 {
            counter = counter + 1;
        }

        if i == 0 {
            counter = 0;
        }

        if counter > max {
            max = counter;
        }
    }

    max
}
