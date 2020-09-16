use std::convert::From;

#[derive(Debug)]
struct Square {
    width: i32,
    height: i32,
}

impl From<i32> for Square {
    fn from(size: i32) -> Self {
        Square {
            width: size,
            height: size,
        }
    }
}

impl Square {
    fn area(&self) -> i32 {
        self.width * self.height
    }
}

#[test]
fn test_from_into() {
    let s = Square::from(10);
    println!("s is {:?},area is {}", s, s.area());
    let a: Square = 20.into(); //这里的Square类型需要显示标注
    println!("a is {:?}", a);
}
