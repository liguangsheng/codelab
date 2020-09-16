use std::collections::HashMap;

fn main() {
    // let mut v: Vec<HashMap<String, String>> =
    //     (0..10).map(|x| HashMap::<String, String>::new()).collect();

    let mut v: Vec<HashMap<String, String>> = vec![HashMap::new(); 10];
    v[1].insert("hello1".to_owned(), "world1".to_owned());
    v[0].insert("hello0".to_owned(), "world0".to_owned());
    v[2].insert("hello2".to_owned(), "world2".to_owned());
    dbg! {&v};
}
