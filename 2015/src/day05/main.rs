use std::collections::{HashMap, HashSet};

fn part1(input: &str) -> usize {
    input
        .lines()
        .fold(0, |acc, s| if is_nice1(s) { acc + 1 } else { acc })
}

fn is_nice1(s: &str) -> bool {
    let vowels = ['a', 'e', 'i', 'o', 'u'];
    let disallowed_strs = [('a', 'b'), ('c', 'd'), ('p', 'q'), ('x', 'y')];

    let mut seen_vowels: HashMap<char, u32> = HashMap::new();
    let mut last_char = ' ';
    let mut double_letter = false;

    for c in s.chars() {
        if vowels.contains(&c) {
            seen_vowels.entry(c).and_modify(|v| *v += 1).or_insert(1);
        }

        if last_char == c {
            double_letter = true;
            continue;
        }

        if disallowed_strs
            .iter()
            .filter(|(prev, curr)| prev == &last_char && curr == &c)
            .count()
            == 1
        {
            return false;
        }

        last_char = c;
    }

    seen_vowels.into_values().sum::<u32>() >= 3 && double_letter
}

fn part2(input: &str) -> usize {
    input
        .lines()
        .fold(0, |acc, s| if is_nice2(s) { acc + 1 } else { acc })
}

fn is_nice2(s: &str) -> bool {
    let mut letters: HashMap<char, (usize, usize)> = HashMap::new();
    let mut pairs = HashSet::new();
    let mut overlap = false;
    let mut last_char = ' ';
    let mut repeated_with_between = false;
    let mut has_pair = false;

    for (i, c) in s.chars().enumerate() {
        if !has_pair {
            if last_char == c && !overlap {
                overlap = true;

                if !pairs.insert((last_char, c)) {
                    has_pair = true;
                }
            } else if last_char == c && overlap {
                overlap = false;
            } else {
                overlap = false;

                if !pairs.insert((last_char, c)) {
                    has_pair = true;
                }
            }
        }

        if let Some(v) = letters.get_mut(&c) {
            if i >= 3 && (v.0 == i - 2 || v.1 == i - 2) {
                repeated_with_between = true;
            } else {
                v.0 = v.1;
                v.1 = i;
            }
        } else {
            letters.insert(c, (0, i));
        }

        if repeated_with_between && has_pair {
            return true;
        }

        last_char = c;
    }

    false
}

fn main() {
    let input = include_str!("input.txt");

    let part1_result = part1(input);
    println!("part1: {}", part1_result);

    let part2_result = part2(input);
    println!("part2: {}", part2_result);
}
