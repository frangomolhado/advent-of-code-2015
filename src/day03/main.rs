use std::collections::HashSet;

#[derive(Eq, PartialEq, Hash, Clone)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn new(x: i32, y: i32) -> Self {
        Self { x, y }
    }
}

fn part1(input: &str) -> usize {
    let mut santa = Point::new(0, 0);
    let mut houses = HashSet::new();
    houses.insert(santa.clone());

    for c in input.chars() {
        if c == '^' {
            santa.y += 1;
        } else if c == 'v' {
            santa.y -= 1;
        } else if c == '<' {
            santa.x -= 1;
        } else if c == '>' {
            santa.x += 1;
        }

        houses.insert(santa.clone());
    }

    houses.len()
}

fn part2(input: &str) -> usize {
    let mut santa = Point::new(0, 0);
    let mut robo_santa = Point::new(0, 0);
    let mut houses = HashSet::new();
    houses.insert(santa.clone());

    for (i, c) in input.chars().enumerate() {
        if c == '^' {
            if i % 2 == 1 {
                santa.y += 1;
            } else {
                robo_santa.y += 1;
            }
        } else if c == 'v' {
            if i % 2 == 1 {
                santa.y -= 1;
            } else {
                robo_santa.y -= 1;
            }
        } else if c == '<' {
            if i % 2 == 1 {
                santa.x -= 1;
            } else {
                robo_santa.x -= 1;
            }
        } else if c == '>' {
            if i % 2 == 1 {
                santa.x += 1;
            } else {
                robo_santa.x += 1;
            }
        }

        if i % 2 == 1 {
            houses.insert(robo_santa.clone());
        } else {
            houses.insert(santa.clone());
        }
    }

    houses.len()
}

fn main() {
    let input = include_str!("input.txt");

    let part1_result = part1(input);
    println!("part1: {}", part1_result);

    let part2_result = part2(input);
    println!("part2: {}", part2_result);
}
