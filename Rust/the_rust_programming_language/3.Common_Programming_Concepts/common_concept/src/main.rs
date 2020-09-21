use std::string::{String};

const PI: f32 = 3.141592;

fn variables(name: String) -> bool {

    println!("function: {}", name);

    // immutable
    let x = 5;
    println!("The value of x is: {}", x);
    // x = 6;  // cannot assign twice to immutable variable `x`
    // println!("The value of x is: {}", x);

    // mutable
    let mut a = 10;
    println!("The value of a is: {}", a);
    a = 11;
    println!("The value of a is: {}", a);

    // constants
    const MAX_POINTS: u32 = 100000;
    println!("The value of constant MAX_POINTS is: {}", MAX_POINTS);
    println!("The value of constant PI is: {}", PI);

    // shadowing
    let b = 5;
    println!("The value of b is: {}", b);
    let b = b + 1;
    println!("The value of b is: {}", b);
    let b = b * 2;
    println!("The value of b is: {}", b);
    let b = "bbbbbb";
    println!("The value of b is: {}", b);

    return true;
}

fn data_types(name: String) -> (bool, u32) {

    println!("function: {}", name);

    // numberic operations
    let sum = 5 + 10;
    println!("sum: {}", sum);
    let difference = 95.5 - 4.3;
    println!("difference: {}", difference);
    let product = 4 * 30;
    println!("product: {}", product);
    let quotient = 56.7 / 32.2;
    println!("quotient: {}", quotient);
    let remainder = 43 % 5;
    println!("remainder: {}", remainder);

    // char
    let c = 'z';
    let z = 'ℤ';
    let heart_eyed_cat = '😻';
    println!("c: {}, z: {}, cat: {}", c, z, heart_eyed_cat);

    // tuple (once declared, they cannot grow or shrink in size)
    let tup: (i32, f64, u8, bool) = (500, 6.4, 1, false);
    let (a, b, c, d) = tup;
    println!("[tup] a: {}, b: {}, c: {}, d: {}", a, b, c, d);
    println!("tup[0]: {}", tup.0);

    // array (fixed length)
    let months = ["January", "February", "March", "April", "May", "June", "July",
              "August", "September", "October", "November", "December"];
    println!("months[0]: {}", months[0]);
    let arr_a = [5, 4, 3, 2, 1];
    println!("arr_a[0]: {}", arr_a[0]);

    let arr_b: [i32; 5] = [1, 2, 3, 4, 5];
    println!("arr_b[0]: {}", arr_b[0]);

    let arr_c = [3; 5];  // the same as "let arr_c = [3, 3, 3, 3, 3];"
    println!("arr_c[0]: {}", arr_c[0]);

    let num = 10;
    plus_one(num);
    println!("num plus_one: {}", num);

    return (true, 1);
}

fn plus_one(x: i32) {
    x + 1;
}

fn main() {

    let mut case = "test variables";
    let res = variables(case.to_string());
    println!("res: {}.", res);
    println!("");
    println!("");
    println!("");
    println!("");

    case = "test data types";
    let res = data_types(case.to_string());
    println!("res[0]: {}, res[1]: {}", res.0, res.1);
    println!("");
    println!("");
    println!("");
}
