# Go's Data Structures: Slices, Maps, and Structs

> **Note:** These notes build on basic Go knowledge. If you understand variables and types from Day 1, you're ready to go!

## Arrays vs Slices: What's the Difference?

This is one of the most important concepts in Go. Let's break it down simply.

### Arrays: The Fixed Box

Think of an **array** as a box with a fixed number of compartments. Once you create it, you can't add more compartments.

```go
var numbers [5]int  // A box with exactly 5 compartments
```

**Key things about arrays:**
- The size is fixed forever (can't grow or shrink)
- The size is part of the type (`[5]int` and `[10]int` are different types!)
- When you pass it to a function, Go makes a complete copy of everything

**Why this matters:** If you have an array with 1000 items and pass it to a function, Go copies all 1000 items. That's slow and uses lots of memory!

```go
func doSomething(arr [1000]int) {
    // Go just copied 1000 integers to get here
}
```

**When do people use arrays?** Almost never! Slices are better in most cases. You'll mostly see arrays in special situations where you need a fixed size.

### Slices: The Flexible List

A **slice** is like a sticky note pointing to an array. The sticky note says "look at these specific items in that array over there."

```go
numbers := []int{1, 2, 3}  // This is a slice (notice: no size in the brackets!)
```

**Key things about slices:**
- Can grow and shrink as needed
- Don't copy all the data when passed to functions (just pass the sticky note!)
- Much more flexible and commonly used

**Think of it like this:**
- **Array** = Owning a house. If you give it to someone, they get their own complete copy of the house (expensive!)
- **Slice** = Having the address to a house. You can give someone the address, and they visit the same house (cheap and easy!)

### How to Tell Them Apart

```go
// Array - has a number in the brackets
var arr [5]int

// Slice - brackets are empty
var slice []int
```

**Simple rule:** In Go, you'll almost always use slices, not arrays. If you see empty brackets `[]`, that's a slice - the flexible, useful one!

## Understanding Length vs Capacity

Every slice has **two numbers** that describe it. This is like a parking lot: one number tells you how many cars are currently parked, another tells you how many parking spots exist total.

### Length: How Many Items Are Actually There

**Length** is how many items you have right now in your slice. You can use and access all of these items.

```go
s := []int{10, 20, 30}
fmt.Println(len(s))  // 3 - there are 3 items
```

Think of this as: "How many items am I using right now?"

### Capacity: How Much Room Do I Have?

**Capacity** is how much space the underlying array has before it needs to get bigger.

```go
s := make([]int, 3, 5)  // I'm using 3, but have room for 5

fmt.Println(len(s))  // 3 - using 3 items
fmt.Println(cap(s))  // 5 - have room for 5 items

// You can access s[0], s[1], s[2]
// The array has 2 more empty spots ready to use
```

**Real-world analogy:** You have a bookshelf with 10 shelves (capacity), but only 4 books on it (length). You have room for 6 more books before you need a bigger bookshelf.

### Why Should You Care About Capacity?

**Performance!** When you add a new item to a slice:
- If there's room in the capacity → Fast! Just put it in an empty spot
- If capacity is full → Slow! Go has to build a bigger array and move everything

```go
s := make([]int, 0, 100)  // Start with room for 100

// Adding items is fast because there's plenty of room
for i := 0; i < 100; i++ {
    s = append(s, i)  // Fast - just using existing space
}
```

**Simple rule:** Go handles capacity automatically. Understanding it helps you write more efficient code.

## What Happens When You Run Out of Room?

Imagine you're adding items to your slice, and suddenly you've filled up all the capacity. What happens next?

### Go Automatically Gets You a Bigger Array

When you append to a slice and run out of room, Go does this automatically:

1. **Makes a bigger array** (usually doubles the size - if you had room for 4, now you get room for 8)
2. **Copies everything** from the old array to the new one
3. **Updates your slice** to point at the new, bigger array
4. **Throws away the old array** (Go's garbage collector cleans it up)

```go
s := make([]int, 0, 2)  // Room for 2 items
s = append(s, 1)        // [1] - still room
s = append(s, 2)        // [1, 2] - now full!
s = append(s, 3)        // [1, 2, 3] - Go made a bigger array!
```

**Think of it like this:** Your bookshelf is full, so Go goes and buys you a bigger bookshelf, moves all your books to it, and gets rid of the old one.

### The Most Important Rule: Always Reassign append

Here's the thing that trips up beginners: **You MUST capture the result of append**.

```go
s := []int{1, 2, 3}
s = append(s, 4)  // ✅ CORRECT: reassigning to s
```

**Why?** Because `append` might have created a new array somewhere else in memory. The old `s` could be pointing to the wrong place!

### Common Mistake to Avoid

```go
s := []int{1, 2, 3}
append(s, 4)  // ❌ WRONG: Not saving the result!

fmt.Println(s)  // Still [1, 2, 3] - the 4 disappeared!
```

This is like someone handing you a new address for your books, and you just throw the address away. Now you can't find your books!

**Simple rule:** Always write `s = append(s, ...)` Never just `append(s, ...)`

## Maps: Your Key-Value Phone Book

A **map** is like a phone book. You look up someone's name (the key) and get their phone number (the value).

```go
ages := make(map[string]int)  // Name → Age
ages["Alice"] = 25
ages["Bob"] = 30

fmt.Println(ages["Alice"])  // 25
```

**Think of it like this:** Instead of remembering positions (like "the 3rd item"), you remember names. Much more useful!

### How Do Maps Work Behind the Scenes?

Maps use something called a **hash table**. When you give Go a key:

1. **Go converts the key to a number** (this is called "hashing")
2. **That number points to a storage spot** (like a locker in a gym)
3. **Your value gets stored in that spot**
4. **Looking it up again is super fast** - just hash the key and check that spot!

```go
ages["Alice"] = 25
// Go: "Alice" → hash → number 42 → store 25 in spot 42

fmt.Println(ages["Alice"])
// Go: "Alice" → hash → number 42 → look in spot 42 → found 25!
```

### How Fast Are Maps?

**Really fast!** For most purposes, maps are instant:
- Looking up a value: Nearly instant
- Adding a new item: Nearly instant
- Deleting an item: Nearly instant

This is called **O(1) time complexity** in computer science - constant time regardless of map size.

### Important Things to Know About Maps

**Maps grow automatically:**
- Start small, add items, and Go makes the map bigger as needed
- You don't have to think about this - it just works!

**Maps aren't safe for concurrent use:**
- If multiple goroutines (think: parallel tasks) access the same map, weird bugs happen
- Solution: Use locks or `sync.Map` for concurrent access

## Structs vs Maps: Which One Should I Use?

Both structs and maps let you group related information together. But when do you use which?

### What's a Struct?

A **struct** is like a form with labeled blanks. You decide what the labels are when you write your code.

```go
type Person struct {
    Name string  // Everyone has a Name
    Age  int     // Everyone has an Age
}

alice := Person{Name: "Alice", Age: 25}
```

**Think of it like a job application:** The fields are printed on the form. Everyone fills out the same fields.

### What's a Map?

A **map** is flexible. You can add any keys you want, whenever you want.

```go
info := make(map[string]int)
info["height"] = 170      // Add height
info["weight"] = 65       // Add weight
info["shoe_size"] = 9     // Add whatever!
```

**Think of it like a notebook:** You can write whatever you want. The entries aren't predetermined.

### When to Use Structs

Use a **struct** when you know what fields you need:

```go
type Book struct {
    Title  string
    Author string
    Pages  int
}
```

**Reasons to use structs:**
- ✅ **Typo protection:** If you type `book.Auther`, Go says "that's not a field!"
- ✅ **Faster:** Structs are more efficient than maps
- ✅ **Better code:** Other programmers can see exactly what fields exist
- ✅ **Methods:** You can attach functions to structs

```go
func (b Book) Summary() string {
    return b.Title + " by " + b.Author
}
```

### When to Use Maps

Use a **map** when you don't know what the keys will be ahead of time:

```go
// User preferences loaded from a file
preferences := map[string]string{
    "theme": "dark",
    "language": "en",
    // Could be any settings!
}
```

**Reasons to use maps:**
- ✅ **Flexibility:** Add and remove keys as you go
- ✅ **Dynamic data:** Keys come from user input or files
- ✅ **Unknown structure:** Like parsing JSON with unknown fields

### Real Examples

**Use a struct for:**
```go
type User struct {
    Username string
    Email    string
    Age      int
}
// Every user has these exact fields
```

**Use a map for:**
```go
// Storing arbitrary settings from a config file
settings := make(map[string]string)
// Could contain anything!
```

**Simple rule:**
- Know the fields when writing code? → Use a **struct**
- Fields come from user/files/unknown? → Use a **map**

## Slice Sharing: The Tricky Part

This is probably the most confusing thing about slices.

### The Problem: Multiple Slices, Same Data

Remember how a slice is like a sticky note pointing to an array? Well, you can have **multiple sticky notes pointing to the same array**.

```go
original := []int{1, 2, 3, 4, 5}
slice1 := original[0:3]  // Points to items 0, 1, 2
slice2 := original[2:5]  // Points to items 2, 3, 4

fmt.Println(slice1)  // [1, 2, 3]
fmt.Println(slice2)  // [3, 4, 5]
```

So far, so good. But watch what happens when you modify `slice1`:

```go
slice1[2] = 99  // Change the 3rd item in slice1

fmt.Println(slice1)  // [1, 2, 99]
fmt.Println(slice2)  // [99, 4, 5] - WHOA! slice2 changed too!
```

**Why did this happen?** Because both slices point to the same underlying array. When you changed position 2 through `slice1`, you changed the actual array. And `slice2` is looking at that same array!

**Think of it like this:**
- You have one notebook (the array)
- You give Alice a bookmark to pages 1-3
- You give Bob a bookmark to pages 3-5
- Alice writes on page 3
- Bob sees the change because it's the same notebook!

### Real Danger: Memory Leaks

Here's a sneaky bug that can waste a ton of memory:

```go
func getFirstTen(data []byte) []byte {
    return data[0:10]  // Return just the first 10 bytes
}

hugeFile := readFile("huge-file.txt")  // This is 1GB!
small := getFirstTen(hugeFile)         // Just 10 bytes

// Problem: The entire 1GB file stays in memory!
// Even though you only want 10 bytes!
```

**Why?** Because `small` still points to the huge underlying array. As long as `small` exists, Go can't throw away that 1GB array.

**Think of it like this:** You photocopy one page from a giant encyclopedia, but you keep the entire encyclopedia on your bookshelf just to keep that one photocopied page valid. Waste of space!

### How to Fix This: Make a Real Copy

When you need independence, **make a real copy** of the data:

```go
original := []int{1, 2, 3, 4, 5}

// Make a new slice with its own array
copy1 := make([]int, 3)
copy(copy1, original[0:3])  // Copy the data over

copy1[2] = 99
fmt.Println(original)  // [1, 2, 3, 4, 5] - original unchanged!
fmt.Println(copy1)     // [1, 2, 99] - only copy1 changed
```

**What happened?** You created a brand new array and copied the values into it. Now they're completely independent.

**Think of it like this:** Instead of both people using bookmarks in the same notebook, you made a photocopy. Now each person has their own notebook.

### Advanced: The Three-Number Slice Trick

There's a more advanced technique using `s[i:j:k]` (three numbers instead of two):

```go
original := []int{1, 2, 3, 4, 5}
slice1 := original[0:3:3]  // [start:end:capacity]
```

This limits the capacity, which prevents `append` from accidentally writing to the shared array. But honestly, if this confuses you, just use the copy method above - it's clearer!

### When Should You Worry About This?

**Be careful when:**
- ⚠️ Returning slices from functions
- ⚠️ Taking a small piece of a huge slice
- ⚠️ Passing slices to other functions that might modify them

**You're safe when:**
- ✅ Creating new slices with `make()`
- ✅ Using `append()` correctly (it often creates new arrays automatically)
- ✅ Making explicit copies when needed

**Practical advice:** When in doubt, make a copy! It's a tiny bit slower, but it prevents confusing bugs.

---

## Quick Reference Cheat Sheet

### Working with Slices

**Creating slices:**
```go
s := []int{1, 2, 3}           // Create with values
s := make([]int, 5)           // Create with length 5, all zeros
s := make([]int, 3, 5)        // Length 3, capacity 5
```

**Using slices:**
```go
len(s)                        // How many items?
cap(s)                        // How much room total?
s = append(s, 4)              // Add item (always reassign!)
s[1:3]                        // Get items 1 and 2 (shares array)
```

**Copying slices (to avoid sharing):**
```go
newSlice := make([]int, len(oldSlice))
copy(newSlice, oldSlice)      // Now independent!
```

### Working with Maps

**Creating and using maps:**
```go
m := make(map[string]int)     // Create empty map
m["Alice"] = 25               // Add or update value
age := m["Alice"]             // Get value
age, exists := m["Alice"]     // Check if key exists
delete(m, "Alice")            // Remove key
```

**Looping through maps:**
```go
for key, value := range m {
    fmt.Println(key, value)
}
```

### Working with Structs

**Defining and using structs:**
```go
type Person struct {          // Define the structure
    Name string
    Age  int
}

p := Person{Name: "Bob", Age: 30}  // Create instance
fmt.Println(p.Name)                // Access field
```

### Key Things to Remember

**About slices:**
- ✅ Always write `s = append(s, ...)` - don't forget the `=`
- ⚠️ Sub-slices share the same array - make copies if you need independence
- 📊 `len()` = items you're using, `cap()` = total room available

**About maps:**
- 📖 Like a dictionary: look things up by key, not position
- 🚀 Super fast for lookups
- ⚠️ Not safe with goroutines without extra protection

**About structs:**
- 📋 Use when you know the fields ahead of time
- ✅ Type-safe and efficient
- 🎯 Better than maps when structure is known

**Choosing between them:**
- Known fields at compile time? → **Struct**
- Keys from user/files/runtime? → **Map**
- Ordered list of items? → **Slice**
