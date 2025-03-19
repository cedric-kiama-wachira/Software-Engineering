# 🏗️ Understanding Go: Variables, Functions, and Packages

This chapter is not only about variables, as they don't work alone. It will include the basics of other blocks of Go code, plus logical components as well, and they are:

- 📦 **Packages**
- 🌎 **Namespaces**
- 🔧 **Functions**
- 🔢 **Variables, Constants, and Declarations**  
  - 4.1 **Global Variables**  
  - 4.2 **Local Variables**  
  - 4.3 **Constants**  
  - 4.4 **Variable Declarations**  
- 📂 **`main.go` File**
- 📌 **More on Variables and Types** (as listed in sections {4.1–4.4} above)
- 🏆 **Exercise**: A program that calculates age and prints it out. The system specification requires using all the building blocks explained below.
- ✅ **Idiomatic Go**: Writing a program that follows Go idioms.
- ▶️ **Running the Code**: Show how to run the program and its output.

---

## 🧐 Before Writing Code: The What & Why Questions

### 📦 What is a Package? Why Do We Need Them?
According to Go's core principles, **packages** group different pieces of code that work towards a common goal. These pieces include variables, functions, etc. Packages provide:
- **Code organization**
- **Reusability**
- **Encapsulation**
- **Namespace management**
- **Support for collaboration**

---

### 🔧 What is a Function? Why Do We Need Them?
A **function** is a block of code that performs a specific task. While not mandatory, functions are:
- Essential for **modularity and reusability**.
- Crucial for **clarifying program intent**.

---

### 🌎 What is a Namespace? Why Do We Need Them?
A **namespace** is a logical label created when we define packages. Namespaces help by:
- **Preventing conflicts** between different code pieces.
- **Organizing code efficiently** when the program runs.

---

### 🔢 What is a Variable? Why Do We Need It?
A **variable** is a container in a program that stores data. It consists of:
- A **name**
- A **data type**
- A **value**

Variables reference memory locations where values are stored, allowing access and modification during program execution.

---

### 📂 What is the `main.go` File? Why is it Required?
- **`main.go`** is the primary Go file containing the **`main()`** function.
- It serves as the **entry point** for the program.
- Without it, our Go code **will not execute**.

---

## 🔍 Expanding Further on Variables

### 6.1 🌎 What is a Global Variable? Why is it Required?
- **Global variables** are declared at the **top of a Go file**.
- They are **accessible across multiple functions**.
- **Pros:** Useful for storing shared data.
- **Cons:** May cause unintended modifications in larger programs.

---

### 6.2 🏠 What is a Local Variable? Why is it Required?
- **Local variables** are defined **inside a function**.
- They exist **only within that function** and disappear after execution.
- **Pros:** Useful for function-specific isolation.
- **Cons:** Cannot be accessed outside the function.

---

### 6.3 🔐 What is a Constant? Why is it Required?
- **Constants (`const`)** define values **that never change** in memory.
- They are useful for **mathematical operations** and **configuration settings**.
- Constants can be declared at the **top of a file** or **inside functions**.

---

### 6.4 ✍️ What is a Variable Declaration? Why is it Required?
Variable declarations can be **explicit** or **shorthand**:

#### 🔹 **Explicit Declaration**
```go
var name string = "Go Language"
```
- Uses the `var` keyword.
- Specifies the **data type** (`string`, `int`, etc.).

#### 🔹 **Shorthand Declaration**
```go
name := "Go Language"
```
- Uses `:=` for implicit typing.
- Go **infers the data type** automatically.

---

## ✅ Writing Go Code the Right Way: Idiomatic Go
Follow Go's **best practices** to write clean, readable, and efficient code.

---

## 🏆 Tying It All Together: An Idiomatic Go Program

```go
package main

import "fmt"

const fullnames = "Cedric K. Wachira" // Untyped constant—type inferred
var birthyear int = 1982 // Explicit global—stays as is

func main() {
    age := 2025 - birthyear // Short declaration for local variable
    fmt.Println(age)
}
```

---

## ▶️ Running the Code

To run the program, use the following command in your terminal:

```sh
go run main.go
```

### 🖥️ Expected Output:
```sh
43
```

🚀 **Happy Coding in Go!** 🚀

