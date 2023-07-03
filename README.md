# Puppeto@0.3

Puppeto is a programming language designed to provide a simplified and explicit alternative to PHP, removing unnecessary features like function type-checking, constants, and more. With its focus on maintaining implicit dynamic typing and simplicity, Puppeto aims to streamline the development process for web developers who prefer a lightweight and straightforward language.

Puppeto embraces a minimalistic approach by offering only eight primitive types: String, Integer, Float, Boolean, Function, Array, Object, and Nil. By limiting the number of types, Puppeto encourages developers to focus on the core functionality of their code while still allowing flexibility in handling different data structures and values.

One of Puppeto's distinctive features is its requirement to declare variables before their usage. This approach promotes explicitness and helps catch potential errors early in the development process. Developers must declare variables before using them, ensuring clear intentions and reducing ambiguity.

To enhance ease of use, Puppeto supports optional semicolons, allowing developers to write code with or without them, based on their personal preference. This flexibility accommodates different coding styles while maintaining code readability.

Puppeto is specifically designed to be used in conjunction with HTML, seamlessly integrating with front-end web development. It provides a straightforward syntax for embedding Puppeto code within HTML, allowing developers to enhance their web applications with dynamic functionality.

Whether you are a web developer seeking a more streamlined alternative to PHP or someone looking to explore a new language for web development, Puppeto offers simplicity, clarity, and a reduced feature set that focuses on what matters most—building efficient and effective web applications.

# How to Use
```
go run cli/html/main.go <file>
```

Though, there's also another CLI file. 
```
go run cli/core/main.go <file>
```
Unlike the first one, this one will expect the file starts with `<?pup`.

# Guide
Please check GUIDE.md m8~

# File Extension
`.pup`