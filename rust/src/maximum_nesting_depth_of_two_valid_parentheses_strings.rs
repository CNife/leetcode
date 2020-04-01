pub fn max_depth_after_split(input: String) -> Vec<i32> {
    let mut result = Vec::with_capacity(input.len());
    for (i, byte) in input.into_bytes().into_iter().enumerate() {
        let result_bit = if byte == b'(' { i & 1 } else { (i + 1) & 1 };
        result.push(result_bit as i32);
    }
    result
}
