pub fn my_atoi(s: String) -> i32 {
    let s = s.into_bytes();
    let mut itr = s.into_iter();
    let mut is_positive = true;
    let mut result = 0i64;

    while let Some(ch) = itr.next() {
        match ch {
            b' ' => {}
            b'-' => {
                is_positive = false;
                break;
            }
            b'+' => break,
            c if c.is_ascii_digit() => {
                result = (c - b'0') as i64;
                break;
            }
            _ => return 0,
        }
    }

    while let Some(ch) = itr.next() {
        match ch {
            c if c.is_ascii_digit() => {
                result = result * 10 + (c - b'0') as i64;
                if result > std::i32::MAX as i64 && is_positive {
                    return std::i32::MAX;
                } else if result > (std::i32::MAX as i64 + 1) && !is_positive {
                    return std::i32::MIN;
                }
            }
            _ => break,
        }
    }

    (if is_positive { result } else { -result }) as i32
}

#[test]
fn test() {
    let cases = vec![
        ("42", 42),
        ("   -42", -42),
        ("4193 with words", 4193),
        ("words and 987", 0),
        ("-91283472332", -2147483648),
        ("-+1", 0),
    ];
    for (str, expected) in cases {
        assert_eq!(my_atoi(str.to_string()), expected);
    }
}
