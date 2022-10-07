use std::vec::Vec;

// Write a function that takes in a non-empty array of distinct integers and an
// integer representing a target sum. The function should find all triplets in
// the array that sum up to the target sum and return a two-dimensional array of
// all these triplets. The numbers in each triplet should be ordered in ascending
// order, and the triplets themselves should be ordered in ascending order with
// respect to the numbers they hold.

// If no three numbers sum up to the target sum, the function should return an
// empty array.

pub fn input() -> (Vec<i64>, i64) {
    let array = vec![12,3,1,2,-6,5,-8,6];
    let target = 0;

    (array, target)
}

pub fn solution(mut array: Vec<i64>, target: i64) -> Vec<Vec<i64>> {
    array.sort();

    let len = array.len();
    let mut result = vec![];

    for i in 0..(len-2) {
        let mut left = i + 1;
        let mut right = len - 1;

        while left < right {
            let current = array[i] + array[left] + array[right];

            if current == target {
                result.push(vec![array[i],array[left],array[right]]);

                left = left + 1;
                right = right - 1;

            } else if current < target {
                left = left + 1;
            } else if current > target {
                right = right - 1;
            }
        }
    }

    result
}