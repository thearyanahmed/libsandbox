pub fn input() -> Vec<i32> {
    // Vec::from([1,0,2,3,0,4,5,0])
    Vec::from([1,2,3])
}

pub fn solution(arr: &mut Vec<i32>) {
    let mut i = 0;

    let length = arr.len();

    while i < length - 1 {
        if arr[i] == 0 {
            // shift right
            // rotate from the last
            let mut j = length - 1;

            while i + 1 < j {
                arr[j] = arr[j - 1];
                j = j - 1;
            }

            arr[i + 1] = 0;
            i = i + 1;
        }

        i = i + 1;
    }

    println!("{:?}",arr);

    ()
}