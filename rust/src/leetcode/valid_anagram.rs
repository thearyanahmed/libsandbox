use std::collections::HashMap;

pub fn input() -> (String, String) {
    // (String::from("anagram"), String::from("nagaram"))
    (String::from("art"), String::from("bat"))
}

pub fn solution(s: String, t: String) -> bool {
    let mut first_map: HashMap<char, i32> = HashMap::new();

    for c in s.chars() {
        if first_map.contains_key(&c) {
            let v = *first_map.get(&c).unwrap() + 1;
            first_map.insert(c,v);
        } else {
            first_map.insert(c, 1);
        }
    }

    let mut second_map: HashMap<char, i32> = HashMap::new();

    for c in t.chars() {
        if second_map.contains_key(&c) {
            let v = *second_map.get(&c).unwrap() + 1;
            second_map.insert(c,v);
        } else {
            second_map.insert(c, 1);
        }
    }

    if first_map.len() != second_map.len() {
        return false;
    }

    for (k, v) in first_map {
        if ! second_map.contains_key(&k) {
            return false;
        }

        if v != second_map[&k] {
            return false;
        }
    }

    true

}