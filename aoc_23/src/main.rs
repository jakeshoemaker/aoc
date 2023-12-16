use std::{
    fs::File,
    io::{prelude::*, BufReader},
    path::Path,
};

fn lines_from_file(filename: impl AsRef<Path>) -> Vec<Vec<char>> {
    let file = File::open(filename).expect("no such file");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("Could not parse line").chars().collect())
        .collect()
}

fn is_num(ch: char) -> Option<char> {
    match ch {
        '0'..='9' => Some(ch),
        _ => None,
    }
}

fn main() {
    let mut total = 0;
    let calibrations = lines_from_file("inputs/day01.txt");
    for vec_c in calibrations {
        let mut temp: Vec<char> = Vec::new();
        vec_c.iter().for_each(|c| match is_num(*c) {
            Some(n) => temp.push(n),
            None => (),
        });

        let mut num_as_string = String::new();
        num_as_string.push(temp[0]);
        num_as_string.push(temp.pop().expect("expected a number here!"));
        total += num_as_string
            .parse::<i32>()
            .expect("should be a string number");
    }
    println!("Day 1 part 1: {:?}", total);
}
