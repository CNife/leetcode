use std::fmt::Debug;

pub fn v<F, T>(src: Vec<F>) -> Vec<T>
where
    F: Into<T>,
{
    src.into_iter().map(|item| item.into()).collect()
}

pub fn vv<F, T>(src: Vec<Vec<F>>) -> Vec<Vec<T>>
where
    F: Into<T>,
{
    src.into_iter().map(v).collect()
}

pub fn assert_same_set<Lhs, Rhs>(mut output: Vec<Lhs>, want: Vec<Rhs>)
where
    Lhs: From<Rhs> + Ord + Debug,
{
    let mut want: Vec<Lhs> = want.into_iter().map(|item| item.into()).collect();
    output.sort_unstable();
    want.sort_unstable();
    assert_eq!(output, want);
}
