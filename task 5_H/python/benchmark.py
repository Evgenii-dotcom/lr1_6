import time
import subprocess
import json
from typing import List
import fastcalc


def python_sum_squares(numbers: List[int]) -> int:
    return sum(x * x for x in numbers)


def measure_time(func, *args, **kwargs):
    start = time.time()
    result = func(*args, **kwargs)
    end = time.time()
    return result, end - start


def run_go_sum_squares(numbers: List[int], executable_path: str) -> int:
    input_data = json.dumps({"numbers": numbers})
    proc = subprocess.Popen(
        [executable_path],
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
    )
    output, error = proc.communicate(input_data.encode())

    if proc.returncode != 0:
        raise RuntimeError(f"Go process failed: {error.decode()}")

    result_json = json.loads(output.decode())
    return result_json["sum"]


def main():
    data = list(range(1_000_000))

    _, python_time = measure_time(python_sum_squares, data)
    print(f"Python time: {python_time:.3f} seconds")

    _, rust_time = measure_time(fastcalc.sum_squares, data)
    print(f"Rust time: {rust_time:.3f} seconds")

    go_path = "../go/calculator"
    _, go_time = measure_time(run_go_sum_squares, data, go_path)
    print(f"Go time: {go_time:.3f} seconds")


if __name__ == "__main__":
    main()