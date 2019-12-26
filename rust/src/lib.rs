pub mod longest_increasing_path_in_a_matrix;

pub mod utils {
    use std::collections::HashSet;
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
}
