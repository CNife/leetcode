use std::collections::HashSet;
use std::hash::{Hash, Hasher};
use std::iter::FromIterator;

#[derive(Eq, PartialEq, Debug)]
pub struct Collection<T>(HashSet<T>)
where
    T: Eq + Hash;

impl Hash for Collection<i32> {
    fn hash<H: Hasher>(&self, state: &mut H) {
        let mut v: Vec<i32> = self.0.iter().copied().collect();
        v.sort_unstable();
        for n in v {
            n.hash(state);
        }
    }
}

impl From<Vec<i32>> for Collection<i32> {
    fn from(v: Vec<i32>) -> Self {
        Collection(HashSet::from_iter(v.into_iter()))
    }
}

impl From<Vec<Vec<i32>>> for Collection<Collection<i32>> {
    fn from(v: Vec<Vec<i32>>) -> Self {
        Collection(HashSet::from_iter(v.into_iter().map(|inner| inner.into())))
    }
}
