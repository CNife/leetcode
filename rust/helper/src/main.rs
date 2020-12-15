use regex::Regex;
use std::{fs, io, process};

fn main() -> io::Result<()> {
    let lib_rs = "../src/lib.rs";
    let finished_lib_rs = "../finished/src/lib.rs";
    let module_regex = Regex::new(r"mod\s+(.*);").unwrap();

    let lib_rs_content = fs::read_to_string(lib_rs)?;
    let mut finished_modules: Vec<String> = scan_finished()?;
    for capture in module_regex.captures_iter(&lib_rs_content) {
        let module_name = &capture[1];
        fs::rename(
            format!("../src/{}.rs", module_name),
            format!("../finished/src/{}.rs", module_name),
        )?;
        finished_modules.push(format!("mod {};", module_name));
    }

    fs::write(lib_rs, "")?;

    finished_modules.sort_unstable();
    let mut new_finished_lib_rs = String::new();
    for module in finished_modules {
        new_finished_lib_rs.push_str("pub mod ");
        new_finished_lib_rs.push_str(&module);
        new_finished_lib_rs.push_str(";\n");
    }
    fs::write(finished_lib_rs, new_finished_lib_rs)?;

    process::Command::new("cargo")
        .arg("test")
        .current_dir("../finished")
        .spawn()?
        .wait()?;
    process::Command::new("cargo")
        .arg("fmt")
        .current_dir("..")
        .spawn()?
        .wait()?;

    Ok(())
}

fn scan_finished() -> io::Result<Vec<String>> {
    let mut result = Vec::new();
    for entry in fs::read_dir("../finished/src")? {
        let entry = entry?;
        if entry.metadata()?.is_file() {
            let path = entry.path();
            match (path.extension(), path.file_stem()) {
                (Some(ext), Some(stem)) => {
                    if ext == "rs" && stem != "lib" {
                        result.push(stem.to_string_lossy().into_owned());
                    }
                }
                _ => {}
            }
        }
    }
    Ok(result)
}
