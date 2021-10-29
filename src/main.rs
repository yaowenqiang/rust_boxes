use std::error::Error;

struct MyError {
    value: u32,
}
fn main() {
    //println!("{}", std::fs::read_to_string("/etc/issue").unwrap());
    //println!("{}", std::fs::read_to_string("/etc/issue").unwrap());

    let e = MyError {
        value: 32
    };
    let e_ptr: *const MyError = &e;
    let e_ref: &MyError = &e;

    //println!("{}", read_issue().unwrap());

    dbg!(std::mem::size_of_val(&e_ptr));

    print_error(e_ptr);

    dbg!(std::mem::size_of_val(&e_ref));

    print_error2(e_ref);
    print_error3(e_ref);

    let e2: *const MyError = std::ptr::null();

    let e_ref2: &MyError = unsafe {&*e2};
    print_error2(e_ref2);
}

fn read_issue() -> Result<String, Box<dyn Error>> {
    //std::fs::read_to_string("/etc/issue").unwrap()
    Ok(std::fs::read_to_string("/etc/issue")?)
}

fn print_error(e: *const MyError) {
    if e != std::ptr::null() {
        println!("MyError (value = {})", unsafe{(*e).value});
    }
}

fn print_error2(e:  &MyError) {
    println!("MyError (value = {})", (*e).value);
}

fn print_error3(e:  &MyError) {
    println!("MyError (value = {})", e.value);
}
