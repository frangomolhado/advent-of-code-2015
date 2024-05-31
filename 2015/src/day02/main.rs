fn part1(input: &str) -> u32 {
    let mut total = 0;

    for line in input.lines() {
        let (l, w, h) = get_dimensions(line);

        let side1 = l * w;
        let side2 = w * h;
        let side3 = h * l;

        total += 2 * (side1 + side2 + side3);
        total += side1.min(side2).min(side3);
    }

    total
}

fn part2(input: &str) -> u32 {
    let mut total = 0;

    for line in input.lines() {
        let (l, w, h) = get_dimensions(line);

        if l <= w && l <= h {
            total += 2 * (l + w.min(h));
        } else if l <= w && l > h {
            total += 2 * (l + h);
        } else if l <= h && l > w {
            total += 2 * (l + w);
        } else {
            total += 2 * (w + h);
        }

        total += l * w * h;
    }

    total
}

fn get_dimensions(line: &str) -> (u32, u32, u32) {
    let dimensions = line.split('x').collect::<Vec<_>>();
    let l: u32 = dimensions[0].parse().unwrap();
    let w: u32 = dimensions[1].parse().unwrap();
    let h: u32 = dimensions[2].parse().unwrap();

    (l, w, h)
}

fn main() {
    let input = include_str!("input.txt");

    let part1_result = part1(input);
    println!("part1: {}", part1_result);

    let part2_result = part2(input);
    println!("part2: {}", part2_result);
}
