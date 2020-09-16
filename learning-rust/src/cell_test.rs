// cell_test.rs
//
// Cell 实现了T类型的内部可变性
// Rust中Cell是一个容器，在整个结构体不能改变的情况下
// 如果结构体的某些成员使用cell包装起来，就使得这个被cell包装的成员变量可以改变
// 例如，实现Foo的get_name每调用一次，count就增加1
use std::cell::Cell;
struct Foo {
    name: &'static str,
    count: Cell<u32>, // 此处使用Cell
}

impl Foo {
    fn get_name(&self) -> &str {
        let new_count = self.count.get() + 1;
        // self.count = new_count; // 如果count是u32，那么此处的count是不能改变的，编译器会报错
        self.count.set(new_count);
        self.name
    }
}

#[test]
fn test_cell() {
    let foo = Foo {
        name: "foo",
        count: Cell::new(0),
    };
    println! {"name={:?}, count={:?}", foo.get_name(), foo.count};
    assert_eq!(foo.count.get(), 1);
    println! {"name={:?}, count={:?}", foo.get_name(), foo.count};
    assert_eq!(foo.count.get(), 2);
    println! {"name={:?}, count={:?}", foo.get_name(), foo.count};
    assert_eq!(foo.count.get(), 3);
    println! {"name={:?}, count={:?}", foo.get_name(), foo.count};
    assert_eq!(foo.count.get(), 4);
    println! {"name={:?}, count={:?}", foo.get_name(), foo.count};
    assert_eq!(foo.count.get(), 5);
}

// ------------------------
// RefCell实现了指针的内部可变性
// RefCell与Cell的区别是，Cell返回T的值的拷贝，而RefCell只能borrow出指针
use std::cell::RefCell;

struct Bar {
    name: &'static str,
    count: RefCell<u32>,
}

impl Bar {
    fn get_name(&self) -> &str {
        let mut ref_count = self.count.borrow_mut();
        *ref_count += 1;
        self.name
    }
}

#[test]
fn test_refcell() {
    let bar = Bar {
        name: "foo",
        count: RefCell::new(0),
    };
    println! {"name={:?}, count={:?}", bar.get_name(), bar.count.borrow()};
    assert_eq!(*(bar.count.borrow()), 1);
    println! {"name={:?}, count={:?}", bar.get_name(), bar.count.borrow()};
    assert_eq!(*(bar.count.borrow()), 2);
    println! {"name={:?}, count={:?}", bar.get_name(), bar.count.borrow()};
    assert_eq!(*(bar.count.borrow()), 3);
    println! {"name={:?}, count={:?}", bar.get_name(), bar.count.borrow()};
    assert_eq!(*(bar.count.borrow()), 4);
    println! {"name={:?}, count={:?}", bar.get_name(), bar.count.borrow()};
    assert_eq!(*(bar.count.borrow()), 5);
}
