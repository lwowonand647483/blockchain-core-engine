use std::collections::HashMap;

struct ContractVM {
    storage: HashMap<String, u64>,
}

impl ContractVM {
    fn new() -> Self {
        ContractVM {
            storage: HashMap::new(),
        }
    }

    fn set(&mut self, key: String, value: u64) {
        self.storage.insert(key, value);
    }

    fn get(&self, key: &str) -> u64 {
        *self.storage.get(key).unwrap_or(&0)
    }

    fn add(&mut self, a: u64, b: u64) -> u64 {
        a + b
    }
}

fn main() {
    let mut vm = ContractVM::new();
    vm.set("count".to_string(), 10);
    println!("{}", vm.get("count"));
    println!("{}", vm.add(3, 5));
}
