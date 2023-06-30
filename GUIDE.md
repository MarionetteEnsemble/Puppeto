Welcome to the fun and exciting world of Puppeto programming! In this beginner's guide, we'll explore the basics of Puppeto. We'll also learn about various data types and built-in functions and how to use them in your code.

# Syntax

1. **Comments:**

In Puppeto, you can add comments to your code using the `//` symbol. Comments are not executed by the interpreter; they are there to help you and other programmers understand your code.

```puppeto
// This is a comment in Puppeto.
```

2. **Variables and Assignments:**

In Puppeto, you can store values in variables using the `let` keyword. Variables must be named in snake_case. To assign a value to a variable, use the `=` sign.

```puppeto
let age = 25;
let user_name = "John Doe";
```

3. **Arithmetic and Assignment Operators:**

Puppeto supports various arithmetic and assignment operators that make it easy to manipulate variables.

```puppeto
let x = 10;
x++; // Equivalent to x = x + 1;
x--; // Equivalent to x = x - 1;
++x;
--x;
x += 5; // Equivalent to x = x + 5;
x -= 3; // Equivalent to x = x - 3;
x *= 2; // Equivalent to x = x * 2;
x /= 4; // Equivalent to x = x / 4;
x %= 3; // Equivalent to x = x % 3;
```

There are also `&=`, `|=`, `^=`, and `**=`.

4. **Binary Operators:**

Puppeto also supports various binary operators for arithmetic operations.

```puppeto
let a = 5;
let b = 2;
let c = a + b; // Addition
let d = a - b; // Subtraction
let e = a * b; // Multiplication
let f = a / b; // Division
let g = a ** b; // Exponentiation
let h = a % b; // Modulus (remainder of division)
```

You can also use `&`, `|`, `^`, `<<`, and `>>`.

5. **Logical Operators:**

Puppeto allows you to perform logical operations using logical operators.

```puppeto
let is_true = true;
let is_false = false;

let result_and = is_true && is_false; // Logical AND
let result_or = is_true || is_false; // Logical OR
let result_not = !is_true; // Logical NOT
```

There are also unary operators; `!` and `-`!

6. **Data Types:**

Puppeto supports various data types:

- `nil`: Represents the absence of a value.
- `boolean`: Represents true or false.
- `float`: Represents decimal numbers (float64).
- `integer`: Represents whole numbers (int64).
- `string`: Represents text enclosed in double or single quotes.
- `function`: Represents a block of code that can be called by its name.
- `array`: Represents a collection of elements.
- `object`: Represents key-value pairs enclosed in curly braces.

```puppeto
let no_value = nil;
let is_raining = true;
let temperature = 27.5;
let count = 42;
let name = "Alice";
fn say_hello() {
    echo "Hello, World!";
}
fn println(s) {
    echo s;
}
fn sum_all(...ns) {
    let n = 0,
        i = 0;

    for i < get_length(ns) {
        n += ns[i];

        i++;
    }

    return n;
}
let numbers = [1, 2, 3, 4];
let person = {
    name: "John",
    age: 30,
    is_student: true,
};
```

7. **Function Calls:**

You can call a function by using its name followed by parentheses.

```puppeto
fn greet(name) {
    echo "Hello, " + name + "!";
}

greet("Alice"); // Output: "Hello, Alice!"
```



8. **Control Flow:**

Puppeto provides control flow statements to make decisions and repeat code blocks.

- `if` statement: Execute a block of code conditionally.

    ```puppeto
    if temperature > 30 {
        echo "It's hot outside!";
    } else {
        echo "Enjoy the weather!";
    }
    ```

- `for` loop: Runs code until expression is falsy.

    ```puppeto
    let i = 0;

    for i < 10 {
        echo i + "\n";

        i++;
    }
    ```

- `try-catch` statement: Handle exceptions that may occur during code execution.

    ```puppeto
    try {
        // Code that might throw an exception
    } catch error {
        // Code to handle the exception
    }
    ```

9. **Printing Output:**

You can print output using the `echo` statement.

```puppeto
echo "Hello, World!"; // Output: "Hello, World!"
```

10. **Error Handling:**

