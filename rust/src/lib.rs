pub mod find_numbers_with_even_number_of_digits;

#[cfg(test)]
pub mod utils {
    use std::collections::HashSet;
    use std::fmt::Debug;
    use std::hash::Hash;
    use std::iter::FromIterator;

    pub fn sv(sv: Vec<&str>) -> Vec<String> {
        sv.into_iter().map(|s| s.to_string()).collect()
    }

    pub fn svv(svv: Vec<Vec<&str>>) -> Vec<Vec<String>> {
        svv.into_iter().map(sv).collect()
    }

    pub fn s<T>(v: Vec<T>) -> HashSet<T>
    where
        T: Eq + Hash,
    {
        HashSet::from_iter(v.into_iter())
    }

    pub fn assert_same_set<T>(mut output: Vec<T>, mut expected: Vec<T>)
    where
        T: Ord + Debug,
    {
        output.sort_unstable();
        expected.sort_unstable();
        assert_eq!(output, expected);
    }
}
