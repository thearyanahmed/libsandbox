mod leetcode;

fn main() {
    let v = vec![1,2,2,3,3];

    let r = leetcode::single_number::solution(v);

    println!("{:?}", r);
}