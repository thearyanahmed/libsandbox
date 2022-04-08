mod leetcode;

fn main() {
    let v = vec![3,2,3];

    let r = leetcode::majority_element::solution(v);

    println!("{:?}", r);
}