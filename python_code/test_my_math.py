from my_math import add

def test_add():
    """Tests the add function."""
    # 测试正数相加
    assert add(2, 3) == 5
    # 测试负数相加
    assert add(-1, -1) == -2
    # 测试正负数相加
    assert add(5, -3) == 2