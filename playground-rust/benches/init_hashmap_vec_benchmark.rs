use criterion::{criterion_group, criterion_main, Criterion};

use std::collections::HashMap;

fn init_by_vec() {
    let mut _v: Vec<HashMap<String, String>> = vec![HashMap::new(); 10];
}

fn init_by_map() {
    let mut _v: Vec<HashMap<String, String>> =
        (0..10).map(|x| HashMap::<String, String>::new()).collect();
}

fn criterion_benchmark(c: &mut Criterion) {
    c.bench_function("init_by_vec", |b| b.iter(|| init_by_vec()));
    c.bench_function("init_by_map", |b| b.iter(|| init_by_map()));
}

criterion_group!(benches, criterion_benchmark);
criterion_main!(benches);
