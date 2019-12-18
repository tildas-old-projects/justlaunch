extern crate xdg;
use xdg::BaseDirectories;

let xdg_dir = xdg::BaseDirectories::with_prefix("justlaunchd").unwrap();

xdg_dir.place_config_file("config.json").expect("Cannot create daemon config")