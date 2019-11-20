use std::collections::HashSet;

pub fn combination_sum2(mut candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
    candidates.sort_unstable();
    let mut results = HashSet::new();
    let mut stack = vec![];
    backtrace(&candidates, target, &mut stack, &mut results);
    results.into_iter().collect()
}

fn backtrace(
    candidates: &[i32],
    target: i32,
    stack: &mut Vec<i32>,
    results: &mut HashSet<Vec<i32>>,
) {
    for (index, &candidate) in candidates.into_iter().enumerate() {
        if candidate <= target {
            stack.push(candidate);
            if candidate < target {
                backtrace(&candidates[index + 1..], target - candidate, stack, results);
            } else {
                results.insert(stack.clone());
            }
            stack.pop().unwrap();
        } else {
            break;
        }
    }
}
