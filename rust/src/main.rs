#![allow(dead_code)]
#![allow(unused_variables)]

mod algoexpert;
mod practice;
mod ds;
mod leetcode;
mod utils;

use crate::leetcode::group_anagrams;

fn main() {
    println!("[+] Practice code");
    let a = group_anagrams::input();
    let b = group_anagrams::input();
    let results = group_anagrams::solution(a);
    let another_results = group_anagrams::another_solution(b);
    println!("{:?}",results);
    println!("{:?}",another_results)
}
