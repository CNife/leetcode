pub fn combination_sum(mut candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
    candidates.sort_unstable();
    let mut results = vec![];
    let mut stack = vec![];
    backtrace(&candidates, target, &mut stack, &mut results);
    results
}

fn backtrace(candidates: &[i32], target: i32, stack: &mut Vec<i32>, results: &mut Vec<Vec<i32>>) {
    for (index, &candidate) in candidates.into_iter().enumerate() {
        if candidate <= target {
            stack.push(candidate);
            if candidate < target {
                backtrace(&candidates[index..], target - candidate, stack, results);
            } else {
                results.push(stack.clone());
            }
            stack.pop().unwrap();
        } else {
            break;
        }
    }
}
