import time
from sum_loop import sum_numbers

start_time = time.time()
total = sum_numbers()
end_time = time.time()
cython_duration = end_time - start_time

print(f"Total: {total}, Cython execution time: {cython_duration:.6f} seconds")
