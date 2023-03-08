
impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let s: String = s
            .chars()
            .filter(|ch| ch.is_alphanumeric())
            .map(|ch| ch.to_ascii_lowercase())
            .collect();
        let half = s.len() / 2;
        let left = s[..half].chars();
        let right = s[(s.len() - half)..].chars();
        left.zip(right.rev()).all(|(a, b)| a == b)
    }
}

