struct Solution {}
use std::collections::HashSet;

impl Solution {
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        let mut set: HashSet<i32> = nums.iter().cloned().collect();

        let mut longest_consecutive = 0;

        for number in nums {
            if !set.contains(&(number - 1)) {
                let mut current = 0;
                while set.contains(&(number + current)) {
                    current += 1;
                }

                longest_consecutive = longest_consecutive.max(current);
            }
        }
        longest_consecutive
    }
}

fn main() {}
