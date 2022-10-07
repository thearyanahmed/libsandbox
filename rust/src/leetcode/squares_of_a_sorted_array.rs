pub fn input() -> Vec<i32> {
    Vec::from([-4,-1,0,3,10])
}

pub fn solution(mut data: Vec<i32>) -> Vec<i32>  {
    for num in data.iter_mut() {
        *num = *num * *num;
    }

    data.sort();

    data
}