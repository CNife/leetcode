pub fn all_cells_dist_order(r: i32, c: i32, r0: i32, c0: i32) -> Vec<Vec<i32>> {
    let mut vp = Vec::new();
    for i in 0..r {
        for j in 0..c {
            let dist = ((i - r0).abs() + (j - c0).abs()) as usize;
            if dist >= vp.len() {
                for _ in vp.len()..=dist {
                    vp.push(Vec::new());
                }
            }
            vp[dist].push(vec![i, j]);
        }
    }
    vp.into_iter()
        .flat_map(|points| points.into_iter())
        .collect()
}
