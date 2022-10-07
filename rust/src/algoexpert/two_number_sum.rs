use std::vec;

// Write a function that takes in a non-empty array of distinct integers and an
// integer representing a target sum. If any two numbers in the input array sum
// up to the target sum, the function should return them in an array, in any
// order. If no two numbers sum up to the target sum, the function should return
// an empty array.

// Note that the target sum has to be obtained by summing two different integers
// in the array; you can't add a single integer to itself in order to obtain the
// target sum.

// You can assume that there will be at most one pair of numbers summing up to
// the target sum.

pub fn input() -> (vec::Vec<i64>, i64) {
    let array = vec![3,5,-4,8,9,1,-1,6];
    let target = 10;

    (array, target)
}

pub fn solution(array: vec::Vec<i64>, target: i64) -> vec::Vec<i64> {
    let len = array.len();

    for i in 0..len - 1 {
        for j in (i + 1)..len {
            if array[i] + array[j] == target {
                return vec![array[i],array[j]]
            }
        }
    }

    vec![]
}