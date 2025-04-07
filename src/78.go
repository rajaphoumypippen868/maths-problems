def is_sum_of_two_close_numbers(num1, num2):
    """
    Check if there are two numbers in nums that sum up to 'num1' or 'num2'.
    
    Args:
    - num1 (int): The target number.
    - num2 (int): The other target number.
    
    Returns:
    - bool: True if there exist two numbers in the list such that their sum equals the specified numbers, otherwise False.
    """
    # Example nums for testing
    example_nums = [5, 10, 3, 4, 8]
    for num1 in example_nums:
        for num2 in example_nums:
            if abs(num1 - num2) == 1 or abs(num1 + num2) == 5:  # Check if their sum is the target
                return True
    return False

# Example usage and check function
print(is_sum_of_two_close_numbers(7, 8))  # Expected output: True
