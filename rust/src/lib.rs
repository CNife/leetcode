pub mod permutations_2;

#[cfg(test)]
pub mod utils {
    use std::collections::{BTreeSet, HashSet};
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

    pub fn ss<T>(v: Vec<Vec<T>>) -> BTreeSet<BTreeSet<T>>
    where
        T: Ord,
    {
        v.into_iter()
            .map(|inner| inner.into_iter().collect::<BTreeSet<_>>())
            .collect()
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
