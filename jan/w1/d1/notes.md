# Go's Type System and Variables

> **Note:** These notes assume you understand basic programming concepts (what a variable is, what a loop is, etc.). If you've done any coding before in any language, you should be fine!

## What are Types?

A **type** tells Go what kind of data a variable holds. Think of it like labeling boxes - an `int` box holds whole numbers, a `string` box holds text, etc.

Go is **strongly typed**, meaning once you say a variable is an `int`, it can only ever hold integers. You can't suddenly put text in it.

## Built-in Types

Go has a straightforward set of built-in types:

**Numbers:**
- `int` - whole numbers like `42`, `-17`, `0` (most commonly used)
- `float64` - decimal numbers like `3.14`, `-0.5`, `100.0`
- Also: `int8`, `int16`, `int32`, `int64` for different sizes

**Text:**
- `string` - text like `"hello"`, `"Go is fun"`
- `rune` - a single character
- `byte` - a single byte of data

**Boolean:**
- `bool` - either `true` or `false`

**More advanced types**:
- Arrays, slices, maps, structs, pointers, functions, interfaces, channels

## Go is Very Picky About Types

Here's something important: **Go won't automatically mix different types**. If you have an integer and a decimal number, Go won't let you add them without converting one first.

```go
var x int = 10        // whole number
var y float64 = 3.5   // decimal number
// result := x + y    // ERROR: Go says "I can't add these!"
```

Why? Because Go wants you to be explicit about what you want. Do you want the result as a decimal? Then convert `x` to a decimal first:

```go
result := float64(x) + y  // Now both are decimals, this works!
```

**Think of types as different languages** - you can't mix English and Spanish words in the same sentence without being clear about what you're doing.

## Converting Between Types

Sometimes you need to change a variable from one type to another. Go makes you do this explicitly.

### How to Convert

The pattern is simple: `newType(value)`

```go
var x int = 10
var y float64 = 3.5

// Convert x to float64, then add
result := float64(x) + y  // result is 13.5
```

### Common Conversions You'll Use

**Whole number to decimal:**
```go
age := 25
price := float64(age) * 1.5  // 37.5
```

**Decimal to whole number (be careful!):**
```go
price := 19.99
dollars := int(price)  // becomes 19 (not 20!)
// It cuts off the decimal part, doesn't round
```

**Number to text:**
```go
import "strconv"  // need this package

age := 25
message := "I am " + strconv.Itoa(age) + " years old"
// strconv.Itoa means "integer to ASCII (text)"
```

### Why So Strict?

Go forces you to be explicit to prevent bugs. Imagine if Go automatically decided what type to use - sometimes it might guess wrong, and your program would have subtle errors that are hard to find.

## Three Ways to Create Variables

All three of these do basically the same thing, just with different levels of explicitness:

**1. Tell Go the type, let it start at zero:**
```go
var name string    // empty string ""
var age int        // 0
```

**2. Give it a value, Go figures out the type:**
```go
var name = "Alice"  // Go knows this is a string
var age = 25        // Go knows this is an int
```

**3. The shortcut (most common):**
```go
name := "Alice"    // Same as #2, just shorter
age := 25          // Same as #2, just shorter
```

**Which one should you use?** Most of the time, use `:=` (option 3). It's short and clear. Use `var` (option 1) when you specifically want the zero value.

## Zero Values (Default Values)

When you create a variable in Go without giving it a value, Go automatically sets it to a sensible default. This is called a **zero value**.

**What are the zero values?**
- Numbers → `0`
- Booleans → `false`
- Text/strings → `""` (empty string)
- Other complex types → `nil` (means "nothing there")

**Example:**
```go
var count int       // automatically 0
var ready bool      // automatically false
var message string  // automatically ""

fmt.Println(count)    // prints: 0
fmt.Println(ready)    // prints: false
fmt.Println(message)  // prints: (nothing)
```

**Why this is helpful:** You never accidentally use a variable that has garbage/random data in it. Every variable starts in a safe, predictable state.

## Two Ways to Create Variables: `:=` vs `var`

Go has two ways to create variables. Both work, but they're used in different situations.

### The Quick Way: `:=`

Use this when you're creating a variable and giving it a value right away:

```go
name := "Alice"      // Go figures out this is a string
age := 25            // Go figures out this is an int
ready := true        // Go figures out this is a bool
```

**Think of `:=` as:** "Create a variable and figure out its type from what I'm giving it"

### The Explicit Way: `var`

Use this when you need more control:

```go
var count int        // Create an int, starts at 0
var price float64    // Create a decimal number, starts at 0.0
```

**Think of `var` as:** "Create a variable with this specific type"

### When to Use Which?

**Use `:=` (most common):**
- When you're giving the variable a value right away
- When the type is obvious

**Use `var` when:**
- You want the variable to start at its zero value (0, false, "")
- You need to be specific about the type

### A Common Mistake

```go
// This creates an int (whole number)
price := 0

// Later... ERROR! Can't add decimals to a whole number
price = price + 3.5  // Won't work!
```

**Fix it by being explicit:**
```go
var price float64  // Specifically a decimal number, starts at 0.0
price = price + 3.5  // Now this works!
```

## Understanding Strings and Characters (Runes)

This part is a bit tricky, but important to understand.

### The Problem with Counting Characters

Let's say you have this string:
```go
s := "Hello"
fmt.Println(len(s))  // 5 - makes sense!
```

But watch what happens with emoji or non-English characters:
```go
s := "世界"  // Two Chinese characters
fmt.Println(len(s))  // 6 - wait, what? There are only 2 characters!
```

### Why This Happens

Behind the scenes, computers store characters as numbers (called Unicode). Simple letters like "H" need only 1 byte of storage. But complex characters like "世" or emoji need 3-4 bytes each.

So `len()` counts **bytes**, not characters!

### What's a Rune?

A **rune** is Go's name for "one actual character" (like what a human would count).

Think of it this way:
- **byte** = one unit of storage (computer's perspective)
- **rune** = one character (human's perspective)

### How to Loop Through Characters

When you loop through a string with `range`, Go gives you **runes** (characters):

```go
s := "Hi世"  // 3 characters: H, i, and 世

for i, char := range s {
    fmt.Printf("%d: %c\n", i, char)
}
// Output:
// 0: H
// 1: i
// 2: 世
```

The loop sees 3 characters (good!), even though the string is 5 bytes behind the scenes.

### When You Need the Actual Character Count

If you really need to count characters (not bytes):

```go
s := "世界"
count := len([]rune(s))  // 2 (correct!)
```

This converts the string to a list of characters first, then counts them.

### Simple Rule

- Use `range` when looping through a string (it handles characters correctly)
- Remember that `len(string)` counts bytes, not characters
- If you need the real character count, use `len([]rune(s))`

---

## Quick Cheat Sheet

**Converting types:**
```go
float64(x)           // whole number → decimal
int(y)               // decimal → whole number (cuts off decimals!)
strconv.Itoa(num)    // number → text
[]rune(s)            // string → list of characters
```

**Making variables:**
```go
name := "Alice"       // := when you know the value
var count int         // var when you want the zero value (0)
var price float64     // var when you need a specific type
```

**Important things to remember:**
- Go is strict: you can't mix types without converting
- `len(string)` counts bytes, not characters (use `[]rune` for character count)
- Use `range` when looping through strings
- Variables always start with safe zero values (0, false, "")
- When you see something confusing, think: "What type is this?"