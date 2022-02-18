fn main() {
    if_control_flow();
    loop_control_flow();
    while_control_flow();
    for_control_flow();
}

fn if_control_flow() {
    let number = 3;

    // Note: Rust will not automatically try to convert non-Boolean types to a Boolean
    if number < 5 {
        println!("condition was true");
    } else {
        println!("condition was false");
    }

    if number != 0 {
        println!("number was something other than zero");
    }

    if number % 4 == 0 {
        println!("number is divisible by 4");
    } else if number % 3 == 0 {
        println!("number is divisible by 3");
    } else if number % 2 == 0 {
        println!("number is divisible by 2");
    } else {
        println!("number is not divisible by 4, 3, or 2");
    }

    // Using if in a let Statement
    let condition = true;
    let number = if condition { 5 } else { 6 };
    println!("The value of number is: {}", number);
    // Note: results from each arm of the if must be the same type
    // let number = if condition { 5 } else { "six" };
}

/*
Rust 有三种循环：loop、while 和 for

loop 关键字告诉 Rust 一遍又一遍地执行一段代码直到你明确要求停止

*/
fn loop_control_flow() {
    // loop
    let mut i = 1;
    loop {
        println!("{}.again!", i);
        i += 1;
        if i > 3 {
            break;
        }
    }

    let mut count = 0;
    'counting_up: loop {
        println!("count = {}", count);
        let mut remaining = 10;

        loop {
            println!("remaining = {}", remaining);
            if remaining == 9 {
                break; // 退出内层循环
            }
            if count == 2 {
                break 'counting_up; // 退出外层循环
            }
            remaining -= 1;
        }

        count += 1;
    }
    println!("End count = {}", count);

    // 从循环返回值
    let mut counter = 0;
    let result = loop {
        counter += 1;

        if counter == 10 {
            break counter * 2;
        }
    };
    println!("The result is {}", result);
}

fn while_control_flow() {
    let mut number = 3;
    while number != 0 {
        println!("number--: {}!", number);
        number -= 1;
    }
    println!("LIFTOFF!!!");

    // range list
    let a = [10, 20, 30, 40, 50];
    let mut index = 0;
    while index < 5 {
        println!("[while] the value is: {}", a[index]);
        index += 1;
    }
}

fn for_control_flow() {
    let a = [10, 20, 30, 40, 50];
    for element in a.iter() {
        println!("[for] the value is: {}", element);
    }
    for element in a {
        println!("the value is: {}", element);
    }

    // range number
    for number in (1..4).rev() {
        println!("range number {}.", number);
    }
}
