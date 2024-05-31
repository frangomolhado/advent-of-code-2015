fn part1(input: &str) -> i32 {
    input.chars().fold(0, |acc, c| {
        if c == '(' {
            acc + 1
        } else if c == ')' {
            acc - 1
        } else {
            acc
        }
    })
}

fn part2(input: &str) -> usize {
    let mut floor = 0;
    for (i, c) in input.chars().enumerate() {
        if floor < 0 {
            // already increased by 1 when the next interation starts
            return i;
        } else if c == '(' {
            floor += 1;
        } else if c == ')' {
            floor -= 1;
        }
    }

    unreachable!("Shouldn't get to this point")
}

fn main() {
    let input = include_str!("input.txt");

    let part1_result = part1(input);
    println!("part 1: {}", part1_result);

    let part2_result = part2(input);
    println!("part 2: {}", part2_result);
}
