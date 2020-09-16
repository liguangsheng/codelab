use serde::{Deserialize, Serialize};

const JSON_STRING: &'static str = r#"
{
    "name": "John Doe",
    "age": 43,
    "address": {
        "street": "10 Downing Street",
        "city": "London"
    },
    "phones": [
        "+44 1234567",
        "+44 2345678"
    ]
}
"#;

#[test]
fn test_serde_json_decode_to_value() {
    let v: serde_json::Value = serde_json::from_str(JSON_STRING).unwrap();
    println!("name={}", v["name"]);
    println!("age={}", v["age"]);
    assert_eq!(v["name"], "John Doe");
    assert_eq!(v["age"], 43);
}

#[derive(Debug, Serialize, Deserialize)]
struct Address {
    street: String,
    city: String,
}

#[derive(Debug, Serialize, Deserialize)]
struct Person {
    name: String,
    age: u8,
    address: Address,
    phones: Vec<String>,
}

#[test]
fn test_serde_json_decode_to_struct() {
    let v: Person = serde_json::from_str(JSON_STRING).unwrap();
    println!("{:?}", v);
    assert_eq!(v.name, "John Doe");
    assert_eq!(v.address.city, "London");
    assert_eq!(v.phones[0], "+44 1234567");
}
