import time

# Normal Python execution
start_time = time.time()

# For loop to sum numbers
total = 0
for i in range(1, 1000001):
    total += i

end_time = time.time()
normal_duration = end_time - start_time

print(f"Total: {total}, Normal execution time: {normal_duration:.6f} seconds")
