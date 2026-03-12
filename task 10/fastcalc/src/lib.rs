use pyo3::prelude::*;

#[pyclass]
struct Counter {
    value: i32,
}

#[pymethods]
impl Counter {
    #[new]
    fn new(value: i32) -> Self {
        Counter { value }
    }

    fn increment(&mut self) {
        self.value += 1;
    }

    fn get(&self) -> i32 {
        self.value
    }
}

#[pymodule]
fn fastcalc(_py: Python, m: &PyModule) -> PyResult<()> {
    m.add_class::<Counter>()?;
    Ok(())
}