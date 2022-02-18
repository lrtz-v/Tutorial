fn main() {
    another_function();
    another_function_with_mutil_parameters(1, 2);
    statements_and_expressions();

    let x = five();
    println!("The value of x is: {}", x);
    println!("The result of plus_one is: {}", plus_one(x));
}

// Rust doesn’t care where you define your functions, only that they’re defined somewhere
fn another_function() {
    println!("Another function.");
}

// 在函数签名中，必须 声明每个参数的类型, 参数是特殊变量，是函数签名的一部分
// 当函数拥有参数（形参）时，可以为这些参数提供具体的值（实参）
fn another_function_with_mutil_parameters(x: i32, y: i32) {
    println!("The value of x is: {}", x);
    println!("The value of y is: {}", y);
}

// 语句（Statements）是执行一些操作但不返回值的指令。表达式（Expressions）计算并产生一个值
fn statements_and_expressions() {
    // 语句
    let y = 6;

    // 表达式
    let y = {
        let x = 3;
        x + 1 // 表达式的结尾没有分号, 如果在表达式的结尾加上分号，它就变成了语句，而语句不会返回值
    };
}

// 在箭头（->）后声明返回值的类型
fn five() -> i32 {
    5 // 表达式
}

fn plus_one(x: i32) -> i32 {
    x + 1 // 表达式
}
