extern crate grpcio;
extern crate protos;

use std::env;
use std::sync::Arc;

use grpcio::{ChannelBuilder, EnvBuilder};

use protos::hello::HelloRequest;
use protos::hello_grpc::HelloServiceClient;

fn main() {
    let port = 50_055;
    let env = Arc::new(EnvBuilder::new().build());
    let ch = ChannelBuilder::new(env).connect(format!("localhost:{}", port).as_str());
    let client = HelloServiceClient::new(ch);

    let mut req = HelloRequest::new();
    req.greeting = "How are you?".to_string();

    let resp = client.say_hello(&req);
    println!("{:?}", resp);
}
