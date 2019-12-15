#[macro_use]
extern crate clap;
use clap::App;
use std::os::unix::net::UnixStream;

pub fn main() {
    let schema = load_yaml!("cli.yml");
    let matches = App::from_yaml(schema).get_matches();
    let mut daemon = UnixStream::connect("/tmp/justlaunchd.sock").unwrap();
}
