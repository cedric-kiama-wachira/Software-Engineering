# 🚀 Enhancing Data Storage & Manipulation in Go

## Introduction
Now that we understand how to store data in computer memory, it's time to explore how to enhance and manipulate it effectively. This introduces us to **Types**.

According to the [Go Specification](https://golang.org/ref/spec), there are four general categories of types, each serving distinct purposes:

### 🏗 Categories of Types in Go:
1. **Built-In Types**
2. **Composite Types** (*struct, array, slice, map, chan, interface, func*)
3. **Custom Types**
4. **Pointer Types**

---

## 📌 Enhancing Our Program

To make our program more functional and technically robust, we will:

- ✅ Understand and use all **Built-In Types** in a basic way.
- ✅ Refactor the code to incorporate **Composite Types**.
- ✅ Further refine the code to leverage **Custom Types**.
- ✅ Ensure the final version follows **Go idiomatic best practices**.

---

## 🔍 Understanding Go Types

### **1️⃣ What Are Built-In Types?** 🏗
These foundational types are always available in Go without any additional installation. They are essential for building basic applications.

#### **Built-In Type Categories:**
- **Numbers**: Includes 14 distinct types (e.g., `int`, `float64`)
- **Strings**: Represents textual data
- **Booleans**: Represents `true` or `false` values
- **Special Types**: Includes `byte` and `rune`

📌 **Why are Built-In Types required?**
> They serve as the backbone of every Go program, allowing us to define fundamental data structures and logic.

---

### **2️⃣ What Are Composite Types?** 🔄
Composite types enable us to build more complex applications by grouping built-in and other types.

#### **Key Composite Types:**
- **`struct`** – Bundles multiple fields together
- **`array`** – Fixed-size sequence of elements of the same type
- **`slice`** – A dynamically sized alternative to arrays
- **`map`** – A key-value store
- **`chan`** – Used for communication between goroutines
- **`interface`** – Defines behavior that a type must implement
- **`func`** – Represents function signatures

📌 **Why are Composite Types required?**
> Without them, our programs would be flat, generic, and fragile, lacking the robustness required for real-world applications.

---

### **3️⃣ What Are Custom Types?** 🎭
Custom types extend the capabilities of built-in and composite types, improving clarity, type safety, and code structure.

📌 **Why are Custom Types required?**
> They make our programs more meaningful, structured, and reusable while enhancing type safety.

---

### **4️⃣ What Are Pointer Types?** 📍
Pointers store the memory addresses of other variables, making them efficient for large-scale applications.

📌 **Why are Pointer Types required?**
> They improve performance by avoiding unnecessary data copying and enabling direct memory access.

---

## 🎯 Data Points for Our Go Program

Here's how we store and manipulate data in Go:

```go
package main

import "fmt"

const fullnames string = "Cedric K. Wachira"  // Full name
var birthyear int = 1982                      // Year of birth
var age int = 2025 - birthyear                // Age in 2025 (43)
var totalFriends int = 150                    // Total friends
var dailyCoffees int8 = 2                     // Daily coffee cups
var stepsPerDay int16 = 8000                  // Daily steps
var totalBooksRead int32 = 250                // Books read lifetime
var lifetimeSeconds int64 = 1356048000        // Seconds lived by 2025
var daysSinceLastVacation uint = 180          // Days since last trip
var favoriteHour uint8 = 7                    // Favorite hour (7 AM)
var yearlyExpenses uint16 = 45000             // Yearly expenses ($)
var stepsLifetime uint32 = 103660000          // Lifetime steps
var starsVisible uint64 = 3000                // Stars visible at night
var height float32 = 1.75                     // Height in meters
var weight float64 = 82.563                   // Weight in kg
var favoriteComplex complex64 = 3 + 4i        // Favorite complex number
var lifeVector complex128 = 43.5 + 1982.0i    // Life as a complex number
var isMorningPerson bool = true               // Morning person?
var favoriteQuote string = "Code is life"     // Favorite quote
var initial byte = 'C'                        // First initial
var favoriteEmoji rune = '🚀'                 // Favorite emoji

func main() {
    fmt.Println(birthyear)
    fmt.Println(favoriteEmoji)
}
```

---

## 🎯 Next Steps
- Refactor the code to use **Composite Types**.
- Further optimize it with **Custom Types**.
- Ensure the final version is **fully idiomatic**.

🔗 *Stay tuned for future updates as we enhance our Go program!* 🚀


