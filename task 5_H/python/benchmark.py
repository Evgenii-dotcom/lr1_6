import time
import subprocess
import json
import fastcalc


def python_sum_squares(numbers):
    return sum(x*x for x in numbers)


data = list(range(1000000))

# Python test
start = time.time()
python_sum_squares(data)
end = time.time()

print("Python time:", end-start)


# Rust test
start = time.time()
fastcalc.sum_squares(data)
end = time.time()

print("Rust time:", end-start)


# Go test
proc = subprocess.Popen(
    ["../go/calculator"],
    stdin=subprocess.PIPE,
    stdout=subprocess.PIPE
)

input_data = json.dumps({"numbers": data})

start = time.time()

output, _ = proc.communicate(input_data.encode())

end = time.time()

print("Go time:", end-start)