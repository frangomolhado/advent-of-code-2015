fn part1(input: &str) -> u32 {
    mine(input, 5)
}

fn part2(input: &str) -> u32 {
    mine(input, 6)
}

fn mine(pass_key: &str, nzeros: usize) -> u32 {
    let mut sufix = 1;
    loop {
        let str_sufix = sufix.to_string();
        let hash = format!(
            "{:x}",
            md5::compute(
                (pass_key[..pass_key.len() - 1].to_string() + str_sufix.as_str()).as_bytes()
            )
        );

        let start = u32::from_str_radix(&hash[0..nzeros], 16).unwrap();
        if start == 0 {
            return sufix;
        }

        sufix += 1;
    }
}

fn main() {
    let input = include_str!("input.txt");

    let part1_result = part1(input);
    println!("part 1: {}", part1_result);

    let part2_result = part2(input);
    println!("part 2: {}", part2_result);
}
