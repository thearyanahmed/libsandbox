pub fn input() -> Vec<i32> {
    Vec::from([555,901,482,1771])
}

pub fn solution(nums: Vec<i32>) -> i32 {
    let mut counter = 0;

    for a in nums {
        let mut d = 0;
        let mut num = a;

        while num > 0 {
            num = num / 10;
            d = d + 1;
        }

        if d % 2 == 0 {
            counter = counter + 1;
        }
    }

    counter
}
