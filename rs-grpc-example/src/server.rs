extern crate futures;
extern crate grpcio;
extern crate protos;

use std::io::Read;
use std::sync::Arc;
use std::{io, thread};

use futures::sync::oneshot;
use futures::Future;
use grpcio::{Environment, RpcContext, ServerBuilder, UnarySink};

use protos::hello::{HelloRequest, HelloResponse};
use protos::hello_grpc::{create_hello_service, HelloService};

#[derive(Clone)]
struct Hello;

impl HelloService for Hello {
    fn say_hello(
        &mut self,
        ctx: RpcContext<'_>,
        req: HelloRequest,
        sink: UnarySink<HelloResponse>,
    ) {
        dbg!(&req.greeting);
        let mut resp = HelloResponse::new();
        resp.reply = "I'm fine, thank you.".to_string();
        sink.success(resp.clone());
    }
}

fn main() {
    let env = Arc::new(Environment::new(1));
    let service = create_hello_service(Hello);
    let mut server = ServerBuilder::new(env)
        .register_service(service)
        .bind("127.0.0.1", 50055)
        .build()
        .unwrap();
    server.start();
    for &(ref host, port) in server.bind_addrs() {
        println!("listening on {}:{}", host, port);
    }
    let (tx, rx) = oneshot::channel();
    thread::spawn(move || {
        println!("Press ENTER to exit...");
        let _ = io::stdin().read(&mut [0]).unwrap();
        tx.send(())
    });
    let _ = rx.wait();
    let _ = server.shutdown().wait();
}
