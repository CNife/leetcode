pub fn group_anagrams(strings: Vec<String>) -> Vec<Vec<String>> {
    const PRIMES: [u32; 26] = [
        2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97,
        101, 103,
    ];

    let mut indexes: Vec<u32> = Vec::new();
    let mut buckets: Vec<Vec<String>> = Vec::new();
    for string in strings {
        let key = string
            .bytes()
            .fold(1, |product, ch| product * PRIMES[(ch - b'a') as usize]);
        match indexes.binary_search(&key) {
            Ok(index) => {
                buckets[index].push(string);
            }
            Err(insert_index) => {
                indexes.insert(insert_index, key);
                buckets.insert(insert_index, vec![string]);
            }
        }
    }
    buckets
}

//use std::collections::HashMap;
//pub fn group_anagrams(strings: Vec<String>) -> Vec<Vec<String>> {
//    let mut map = HashMap::new();
//    for string in strings {
//        let mut char_counts = [0u8; 26];
//        for ch in string.bytes() {
//            let index = ch - b'a';
//            char_counts[index as usize] += 1;
//        }
//        map.entry(char_counts).or_insert_with(Vec::new).push(string);
//    }
//    map.into_iter().map(|(_, v)| v).collect()
//}
