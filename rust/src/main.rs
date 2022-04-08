mod leetcode;

fn main() {
    let v = vec![-1,0,1,2,-1,-4];

    let r = leetcode::three_sum::solution(v);

    println!("{:?}", r);
}