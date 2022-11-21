use std::collections::HashMap;

struct Solution {}

impl Solution {
    pub fn top_k_frequents(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut map: HashMap<i32, i32> = HashMap::new();

        for num in nums {
            *map.entry(num).or_insert(0) += 1;
        }

        let mut x: Vec<(&i32, &i32)> = map.iter().collect();
        x.sort_by(|&a, &b| (b.1).cmp(&a.1));

        let mut result = vec![];

        for i in 0..k {
            result.push(*x[i as usize].0);
        }

        result
    }
}

fn main() {
    Solution::top_k_frequents(vec![4, 1, -1, 2, -1, 2, 3], 2);
}
