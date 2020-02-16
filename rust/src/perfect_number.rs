pub fn check_perfect_number(n: i32) -> bool {
    debug_assert!(n > 0);

    if n == 1 {
        false
    } else {
        let mut sum = 1;
        let bound = (n as f64).sqrt().floor() as i32;
        for i in 2..=bound {
            if n % i == 0 {
                sum += i;
                sum += (n / i);
            }
        }
        sum == n
    }
}

#[test]
fn test_1e8() {
    for i in 1..100_000_000 {
        if i % 1_000_000 == 0 {
            println!("tested {}%", i / 1000000);
        }
        if check_perfect_number(i) {
            println!("{}", i);
        }
    }
}

#[test]
fn test() {
    let cases = vec![(28, true)];
    for (n, expected) in cases {
        assert_eq!(check_perfect_number(n), expected);
    }
}
