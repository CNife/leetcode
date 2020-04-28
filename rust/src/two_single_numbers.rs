pub fn single_numbers(nums: Vec<i32>) -> Vec<i32> {
    let k = nums.iter().fold(0, |acc, num| acc ^ num);
    let mask = (0..32)
        .map(|shift| 1 << shift)
        .find(|&mask| mask & k != 0)
        .unwrap();
    let (a, b) = nums.into_iter().fold((0, 0), |(a, b), num| {
        if mask & num == 0 {
            (a ^ num, b)
        } else {
            (a, b ^ num)
        }
    });
    vec![a, b]
}

#[test]
fn test() {
    let cases = vec![
        (vec![4, 1, 4, 6], vec![6, 1]),
        (vec![1, 2, 10, 4, 1, 4, 3, 3], vec![2, 10]),
    ];
    for (nums, expect) in cases {
        use crate::utils::s;
        assert_eq!(s(single_numbers(nums)), s(expect));
    }
}
