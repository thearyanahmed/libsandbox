use std::collections::HashMap;

fn main() {
    let v = vec![1,2,2,3,3];

    let single_number = single_number(v);

    println!("{:?}", single_number);
}

fn single_number(v: Vec<i32>) -> i32 {
    // We can sort it.
    let mut hashmap = HashMap::<i32,i32>::new();

    for i in v {
        if hashmap.contains_key(&i) {
            hashmap.insert(i,2);
        } else {
            hashmap.insert(i,1);
        }
    }

    for (i,v) in hashmap {
        if v == 1 {
            return i;
        }
    }

    0
}