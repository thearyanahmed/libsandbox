use std::collections::HashMap;

pub fn solution(v: Vec<i32>) -> i32 {
    let mut hashmap = HashMap::<i32, i32>::new();

    for i in v {
        if hashmap.contains_key(&i) {
            hashmap.insert(i, 2);
        } else {
            hashmap.insert(i, 1);
        }
    }

    for (i, v) in hashmap {
        if v == 1 {
            return i;
        }
    }

    0
}
