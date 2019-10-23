use std::collections::BinaryHeap;

pub fn last_stone_weight(stones: Vec<i32>) -> i32 {
    let mut heap = BinaryHeap::from(stones);
    while heap.len() > 1 {
        let lhs = heap.pop().unwrap();
        let rhs = heap.pop().unwrap();
        let result = (lhs - rhs).abs();
        if result != 0 {
            heap.push(result);
        }
    }
    heap.pop().unwrap_or(0)
}
