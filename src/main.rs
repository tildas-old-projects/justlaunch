#[macro_use]
extern crate clap;
use clap::App;

pub fn main() {
    let schema = load_yaml!("cli.yml");
    let matches = App::from_yaml(schema).get_matches();
}
