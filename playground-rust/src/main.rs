fn main() {
    fn fact(n: i32) -> i32 {
        if n == 1 {
            return 1;
        }
        n * fact(n - 1)
    }

    println!("{}", fact(5));
}
