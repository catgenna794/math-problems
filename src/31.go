def is_even(n):
    """
    Check if an integer n is even.
    
    Parameters:
    n (int): An integer to check.
    
    Returns:
    bool: True if n is even, False otherwise.
    """
    return n % 2 == 0

# Example usage and test the function
print(is_even(4))  # Output: True
print(is_even(5))  # Output: False
