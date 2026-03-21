

fn takes_ownership(some_string: String) {
    println!("获得所有权: {}", some_string);
}

fn makes_copy(some_integer: i32) {
    println!("获得拷贝: {}", some_integer);
}


use std::thread;
use std::sync::{Arc, Mutex};

/// 并发计数器示例：启动多个线程并安全地累加计数
fn concurrent_counter() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        let handle = thread::spawn(move || {
            for _ in 0..1000 {
                let mut num = counter.lock().unwrap();
                *num += 1;
            }
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("并发计数结果: {}", *counter.lock().unwrap());
}


fn main() {
    // 斐波那契数列的前 10 项
    let mut fib = vec![0, 1];
    for i in 2..10 {
        let next = fib[i - 1] + fib[i - 2];
        fib.push(next);
    }
    println!("斐波那契数列前 10 项: {:?}", fib);

    // 使用 match 进行模式匹配
    let number = 7;
    match number {
        1 => println!("数字是一"),
        2 | 3 | 5 | 7 | 11 => println!("数字是质数"),
        13..=19 => println!("数字在 13 到 19 之间"),
        _ => println!("其他数字"),
    }

    // 所有权与借用示例
    let s = String::from("hello");
    takes_ownership(s);
    // println!("{}", s); // 这一行会报错，因为 s 的所有权已被转移

    let x = 5;
    makes_copy(x);
    println!("x 仍然可以使用: {}", x);

    // 闭包示例
    let add = |a: i32, b: i32| a + b;
    println!("3 + 4 = {}", add(3, 4));
}