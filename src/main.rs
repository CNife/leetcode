fn main() {
    let nums = [
        "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];

    print!("      ");
    for ch in 0u8..26 {
        let ch = ch + 'a' as u8;
        print!("{:^3}", ch as char);
    }
    println!();

    for &num in &nums {
        let mut counts = [0; 26];
        for ch in num.as_bytes() {
            let index = ch - 'a' as u8;
            counts[index as usize] += 1;
        }
        // println!("{:?}", counts);

        print!("{:6}", num);
        for &count in &counts {
            if count == 0 {
                print!("   ");
            } else {
                print!("{:^3}", count);
            }
        }
        println!();
    }
}
