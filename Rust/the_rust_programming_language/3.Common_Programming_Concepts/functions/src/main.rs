fn main() {
    println!("Hello, world!");
    another_function();
    another_function_with_parameter(3);
    another_function_with_mutil_parameters(1, 2);
    another_function_with_mutil_types_parameters(1, 2, "123".to_string());

    // create new scopes, {}, is an expression
    let y = {
        let x = 3;
        x + 1
        // Expressions do not include ending semicolons. 
        // If you add a semicolon to the end of an expression, 
        // you turn it into a statement, which will then not return a value
    };
    println!("The value of y is: {}", y);  // The value of y is: 4

    let x = five();
    println!("The value of x is: {}", x);
    println!("The result of plus_one is: {}", plus_one(x));
}

// Rust doesn’t care where you define your functions, only that they’re defined somewhere
fn another_function() {
    println!("Another function.");
}

fn another_function_with_parameter(x: i32) {
    println!("The value of x is: {}", x);
}

fn another_function_with_mutil_parameters(x: i32, y: i32) {
    println!("The value of x is: {}", x);
    println!("The value of y is: {}", y);
}

fn another_function_with_mutil_types_parameters(x: i32, y: i32, z: String) {
    println!("The value of x is: {}", x);
    println!("The value of y is: {}", y);
    println!("The value of z is: {}", z);
}

fn five() -> i32 {
    5
}

fn plus_one(x: i32) -> i32 {
    x + 1
}