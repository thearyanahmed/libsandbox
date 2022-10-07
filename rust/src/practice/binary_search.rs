use crate::utils::random_vector;

pub fn input() -> Vec<i32> {
    let v : Vec<i32 > = random_vector(0, 20, 20);

    v
}

pub fn solution(mut array: Vec<i32>, target: i32) -> Option<i32> {
    array.sort();

    let mut left : i32 = 0;
    let mut right : i32 = (array.len() - 1) as i32;

    while left <= right {
        let mid = left + (right - left) / 2;

        if array[mid as usize] == target {
            return Some(mid);
        }

        if array[mid as usize] < target {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }

    None
}