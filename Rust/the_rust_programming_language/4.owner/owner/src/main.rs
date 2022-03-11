// 移动
fn var_move() {
    // String 由三部分组成:一个指向存放字符串内容内存的指针，一个长度，和一个容量
    let s1 = String::from("hello");
    // 将 s1 赋值给 s2，String 的数据被复制了，这意味着我们从栈上拷贝了它的指针、长度和容量, 并没有复制指针指向的**堆上数据**
    // 二次释放问题：当变量离开作用域后，Rust 自动调用 drop 函数并清理变量的堆内存。当 s2 和 s1 离开作用域，他们都会尝试释放相同的内存。为了确保内存安全，在 let s2 = s1 之后，Rust 认为 s1 不再有效，因此 Rust 不需要在 s1 离开作用域后清理任何东西
    let s2 = s1; // 移动 (Move)，不同于浅拷贝的是，第一个变量无效了，可以解读为 s1 被 移动 到了 s2 中， 只有 s2 是有效的
                 // println!("{}, world!", s1);  // error
    println!("{}, world!", s2);

    // 整型这样的在编译时已知大小的类型被整个存储在栈上，所以拷贝其实际的值是快速的, 没有理由在创建变量 y 后使 x 无效
    // 这里没有深浅拷贝的区别
    let mut x = 5;
    let y = x;
    println!("x: {}, y: {}", x, y);
    x = 10;
    println!("x: {}, y: {}", x, y);

    /*
        Rust 有一个叫做 Copy trait 的特殊注解，可以用在类似整型这样的存储在栈上的类型上
        如果一个类型实现了 Copy trait，那么一个旧的变量在将其赋值给其他变量后仍然可用
        Rust 不允许自身或其任何部分实现了 Drop trait 的类型使用 Copy trait
        如果我们对其值离开作用域时需要特殊处理的类型使用 Copy 注解，将会出现一个编译时错误
        一些 Copy 的类型：
            所有整数类型，比如 u32。
            布尔类型，bool，它的值是 true 和 false。
            所有浮点数类型，比如 f64。
            字符类型，char。
            元组，当且仅当其包含的类型也都实现 Copy 的时候。比如，(i32, i32) 实现了 Copy，但 (i32, String) 就没有
    */
}

// 克隆
fn var_clone() {
    let s1 = String::from("hello");
    let mut s2 = s1.clone(); // 深度复制 String 中堆上的数据
    println!("s1 = {}, s2 = {}", s1, s2);
    s2.push_str(", this is s2");
    println!("s1 = {}, s2 = {}", s1, s2);
}

// 所有权与函数调用
fn owner_and_function() {
    fn takes_ownership(mut some_string: String) {
        // some_string 进入作用域
        some_string.push_str(" in funcetion");
        println!("{}", some_string);
    } // 这里，some_string 移出作用域并调用 `drop` 方法。占用的内存被释放

    fn makes_copy(mut some_integer: i32) {
        // some_integer 进入作用域
        some_integer += 1;
        println!("{}", some_integer);
    } // 这里，some_integer 移出作用域。没有特殊之处

    let s = String::from("hello"); // s 进入作用域
    takes_ownership(s); // s 的值移动到函数里
                        // println!("{}", s); // s 到这里不再有效, err: move occurs

    let x = 5; // x 进入作用域
    makes_copy(x); // x 应该移动函数里
                   // 但 i32 是 Copy 的，所以在后面可继续使用 x
    println!("{}", x);
}

// 返回值与作用域
fn return_values_and_scope() {
    fn gives_ownership() -> String {
        // gives_ownership 会将返回值移动给, 调用它的函数
        let some_string = String::from("yours"); // some_string 进入作用域.
        some_string // 返回 some_string, 并移出给调用的函数
    }

    // takes_and_gives_back 将传入字符串并返回该值
    fn takes_and_gives_back(mut a_string: String) -> String {
        // a_string 进入作用域
        a_string.push_str(", in takes_and_gives_back");
        a_string // 返回 a_string 并移出给调用的函数
    }

    let s1 = gives_ownership(); // gives_ownership 将返回值转移给 s1
    let s2 = String::from("hello"); // s2 进入作用域
    println!("s2: {}", s2);
    let s3 = takes_and_gives_back(s2); // s2 被移动到takes_and_gives_back 中,它也将返回值移给 s3
    println!("s3: {}", s3);

    // 这里, s3 移出作用域并被drop, s2 也移出作用域，但已被移走，所以什么也不会发生, s1 离开作用域并被drop
}

fn main() {
    var_move();
    var_clone();
    owner_and_function();
    return_values_and_scope();
}
