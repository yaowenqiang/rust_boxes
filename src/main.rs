use std::error::Error;
fn main() {
    //println!("{}", std::fs::read_to_string("/etc/issue").unwrap());
    //println!("{}", std::fs::read_to_string("/etc/issue").unwrap());
    println!("{}", read_issue().unwrap());
}

fn read_issue() -> Result<String, Box<dyn Error>> {
    //std::fs::read_to_string("/etc/issue").unwrap()
    Ok(std::fs::read_to_string("/etc/issue")?)
}
