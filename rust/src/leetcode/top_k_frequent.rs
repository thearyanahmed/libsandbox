use std::collections::HashMap;

use crate::utils::{print_type_of, random_vector};

pub fn input() -> (Vec<i32>, i32) {
    let v : Vec<i32> = vec![1,1,1,2,2,3,3,3,3,33,3,3,4,0,0,0,0,0,100];
    (v, 2)
}

pub fn solution(nums: Vec<i32>, k: i32) -> Vec<i32> {
    let mut hashmap = HashMap::new();
    for num in nums {
        *hashmap.entry(num).or_insert(0) += 1;
    }
    let mut vec: Vec<(i32, i32)> = hashmap
        .into_iter()
        .collect();
    vec.sort_by(|(_, x), (_, y)| y.cmp(x));
    let mut res = Vec::new();

    for num in 0..k as usize {
        res.push(vec[num].0)
    }
    res
}