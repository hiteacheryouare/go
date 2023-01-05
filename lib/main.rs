use std::env;
use std::error::Error;
use std::str::FromStr;

use hyper::{Body, Method, Response, Server, StatusCode};
use hyper::header::{HeaderValue, SERVER};
use hyper::service::{make_service_fn, service_fn};
use hyper::{header, Body};
use serde::{Deserialize, Serialize};
use serde_json::{json, Value};

#[derive(Serialize, Deserialize)]
struct ResponseData {
    data: String,
    code: i32,
    message: String,
}

fn reverse_string(s: String) -> String {
    s.chars().rev().collect()
}

async fn simple_json_response(e: &str, d: String, c: i32) {
    let service = make_service_fn(|_| async {
        Ok::<_, hyper::Error>(service_fn(|req| {
            let mut response = Response::new(Body::empty());
            match (req.method(), req.uri().path()) {
                (&Method::GET, e) => {
                    let code = c;
                    response.headers_mut().insert(
                        SERVER,
                        HeaderValue::from_str(&format!("{}", env!("CARGO_PKG_VERSION"))).unwrap(),
                    );
                    *response.status_mut() = StatusCode::from_u16(code as u16).unwrap();
                    let data = ResponseData {
                        data: d,
                        code: c,
                    };
                    let json_data = serde_json::to_string(&data).unwrap();
                    *response.body_mut() = json_data.into();
                }
                _ => {
                    *response.status_mut() = StatusCode::NOT_FOUND;
                }
            };
            async { response }
        }))
    });
    let addr = ([127, 0, 0, 1], 3000).into();
    let server = Server::bind(&addr).serve(service);
    println!("Listening on http://{}", addr);
    server.await;
}

async fn simple_query_params_json_response(e: &str, d: String, p: &str, c: i32) {
    let service = make_service_fn(|_| async {
        Ok::<_, hyper::Error>(service_fn(|req| {
            let mut response = Response::new(Body::empty());
            match (req.method(), req.uri().path()) {
                (&Method::GET, e) => {
                    let input = req
                        .uri()
                        .query()
                        .unwrap_or("")
                        .split('=')
                        .nth(1)
                        .unwrap_or("");
                    let reversed_input = reverse_string(input.to_string());
                    let code = c;
                    response.headers_mut().insert(
                        SERVER,
                        HeaderValue::from_str(&format!("{}", env!("CARGO_PKG_VERSION"))).unwrap(),
                    );
                    *response.status_mut()
= StatusCode::from_u16(code as u16).unwrap();
                    let data = ResponseData {
                        data: reversed_input,
                        message: d,
                        code: c,
                    };
                    let json_data = serde_json::to_string(&data).unwrap();
                    *response.body_mut() = json_data.into();
                }
                _ => {
                    *response.status_mut() = StatusCode::NOT_FOUND;
                }
            };
            async { response }
        }))
    });
    let addr = ([127, 0, 0, 1], 3000).into();
    let server = Server::bind(&addr).serve(service);
    println!("Listening on http://{}", addr);
    server.await;
}

async fn simple_html_response(e: &str, h: String, c: i32) {
    let service = make_service_fn(|_| async {
        Ok::<_, hyper::Error>(service_fn(|req| {
            let mut response = Response::new(Body::empty());
            match (req.method(), req.uri().path()) {
                (&Method::GET, e) => {
                    let code = c;
                    response.headers_mut().insert(
                        SERVER,
                        HeaderValue::from_str(&format!("{}", env!("CARGO_PKG_VERSION"))).unwrap(),
                    );
                    response.headers_mut().insert(
                        header::CONTENT_TYPE,
                        HeaderValue::from_str("text/html").unwrap(),
                    );
                    *response.status_mut() = StatusCode::from_u16(code as u16).unwrap();
                    *response.body_mut() = h.into();
                }
                _ => {
                    *response.status_mut() = StatusCode::NOT_FOUND;
                }
            };
            async { response }
        }))
    });
    let addr = ([127, 0, 0, 1], 3000).into();
    let server = Server::bind(&addr).serve(service);
    println!("Listening on http://{}", addr);
    server.await;
}

async fn redirect_not_found(e: &str) {
    let service = make_service_fn(|_| async {
        Ok::<_, hyper::Error>(service_fn(|req| {
            let mut response = Response::new(Body::empty());
            match (req.method(), req.uri().path()) {
                (&Method::GET, e) => {
                    *response.status_mut() = StatusCode::NOT_FOUND;
                    let location = "https://porkyproductions.github.io/404.html";
                    response.headers_mut().insert(
                        header::LOCATION,
                        HeaderValue::from_str(location).unwrap(),
                    );
                }
                _ => {
                    *response.status_mut() = StatusCode::NOT_FOUND;
                }
            };
            async { response }
        }))
    });
    let addr = ([127, 0, 0, 1], 3000).
into();
    let server = Server::bind(&addr).serve(service);
    println!("Listening on http://{}", addr);
    server.await;
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    let port = env::var("PORT").unwrap_or("3000".to_string());
    let port: u16 = port.parse()?;
    let addr = ([127, 0, 0, 1], port).into();
    simple_json_response("/", "It works!".to_string(), 200).await;
    simple_query_params_json_response("/reverse", "Use data".to_string(), "input", 203).await;
    simple_html_response("/dmv1", dmv1_html.to_string(), 200).await;
    redirect_not_found("/goober").await;
    println!("Listening on http://{}", addr);
    Ok(())
}