Puppeto allows you to handle errors gracefully using the `try-catch` statement.

```puppeto
try {
    // Code that might throw an exception
} catch error {
    // Code to handle the exception
}
```

# Built-in Functions

Puppeto provides several built-in functions that you can use to perform various operations. Let's explore each of them:

1. **typeof(any) -> string:**
    - Returns the type of the given value as a string.
    - Example:
        ```puppeto
        let value = 42;
        let type = typeof(value);
        ECHO type; // Output: "integer"
        ```

2. **is_nan(float64|any) -> boolean:**
    - Checks if the given value is NaN (Not a Number).
    - Example:
        ```puppeto
        let result = is_nan(10.5 / 0);
        ECHO result; // Output: true
        ```

3. **inspect(any) -> string:**
    - Returns a string representation of the given value, useful for debugging.
    - Example:
        ```puppeto
        let obj = { name: "John", age: 30 };
        let str = inspect(obj);
        ECHO str; // Output: "{name: 'John', age: 30}"
        ```

4. **to_string(any) -> string:**
    - Converts the given value to a string.
    - Example:
        ```puppeto
        let value = 42;
        let str = to_string(value);
        ECHO str; // Output: "42"
        ```

5. **to_integer(any) -> integer:**
    - Converts the given value to an integer.
    - Example:
        ```puppeto
        let value = "42";
        let integer = to_integer(value);
        ECHO integer; // Output: 42
        ```

6. **to_boolean(any) -> boolean:**
    - Converts the given value to a boolean.
    - Example:
        ```puppeto
        let value = "true";
        let bool = to_boolean(value);
        ECHO bool; // Output: true
        ```

7. **parse_float(string|any) -> float:**
    - Parses the given value as a float (decimal number).
    - Example:
        ```puppeto
        let value = "3.14";
        let floatNum = parse_float(value);
        ECHO floatNum; // Output: 3.14
        ```

8. **parse_integer(string|any) -> integer:**
    - Parses the given value as an integer (whole number).
    - Example:
        ```puppeto
        let value = "42";
        let integerNum = parse_integer(value);
        ECHO integerNum; // Output: 42
        ```

9. **parse_error(string|any) -> string:**
    - Parses the given value as an error message.
    - Example:
        ```puppeto
        let value = "Error: Something went wrong!";
        let errorMessage = parse_error(value);
        ECHO errorMessage; // Output: "Something went wrong!"
        ```

10. **import_string_as_file(file: string|any, code: string|any) -> string:**
    - Imports the given code as a file with the specified name.
    - Example:
        ```puppeto
        let fileName = "my_file.pup";
        let code = "ECHO 'Hello, World!';";
        import_string_as_file(fileName, code);
        import_file(fileName); // Output: "Hello, World!"
        ```

11. **import_file(string) -> string:**
    - Imports and executes the code from the specified file.
    - Example:
        ```puppeto
        import_file("my_script.pup");
        ```

12. **get_length(array) -> integer:**
    - Get the length of an array.
    - Example:
        ```puppeto
        get_length([]); //0
        ```

13. **get_entries(object) -> array:**
    - Returns an array of a given object's key-value pairs.
    - Example:
        ```puppeto
        get_entries({a:0}); //[["a", 0]]
        ```

14. **get_keys(object) -> array:**
    - Returns an array of a given object's keys.
    - Example:
        ```puppeto
        get_keys({a:0}); //["a"]
        ```

15. **get_values(object) -> array:**
    - Returns an array of a given object's values.
    - Example:
        ```puppeto
        get_values({a:0}); //[0]
        ```

These built-in functions provide powerful tools for manipulating and transforming data in your Puppeto programs. Experiment with them to see how they can enhance your code!

# Program
Your code should be inside `<?php` and `?>`.
```puppeto
<?pup

?>
```
You can put any HTML tags in `.pup` files just like how PHP works.
<br />
You also may do this
```puppeto
<?pup if exp { ?>
Hello
<?pup } else { ?>
World!
<?pup } ?>
```

Remember, practice is key to mastering any programming language. So feel free to experiment with these concepts and have fun exploring the world of Puppeto programming!