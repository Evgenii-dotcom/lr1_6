use pyo3::prelude::*;

fn calculate_sum_of_squares(numbers: &[i64]) -> i64 {
    numbers.iter().map(|x| x * x).sum()
}

#[pyfunction]
fn sum_squares(numbers: Vec<i64>) -> PyResult<i64> {
    Ok(calculate_sum_of_squares(&numbers))
}

#[pymodule]
fn fastcalc(_py: Python, m: &PyModule) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(sum_squares, m)?)?;
    Ok(())
}