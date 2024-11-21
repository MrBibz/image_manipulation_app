def test_if(value):
    if value > 0:
        print(f"{value} is positive")
    elif value < 0:
        print(f"{value} is negative")
    else:
        print(f"{value} is zero")

def test_switch(value):
    if value > 0:
        print(f"Value {value} is positive")
    elif value < 0:
        print(f"Value {value} is negative")
    else:
        print(f"Value {value} is zero")

def test_for():
    print("Counting from 1 to 10 using a for loop:")
    for i in range(1, 11):
        print(i)

def test_while():
    print("Counting from 1 to 10 using a while loop:")
    i = 1
    while i <= 10:
        print(i)
        i += 1

def test_list():
    print("Creating a list of 5 elements:")
    lst = [1, 2, 3, 4, 5]
    print(lst)

    print("Adding an element to the list:")
    lst.append(6)
    print(lst)

    print("Removing the last element of the list:")
    lst.pop()
    print(lst)

def test_dict():
    print("Creating a dictionary of 3 elements:")
    d = {
        "one": 1,
        "two": 2,
        "three": 3
    }
    print(d)

    print("Adding an element to the dictionary:")
    d["four"] = 4
    print(d)

    print("Removing an element from the dictionary:")
    del d["two"]
    print(d)

def main():
    # Tests conditional statements
    print("\nTest if statement:")
    test_if(5)
    test_if(-5)
    test_if(0)

    print("\nTest switch statement:")
    test_switch(5)
    test_switch(-5)
    test_switch(0)

    # Test iteration statements
    print("\nTest for statement:")
    test_for()

    print("\nTest while statement:")
    test_while()

    # Test collections
    print("\nTest list:")
    test_list()

    print("\nTest dictionary:")
    test_dict()

    # Test functions
    print("\nTest functions:")
    print("\nEvery test is in a function in this file.")

if __name__ == "__main__":
    main()
