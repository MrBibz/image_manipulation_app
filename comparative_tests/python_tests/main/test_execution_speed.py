import time

def hash_fnv1a(input_string):
    """
    Hash function that takes a string input and returns a 64-bit hash value.
    Implements a variant of FNV-1a hash with additional bit manipulations.
    """
    # Initial hash value (FNV-1a offset basis)
    hash_value = 0xcbf29ce484222325
    # FNV-1a prime
    prime = 0x100000001b3

    # Iterate over each byte in the input string
    for char in input_string:
        hash_value ^= ord(char)  # XOR the hash with the byte value
        hash_value *= prime      # Multiply the hash by the FNV-1a prime
        hash_value &= 0xFFFFFFFFFFFFFFFF  # Ensure 64-bit overflow
        # Rotate the hash left by 13 bits
        hash_value = ((hash_value << 13) & 0xFFFFFFFFFFFFFFFF) | (hash_value >> 51)
        # XOR the hash with itself shifted right by 7 bits
        hash_value ^= hash_value >> 7

    return hash_value  # Return the final hash value

def main():
    """
    Main function to test the hash function's execution speed.
    """
    data = "Here is a string of characters in order to test Python's execution speed."  # Input data to be hashed

    hash_value = 0  # Variable to store the hash value
    start_time = time.time()  # Record the start time

    # Loop to hash the data 100,000 times
    for _ in range(100000):
        hash_value = hash_fnv1a(data)  # Compute the hash value

    elapsed_time = time.time() - start_time  # Calculate the elapsed time

    # Print the final hash value and the execution time
    print("Hash value:", hash_value)
    print("Execution time: {:.6f} seconds".format(elapsed_time))

if __name__ == "__main__":
    main()
