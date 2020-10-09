pub fn divisor_game(n: i32) -> bool {
    let n = n as usize;
    let mut win = vec![false; if n == 1 { 3 } else { n + 1 }];
    win[2] = true;

    for i in 3..=n {
        for j in 1..=i / 2 {
            if i % j == 0 && !win[i - j] {
                win[i] = true;
            }
        }
    }
    win[n]
}

#[cfg(test)]
mod test {
    use rand::Rng;

    use super::*;

    fn divisor_game_fast(n: i32) -> bool {
        n & 1 == 0
    }

    #[test]
    fn test_divisor_game() {
        let test_times = 100;

        let mut rng = rand::thread_rng();
        let gen_test_case = || {
            let n = rng.gen_range(1, 1001);
            (n, divisor_game_fast(n))
        };

        for (n, expected) in std::iter::repeat_with(gen_test_case).take(test_times) {
            assert_eq!(divisor_game(n), expected);
        }
    }
}
