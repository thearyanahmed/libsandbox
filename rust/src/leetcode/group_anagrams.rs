use std::collections::HashMap;

pub fn input() -> Vec<String> {
    vec![
        // "eat".to_string(),
        // "tea".to_string(),
        // "tan".to_string(),
        // "ate".to_string(),
        // "nat".to_string(),
        // "bat".to_string(),
        "bdddddddddd".to_string(),
        "bbbbbbbbbbc".to_string(),
    ]
}

pub fn another_solution(strs: Vec<String>) -> Vec<Vec<String>> {
    let mut res = vec![];

    let a_ascii = 'a' as u8;

    let mut k: HashMap<String, Vec<String>> = HashMap::new();

    for s in strs {
        let count: [i32; 26] = [0; 26];
        let mut count = count.to_vec();

        for c in s.chars() {
            count[(c as u8 - a_ascii) as usize] += 1;
        }

        let count : String = count.iter().map(|e| e.to_string()).collect();

        if k.contains_key(&count) {
            let v = k.get_mut(&count).unwrap();
            v.push(s.to_string());
        } else {
            k.insert(count,vec![s]);
        }
    }

    for a in k {
        res.push(a.1);
    }


    res
}

pub fn solution(strs: Vec<String>) -> Vec<Vec<String>> {
    let mut h = HashMap::new();

    for s in strs {
        let mut key: Vec<char> = s.chars().collect();
        key.sort();
        h.entry(key).or_insert(vec![]).push(s);
    }

    h.values().cloned().collect()
}