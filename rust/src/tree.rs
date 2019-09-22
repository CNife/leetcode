use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug, Clone, Eq, PartialEq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

pub fn tree_from_vec(values: Vec<Option<i32>>) -> Option<Rc<RefCell<TreeNode>>> {
    let values = values.into_iter()
        .map(|val| val.map(|n| Rc::new(RefCell::new(TreeNode::new(n)))))
        .collect::<Vec<_>>();
    for i in 0..values.len() / 2 {
        let left_index = i * 2 + 1;
        let right_index = left_index + 1;
        if let Some(ref root) = values[i] {
            root.borrow_mut().left = values[left_index].clone();
            root.borrow_mut().right = values[right_index].clone();
        }
    }
    match values.first() {
        Some(v) => v.clone(),
        None => None,
    }
}

#[test]
fn test_tree_from_vec() {
    let values: Vec<_> = vec![1, 2, 3, 4, 5, 6, 7].into_iter()
        .map(|n| Some(n))
        .collect();
    let root = tree_from_vec(values);
    println!("{:?}", root);
    assert_eq!(None, root);
}
