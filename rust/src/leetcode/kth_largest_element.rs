use std::collections::HashMap;
use crate::utils::print_type_of;

pub fn input() -> (Vec<i32>, i32) {
    let v : Vec<i32> = vec![3,2,3,1,2,4,5,5,6];
    (v, 4)
}

pub fn solution(nums: Vec<i32>, k: i32) -> Vec<i32> {
    println!("nums {:?}", nums);
    let mut hashmap = HashMap::new();
    for num in nums {
        *hashmap.entry(num).or_insert(0) += 1;
    }
    let mut vec: Vec<(i32, i32)> = hashmap
        .into_iter()
        .collect();
    vec.sort_by(|(_, x), (_, y)| y.cmp(x));
    let mut res = Vec::new();

    let a = vec[(k - 1) as usize];

    println!("kth largest element: {:?}",a);

    println!("vec -> {:?}, k -> {}",  vec, k);

    for num in 0..k as usize {
        res.push(vec[num].0)
    }
    res
}