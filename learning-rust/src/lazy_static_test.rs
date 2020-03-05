//-----------------------------
// lazy_static with std HashMap
use std::collections::HashMap;

lazy_static! {
    static ref HASHMAP: HashMap<u32, &'static str> = {
        let mut m = HashMap::new();
        m.insert(0, "foo");
        m.insert(1, "bar");
        m.insert(2, "baz");
        m
    };
    static ref COUNT: usize = HASHMAP.len();
    static ref NUMBER: u32 = times_two(21);
}

fn times_two(n: u32) -> u32 {
    n * 2
}

#[test]
fn test_lazy_static_std_hashmap() {
    assert_eq!(HASHMAP.get(&0).unwrap(), &"foo");
    assert_eq!(HASHMAP.get(&1).unwrap(), &"bar");
    assert_eq!(HASHMAP[&2], "baz");
    assert_eq!(*COUNT, 3);
    assert_eq!(*NUMBER, 42);
}

//--------------------------------------
// lazy_static with maplit hashmap macro

lazy_static! {
    static ref MAPLIT_HASHMAP: HashMap<&'static str, &'static str> = hashmap! {
    "foo" => "bar",
    "bar" => "foo",
    };
}

#[test]
fn test_lazy_static_maplit_hashmao() {
    assert_eq!(MAPLIT_HASHMAP[&"foo"], "bar");
    assert_eq!(MAPLIT_HASHMAP.get("foo").unwrap(), &"bar");
}
