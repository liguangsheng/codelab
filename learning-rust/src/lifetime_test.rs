// read https://www.jianshu.com/p/c186f67a0cf2

// 此max函数将返回m和n两个数中较大的一个
// 由于此函数没有用到引用，所以不存在内存安全问题
fn max(m: i32, n: i32) -> i32 {
    if m > n {
        return m;
    }
    n
}

#[test]
fn test_without_reference() {
    let i = 1;
    let j = 2;
    let k = max(i, j);
    println!("{}", k)
}

// 此max2函数，在参数中用到引用，但在返回值中没有引用
// 所以直接可以推导出在函数声明收起内，是内存安全的
fn max2(m: &i32, n: &i32) -> i32 {
    if m > n {
        return *m;
    }
    *n
}

#[test]
fn test_without_lifetime() {
    let i = 3;
    let j = 4;
    let k = max2(&i, &j);
    println!("{}", k)
}

// 此max3函数由于返回了引用类型，所以不能保证所引用的内容在函数的声明周期结束后仍然是安全的
// 因此此函数不能通过编译，要想通过编译，就需要为此函数加上声明周期的声明
// fn max3(m: &i32, n: &i32) -> &i32 {
//     if m > n {
//         return *m;
//     }
//     n
// }

// 此max4函数加上了单个lifetime的声明
// 这个声明的含义是，声明了名为a的lifetime
// max4返回的引用的lifetime是a，那么m和n的lifetime至少也要是a，才能满足要求
fn max4<'a>(m: &'a i32, n: &'a i32) -> &'a i32 {
    if m > n {
        return m;
    }
    n
}

// #[test]
// fn test_single_lifetime() {
//     let i = 1;
//     let k = 100;
//     {
//         let j = 2;
//         k = max4(&i, &j); // 此处j的lifetime小于k的lifetime，所以不能通过编译
//     }
//     println!("{}", k)
// }

#[test]
fn test_single_lifetime() {
    let i = 5;
    let j = 6;
    let k = max4(&i, &j); // 此处i,j,k的lifetime都相同，所以正确
    println!("{}", k)
}

#[test]
fn test_single_lifetime2() {
    let i = 7;
    let j = 8;
    {
        let k = max4(&i, &j); // 此处i,j的lifetime均大于k的lifetime，所以正确
        println!("{}", k)
    }
}

// 此max5函数，引用了o，但返回值只可能是m，或者n，所以o的lifetime实际上可以不设计为a
// o的lifetime可以不写，也可以显示的再声明一个b
fn max5<'a>(m: &'a mut i32, n: &'a mut i32, o: &i32) -> &'a i32 {
    if m > n {
        *m += o;
        return m;
    }
    *n += o;
    n
}

#[test]
fn test_multi_lifetime() {
    let mut i = 9;
    let mut j = 10;
    let k: &i32;
    {
        let o = 3;
        k = max5(&mut i, &mut j, &o);
    }
    println!("{}", k);
}
