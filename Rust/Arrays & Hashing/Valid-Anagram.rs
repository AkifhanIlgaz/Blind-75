impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }

        let mut count = [0; 26];

        for (s_char, t_char) in s.chars().zip(t.chars()) {
            count[(s_char as u8 - b'a') as usize] += 1;
            count[(t_char as u8 - b'a') as usize] -= 1;
        }

        count.iter().all(|x| *x == 0)
    }
}