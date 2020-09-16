#[macro_use]
extern crate lazy_static;
#[macro_use]
extern crate maplit;

mod cell_test;
mod from_into_test;
mod lazy_static_test;
mod lifetime_test;
mod serde_json_test;

fn main() {
    println! {"hello, world!"};
}
