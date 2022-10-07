pub fn input() -> Vec<i32> {
    Vec::from([1,2,3,1])
}

pub fn solution(arr: Vec<i32>) -> bool{
    let mut arr = arr;
    let length = arr.len();

    if length == 0 || length == 1 {
        return false
    }

    if length == 2 {
        return arr[0] == arr[1];
    }

    arr.sort();

    let mut leader = 1;
    let mut follower = 0;

    while leader < length {
        if arr[leader] == arr[follower] {
            return true;
        }

        leader += 1;
        follower += 1;
    }

    false
}