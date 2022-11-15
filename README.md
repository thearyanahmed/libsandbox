# LibSandbox

This is a repository where I put random code, whether it's problem solving or an article about kubernetes or something else. This is called LibSandbox, lib indicating library and sandbox for tinkering with anything and everything that I find interesting.

If anything interests you, feel free to use it/ clone it/update it. 

### Postgres
- [PostgreSql](postgres/README.md)

### Kafka
- [Building Scalable Data Pipelines with Kafka](kafka/README.md)

### Redis
- [Redis Clustering](redis/README.md)

### Kubernetes
- [Cluster architecture and general notes](kubernetes/README.md)

### Data structure & Algorithms
**r/781691**

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
      - Take left = 0 and right = target
      - Binary search, take `mid = left + ((right - left) / 2))` and check 
      - mid * mid is left or equals to `target`
      - If yes -> move left -> mid + 1, res = mid 
      - Else -> move right -> mid - 1 
  - [Search in 2d matrix](go/binary_search/search_in_2d_matrix.go)
    - **solution** 
      - Use simple logic, iterate over the array
      - Check if the right row's value is greater than the target
      - If yes, increase `left++` and `continue`
      - Else, do a linear search in the current one
    - **another solution** 
      - Iterate over the array and find the boundery where `row[0] <= target <= row[last]`
      - Hold the value of the row in a variable
      - After the first loop, check if value was in boundery, if not found, return false
      - Else do a simple binary search within the row
  - [Koko eating bananas](go/binary_search/koko_eating_bananas.go)
    - **solution**
      - Loop for minimum to maximum
      - Set mid, calculate total hours to eat banans assuming mid is the result
      - Calculation formula `piles.forEach (p) => p / k` 
      - Check if calculated hours is less than target h
        - Yes: set `r = k - 1` , set res's new value as `min(res,k)`
        - No: set `l = k + 1`
   - [Product of array except self](go/array/product_of_array_except_self.go)
     - **solution O(n)**
       - Multiplying anything with 1 doesn't change the value
       - Imagine there is a prefix and postfix for each element (including first and last)
       - Create an array that contains prefixs
       - Then multiply that array with postfixs
   - [Longest consecutive sequence](go/array/longest_consecutive_sequence.go)
     - **solution**
       - Put the array in a hashmap
       - Skip if it has `n-1` in hashmap. So if n = 100, if it has 99, skip
       - Else, count until it doesn't have any sequenced number, `n = $n, while map[$n] n++`
       - Take `max(oldLongest, newCount)`
   - [Implement Time based key value store](go/binary_search/time_based_key_value_store.go)
   - [Is valid palindrome I](go/two_pointers/valid_palindrome_i.go)
     - **solution O(n)**
       - Use two pointers
   - [Is valid palindrome II](go/two_pointers/valid_palindrome_ii.go)
     - **solution O(n)**
       - Use two pointers, during the check, if it is not a palindrom, return s, and l,r
       - Using those l, rs first run from `l + 1` until `r`
       - If thats still false, run `l` to `r-1` 
   - [Two Sum](go/binary_search/two_sum.go)
     - **solution**
       - Do a binary search
       - Start with `left = i + 1, right = len - 1`
       - Binary Search `mid = left + ((right - left) /2 )`
       - We are looking for `t = target - numbers[i]`
       - If `array[mid] = t` thats result, `i and mid`
       - If less, `left = mid + 1`
       - ELse, `right = mid - 1` 
   - [Container with most water](go/two_pointers/container_with_most_water.go)
     - **solution**
       - Use two pointer
       - Calculate the area
       - Take the max from `max(area,res)`, res is 0 at first
       - If height[left] is smaller than height[right]
         - increase left
         - else, decrease right
   - [Trapping rain water](go/two_pointers/trapping_rain_water.go)
   - [Reverse Linked List - I](go/linked_list/reverse_linked_list_i.go)
     - **solution**
       - take `prev *List`
       - while head exists, set
         - `head.Next = prev`
         - `prev = head` (current)
         - `head (current) = head.Next` (move forward)
 - [Array](go/array)
    - [Rotate array by N](go/array/rotate_array.go)
    - [Reverse array](go/array/reverse_array.go)
    - [Rotate counter clockwise](go/array/rotate_counter_clock_wise.go)
    - [Cyclic Rotation](go/array/cyclic_rotate.go)
    - [Bubble sort O(n^2)](go/array/bubble_sort.go)
    - [Merge sort O(nLogn)](go/array/merge_sort.go)
    - [Insertion sort O(n^2)](go/array/insertion_sort.go)
    - [Selection sort O(n^2)](go/array/selection_sort.go)


