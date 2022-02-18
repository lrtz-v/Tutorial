fn main() {
    var();
    const_var();
    shadow_var();
    data_type();
}

fn var() {
    // 变量
    let y = 5;
    println!("The value of y is: {}", y);
    // cannot assign twice to immutable variable
    // y = 6;

    // 在变量名之前加 mut 来使其可变
    let mut m = 5;
    println!("The value of m is: {}", m);
    m = 6;
    println!("The value of m is: {}", m);
}

fn const_var() {
    // 常量
    // 声明常量使用 const 关键字而不是 let，并且 必须 注明值的类型，不允许对常量使用 mut
    // 常量可以在任何作用域中声明，包括全局作用域，这在一个值需要被很多部分的代码用到时很有用
    // 常量只能被设置为常量表达式，而不可以是其他任何只能在运行时计算出的值
    const THREE_HOURS_IN_SECONDS: u32 = 60 * 60 * 3;
    println!(
        "The value of THREE_HOURS_IN_SECONDS is: {}",
        THREE_HOURS_IN_SECONDS
    );
}

fn shadow_var() {
    // 变量隐藏
    let x = 5;
    let x = x + 1;
    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {}", x);
    }
    println!("The value of x is: {}", x);

    // 再次使用 let 时，实际上创建了一个新变量，我们可以改变值的类型，并且复用这个名字
    let spaces = "   ";
    let spaces = spaces.len();
    println!("The value of spaces is: {}", spaces);
}

fn data_type() {
    /*
    标量类型：标量（scalar）类型代表一个单独的值。Rust 有四种基本的标量类型：整型、浮点型、布尔类型和字符类型

    整型：
    长度	有符号	无符号
    8-bit	i8	u8
    16-bit	i16	u16
    32-bit	i32	u32
    64-bit	i64	u64
    128-bit	i128	u128
    arch	isize	usize   # isize 和 usize 类型依赖运行程序的计算机架构：64 位架构上它们是 64 位的， 32 位架构上它们是 32 位的
    数字类型的数字字面值允许使用类型后缀，例如 57u8 来指定类型，同时也允许使用 _ 做为分隔符以方便读数，例如1_000，它的值与你指定的 1000 相同

    浮点数：
    Rust 的浮点数类型是 f32 和 f64，分别占 32 位和 64 位。默认类型是 f64
    */
    let a = 2.0; // f64
    println!("The value of a is: {}", a);
    let b: f32 = 3.0; // f32
    println!("The value of b is: {}", b);

    // 数值运算
    // 加法
    let sum = 5 + 10;
    println!("The value of sum is: {}", sum);
    // 减法
    let difference = 95.5 - 4.3;
    println!("The value of difference is: {}", difference);
    // 乘法
    let product = 4 * 30;
    println!("The value of product is: {}", product);
    // 除法
    let quotient = 56.7 / 32.2;
    println!("The value of quotient is: {}", quotient);
    let floored = 2 / 3; // 结果为 0
    println!("The value of floored is: {}", floored);
    // 取余
    let remainder = 43 % 5;
    println!("The value of remainder is: {}", remainder);

    /*
    布尔型：
    布尔类型有两个可能的值：true 和 false
    */
    let t = true;
    println!("The value of t is: {}", t);
    let f: bool = false; // 显式指定类型注解
    println!("The value of f is: {}", f);

    /*
    字符类型
    char 类型是语言中最原生的字母类型
    */
    let c = 'z';
    println!("The value of c is: {}", c);
    let z = 'ℤ';
    println!("The value of z is: {}", z);
    let heart_eyed_cat = '😻';
    println!("The value of heart_eyed_cat is: {}", heart_eyed_cat);

    /*
    复合类型
    Rust 有两个原生的复合类型：元组（tuple）和数组（array）

    元组: 元组是一个将多个其他类型的值组合进一个复合类型的主要方式。元组长度固定：一旦声明，其长度不会增大或缩小。
    */
    let tup: (i32, f64, u8) = (500, 6.4, 1);

    let tup = (500, 6.4, 1);
    let (x, y, z) = tup;
    println!("The value of x, y, z is: {}, {}, {}", x, y, z);

    // 可以使用点号（.）后跟值的索引来直接访问
    let x: (i32, f64, u8) = (500, 6.4, 1);
    let five_hundred = x.0;
    let six_point_four = x.1;
    let one = x.2;
    println!(
        "The value of five_hundred, six_point_four, one is: {}, {}, {}",
        five_hundred, six_point_four, one
    );

    /*
    数组类型: 数组中的每个元素的类型必须相同, Rust中的数组长度是固定的
    vector: 是标准库提供的一个 允许 增长和缩小长度的类似数组的集合类型。当不确定是应该使用数组还是 vector 的时候，那么很可能应该使用 vector


    */
    let a = [1, 2, 3, 4, 5];
    let months = [
        // 确定只会有12个元素
        "January",
        "February",
        "March",
        "April",
        "May",
        "June",
        "July",
        "August",
        "September",
        "October",
        "November",
        "December",
    ];
    let a: [i32; 5] = [1, 2, 3, 4, 5]; // i32 是每个元素的类型。分号之后，数字 5 表明该数组包含五个元素
    let a = [3; 5]; // 数组将包含 5 个元素，这些元素的值最初都将被设置为 3

    let first = a[0];
    let second = a[1];

    // 程序在索引操作中使用一个无效的值时导致 运行时 错误
}
