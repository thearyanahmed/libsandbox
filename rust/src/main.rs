#![allow(dead_code)]
#![allow(unused_variables)]

mod algoexpert;
mod practice;
mod ds;
mod leetcode;
mod utils;

use crate::leetcode::{kth_largest_element};

fn main() {
    println!("[+] Practice code");
    let (a,b) = kth_largest_element::input();
    let results = kth_largest_element::solution(a,b);

    println!("final output: {:?}",results);
}
