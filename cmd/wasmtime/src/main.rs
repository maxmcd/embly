//! CLI tool to use the functions provided by the [wasmtime](../wasmtime/index.html)
//! crate.
//!
//! Reads Wasm binary files (one Wasm module per file), translates the functions' code to Cranelift
//! IL. Can also executes the `start` function of the module by laying out the memories, globals
//! and tables, then emitting the translated code with hardcoded addresses to memory.

#![deny(
    missing_docs,
    trivial_numeric_casts,
    unused_extern_crates,
    unstable_features
)]
#![warn(unused_import_braces)]
#![cfg_attr(feature = "clippy", plugin(clippy(conf_file = "../../clippy.toml")))]
#![cfg_attr(
    feature = "cargo-clippy",
    allow(clippy::new_without_default, clippy::new_without_default_derive)
)]
#![cfg_attr(
    feature = "cargo-clippy",
    warn(
        clippy::float_arithmetic,
        clippy::mut_mut,
        clippy::nonminimal_bool,
        clippy::option_map_unwrap_or,
        clippy::option_map_unwrap_or_else,
        clippy::unicode_not_nfc,
        clippy::use_self
    )
)]

extern crate cranelift_codegen;
extern crate cranelift_native;
extern crate docopt;
extern crate wasmtime_execute;
extern crate wasmtime_runtime;
#[macro_use]
extern crate serde_derive;
extern crate file_per_thread_logger;
extern crate pretty_env_logger;
extern crate wabt;

use cranelift_codegen::isa::TargetIsa;
use cranelift_codegen::settings;
use cranelift_codegen::settings::Configurable;
use docopt::Docopt;
use std::error::Error;
use std::fs::File;
use std::io;
use std::io::prelude::*;
use std::path::Path;
use std::path::PathBuf;
use std::process::exit;
use wasmtime_execute::{ActionOutcome, InstancePlus, JITCode, NullResolver, Resolver};
use wasmtime_runtime::{Export};

static LOG_FILENAME_PREFIX: &str = "cranelift.dbg.";

const USAGE: &str = "
Wasm runner.
Takes a binary (wasm) or text (wat) WebAssembly module and instantiates it,
including calling the start function if one is present. Additional functions
given with --invoke are then called.
Usage:
    wasmtime [-od] <file>...
    wasmtime [-od] <file>... --invoke=<fn>
    wasmtime --help | --version
Options:
    --invoke=<fn>       name of function to run
    -o, --optimize      runs optimization passes on the translated functions
    -d, --debug         enable debug output on stderr/stdout
    -h, --help          print this help message
    --version           print the Cranelift version
";

#[derive(Deserialize, Debug, Clone)]
struct Args {
    arg_file: Vec<String>,
    flag_optimize: bool,
    flag_debug: bool,
    flag_invoke: Option<String>,
}


struct MyResolver {}


impl Resolver for MyResolver {
    fn resolve(&mut self, _module: &str, _field: &str) -> Option<Export> {
        println!("{}", _module);
        println!("{}", _field);
        None
    }
}

fn read_to_end(path: PathBuf) -> Result<Vec<u8>, io::Error> {
    let mut buf: Vec<u8> = Vec::new();
    let mut file = File::open(path)?;
    file.read_to_end(&mut buf)?;
    Ok(buf)
}

fn main() {
    let args: Args = Docopt::new(USAGE)
        .and_then(|d| {
            d.help(true)
                .version(Some(String::from("0.0.0")))
                .deserialize()
        })
        .unwrap_or_else(|e| e.exit());

    if args.flag_debug {
        pretty_env_logger::init();
    } else {
        file_per_thread_logger::initialize(LOG_FILENAME_PREFIX);
    }

    let isa_builder = cranelift_native::builder().unwrap_or_else(|_| {
        panic!("host machine is not a supported target");
    });
    let mut flag_builder = settings::builder();

    // Enable verifier passes in debug mode.
    if cfg!(debug_assertions) {
        flag_builder.enable("enable_verifier").unwrap();
    }

    // Enable optimization if requested.
    if args.flag_optimize {
        flag_builder.set("opt_level", "best").unwrap();
    }

    let isa = isa_builder.finish(settings::Flags::new(flag_builder));

    for filename in &args.arg_file {
        let path = Path::new(&filename);
        match handle_module(&args, path, &*isa) {
            Ok(()) => {}
            Err(message) => {
                let name = path.as_os_str().to_string_lossy();
                println!("error while processing {}: {}", name, message);
                exit(1);
            }
        }
    }
}

fn handle_module(args: &Args, path: &Path, isa: &TargetIsa) -> Result<(), String> {
    let mut data =
        read_to_end(path.to_path_buf()).map_err(|err| String::from(err.description()))?;
    // if data is using wat-format, first convert data to wasm
    if !data.starts_with(&[b'\0', b'a', b's', b'm']) {
        data = wabt::wat2wasm(data).map_err(|err| String::from(err.description()))?;
    }
    let mut resolver = MyResolver {};
    let mut jit_code = JITCode::new();
    let mut instance_plus =
        InstancePlus::new(&mut jit_code, isa, &data, &mut resolver).map_err(|e| e.to_string())?;

    if let Some(ref f) = args.flag_invoke {
        match instance_plus
            .invoke(&mut jit_code, isa, &f, &[])
            .map_err(|e| e.to_string())?
        {
            ActionOutcome::Returned { .. } => {}
            ActionOutcome::Trapped { message } => {
                return Err(format!("Trap from within function {}: {}", f, message));
            }
        }
    }

    Ok(())
}

