fn is_palindrome(s: &str, left: usize, right: usize) -> bool {
    let mut left = left;
    let mut right = right;
    while left < right {
        if s.chars().nth(left) != s.chars().nth(right) {
            return false;
        }
        left += 1;
        right -= 1;
    }
    true
}

fn valid_palindrome_by_removing_one_element(s: &str) -> bool {
    let mut left = 0;
    let mut right = s.len() - 1;
    while left < right {
        if s.chars().nth(left) != s.chars().nth(right) {
            // Try removing either s[left] or s[right]
            return is_palindrome(s, left + 1, right) || is_palindrome(s, left, right - 1);
        }
        left += 1;
        right -= 1;
    }
    true
}

pub fn run() {
    let input = "abbaaaaa";
    let result = valid_palindrome_by_removing_one_element(input);
    println!("Result: {}",result);
}
