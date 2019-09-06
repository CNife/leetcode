macro_rules! print_size_of_type {
    ($type: ty) => {
        println!(
            "sizeof({}) = {}",
            stringify!($type),
            std::mem::size_of::<$type>()
        );
    };
}

fn main() {
    print_size_of_type!(&mut Vec<i32>);
    print_size_of_type!(Option<&mut Vec<i32>>);
}
