/*
引用（reference）像一个指针，因为它是一个地址，我们可以由此访问储存于该地址的属于其他变量的数据
与指针不同，引用确保指向某个特定类型的有效值
*/

fn calculate_length(s: &String) -> usize {
    // & 符号就是 引用, 它们允许你使用值但不获取其所有权
    s.len()
} // 这里，s 离开了作用域。但因为它并不拥有引用值的所有权，所以什么也不会发生
  // 当函数使用引用而不是实际值作为参数，无需返回值来交还所有权，因为就不曾拥有所有权


fn change(some_string: &mut String) {
    some_string.push_str(", world");
}

fn main() {
    let s1 = String::from("hello");
    let len = calculate_length(&s1); // & 符号就是 引用, 它们允许你使用值但不获取其所有权
    // 创建一个引用的行为称为 借用（borrowing）
    println!("The length of '{}' is {}.", s1, len);

    // 引用不允许修改值，必须修改为可变引用
    // PS：在同一时间只能有一个对某一特定数据的可变引用
    let mut s = String::from("hello");
    change(&mut s);
    println!("The length of '{}' is {}.", s, s.len());
}
