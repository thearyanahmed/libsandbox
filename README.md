# LibSandbox

This is a repository where I put random code, whether it's problem solving or an article about kubernetes or something else. This is called LibSandbox, lib indicating library and sandbox for tinkering with anything and everything that I find interesting.

If anything interests you, feel free to use it/ clone it/update it. 

### Postgres
- [PostgreSql](postgres/README.md)

### Redis
- [Redis](redis/)
  - [Redis Clustering](redis/README.md)

### Kubernetes
- [Kubernetes](/kubernetes/)
  - [Cluster architecture and general notes](kubernetes/README.md)

### Problem Solving
r/797022

- [leetcode](rust/src/leetcode)
  - [remove nth element](go/linked_list/19_remove_nth_element_from_the_last.go)
  - [LRU cache](go/linked_list/146_LRU_cache.go)
  - [LRU cache 2](go/linked_list/146_2_LRU_cache.go)
  - [Merge in between](go/linked_list/1669_merge_in_between.go)
  - [Swapping nodes](go/linked_list/1721_swaping_nodes.go)
  - [Contains duplicate](rust/src/leetcode/contains_duplicate.rs)
  - [Duplicate zeros](rust/src/leetcode/duplicate_zeros.rs)
  - [Find numbers with even number of digits](rust/src/leetcode/find_numbers_with_even_number_of_digits.rs)
  - [Group anagram](rust/src/leetcode/group_anagrams.rs)
  - [Majority elements](rust/src/leetcode/majority_element.rs)
  - [Max consecutive ones](rust/src/leetcode/max_consecutive_ones.rs)
  - [Merge sorted array](rust/src/leetcode/merge_sorted_array.rs)
  - [Single number](rust/src/leetcode/single_number.rs)
  - [Three sum](rust/src/leetcode/three_sum.rs)
  - [Squares of a sorted array](rust/src/leetcode/squares_of_a_sorted_array.rs)
  - [Top K Frequent](rust/src/leetcode/top_k_frequent.rs)
    - **solution**
      - Count the number of occurrence and put it in a hashmap
      - Sort the values by converting it to a vec (rust)
      - Take the first k elements, so `0..k` to create the final result 
  - [Valid anagram](rust/src/leetcode/valid_anagram.rs)
    - **solution**: 
      - Check the length of characters
      - Count occurrence of characters and putting it inside hashmap.
      - Match the counts, character by character
    - **another solution**:
      - Check the length of characters
      - Sort the strings in place. Do a linear search and match on both scenarios.
  - [Merge in between](go/linked_list/leetcode_1669_merge_in_between.go) 
  - [Remove Nth Node From End of List](go/linked_list/19_remove_nth_element_from_the_last.go) 
  - [Swapping Nodes in a Linked List](go/linked_list/1721_swaping_nodes.go) 
  - [Square root](go/binary_search/square_root.go)
    - **solution**
      - take left = 0 and right = target
      - binary search, take `mid = left + ((right - left) / 2))` and check 
      - mid * mid is left or equals to `target`
      - if yes -> move left -> mid + 1, res = mid 
      - else -> move right -> mid - 1 
