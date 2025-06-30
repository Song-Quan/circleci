from my_math import add

def test_add_positive_numbers():
    """Tests adding two positive numbers."""
    assert add(2, 3) == 5

def test_add_negative_numbers():
    """Tests adding two negative numbers."""
    assert add(-2, -3) == -5

def test_add_mixed_numbers():
    """Tests adding a positive and a negative number."""
    assert add(5, -3) == 2