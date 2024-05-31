#[derive(Debug)]
enum Action {
    TurnOff,
    TurnOn,
    Toggle,
}

#[derive(Debug)]
struct Point {
    x: u32,
    y: u32,
}

impl Point {
    fn new(x: u32, y: u32) -> Self {
        Self { x, y }
    }
}

#[derive(Debug)]
struct Instruction {
    action: Action,
    start: Point,
    end: Point,
}

impl Instruction {
    fn new(action: Action, start: Point, end: Point) -> Self {
        Self { action, start, end }
    }
}

fn parse_input(input: &str) -> Vec<Instruction> {
    let mut result = vec![];
    for line in input.lines() {
        let splitted: Vec<_> = line.split(" ").collect();
        if splitted[0] == "turn" {
            let action = if splitted[1] == "on" {
                Action::TurnOn
            } else {
                Action::TurnOff
            };
            let p1: Vec<_> = splitted[2].split(",").collect();
            let p2: Vec<_> = splitted[4].split(",").collect();
            result.push(Instruction::new(
                action,
                Point::new(p1[0].parse().unwrap(), p1[1].parse().unwrap()),
                Point::new(p2[0].parse().unwrap(), p2[1].parse().unwrap()),
            ));
        } else {
            let action = Action::Toggle;
            let p1: Vec<_> = splitted[1].split(",").collect();
            let p2: Vec<_> = splitted[3].split(",").collect();
            result.push(Instruction::new(
                action,
                Point::new(p1[0].parse().unwrap(), p1[1].parse().unwrap()),
                Point::new(p2[0].parse().unwrap(), p2[1].parse().unwrap()),
            ));
        }
    }
    result
}

fn part1(input: &Vec<Instruction>) -> u32 {
    let mut grid = [0u8; 1000 * 1000];

    for inst in input.into_iter() {
        for x in inst.start.x..=inst.end.x {
            for y in inst.start.y..=inst.end.y {
                let i = (x * 1000 + y) as usize;
                match inst.action {
                    Action::TurnOff => {
                        grid[i] = 0;
                    }
                    Action::TurnOn => {
                        grid[i] = 1;
                    }
                    Action::Toggle => {
                        grid[i] ^= 1;
                    }
                }
            }
        }
    }

    // can't use `iter::sum()` because of grid beeing u8 and return u32
    return grid.into_iter().fold(0, |acc, n| acc + n as u32);
}

fn part2(input: &Vec<Instruction>) -> u32 {
    let mut grid = [0u8; 1000 * 1000];

    for inst in input.into_iter() {
        for x in inst.start.x..=inst.end.x {
            for y in inst.start.y..=inst.end.y {
                let i = (x * 1000 + y) as usize;
                match inst.action {
                    Action::TurnOff => {
                        grid[i] = if grid[i] > 0 { grid[i] - 1 } else { 0 };
                    }
                    Action::TurnOn => {
                        grid[i] += 1;
                    }
                    Action::Toggle => {
                        grid[i] += 2;
                    }
                }
            }
        }
    }

    // can't use `iter::sum()` because of grid beeing u8 and return u32
    return grid.into_iter().fold(0, |acc, n| acc + n as u32);
}

fn main() {
    let input = include_str!("input.txt");
    let parsed_input = parse_input(input);

    let part1_result = part1(&parsed_input);
    println!("part 1: {}", part1_result);

    let part2_result = part2(&parsed_input);
    println!("part 2: {}", part2_result);
}
