use rand::Rng;
use std::cmp::Ordering;
use std::io;

fn main() {
    println!("Guess the number!");

    // 使用 let 语句来创建变量, 变量默认是不可变的; 在变量名前使用 mut 来使一个变量可变
    let secret_number = rand::thread_rng().gen_range(0..101);

    // loop 关键字创建了一个无限循环
    loop {
        println!("Please input your guess.");
        // 创建了一个新的可变空字符串
        let mut guess = String::new();

        io::stdin()
            // & 表示这个参数是一个引用（reference），它允许多处代码访问同一处数据，而无需在内存中多次拷贝
            .read_line(&mut guess)
            .expect("Failed to read line.");

        // parse 方法 将字符串解析成数字，通过 let guess: u32 指定数字类型,guess 后面的冒号（:）告诉 Rust 我们指定了变量的类型
        let guess: u32 = match guess.trim().parse() {
            // parse 返回一个 Result 类型，而 Result 是一个拥有 Ok 或 Err 成员的枚举
            Ok(num) => num,
            // _ 是一个通配符值
            Err(_) => {
                println!("Please type a number.");
                continue;
            }
        };
        // 花括号占位
        println!("You guessed: {}", guess);

        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small."),
            Ordering::Greater => println!("Too big."),
            Ordering::Equal => {
                println!("you win!");
                break;
            }
        }
    }
}
