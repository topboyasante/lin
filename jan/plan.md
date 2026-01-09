# 23-Day Go Mastery Plan (Flexible Schedule)

## Table of Contents

- [Day 0: Setup & Preparation](#day-0-setup--preparation)
- [Week 1: Foundations & Type System Deep Dive](#week-1-foundations--type-system-deep-dive)
  - [Day 1](#day-1)
  - [Day 2](#day-2)
  - [Day 3](#day-3)
  - [Day 4](#day-4)
  - [Day 5](#day-5)
  - [Day 6](#day-6)
  - [Day 7](#day-7)
- [Week 2: Advanced Patterns & Error Handling](#week-2-advanced-patterns--error-handling)
  - [Day 8](#day-8)
  - [Day 9](#day-9)
  - [Day 10](#day-10)
  - [Day 11](#day-11)
  - [Day 12](#day-12)
  - [Day 13](#day-13)
  - [Day 14](#day-14)
- [Week 3: Concurrency Mastery & Testing](#week-3-concurrency-mastery--testing)
  - [Day 15](#day-15)
  - [Day 16](#day-16)
  - [Day 17](#day-17)
  - [Day 18](#day-18)
  - [Day 19](#day-19)
  - [Day 20](#day-20)
  - [Day 21](#day-21)
  - [Day 22 - FINAL DAY](#day-22---final-day)
- [Resources](#resources-to-supplement-learning-go)
- [Success Metrics](#success-metrics)
- [Daily Habits](#daily-habits)

---

## Day 0: Setup & Preparation

### READ
By the end of this reading, can you answer:

- What are the main chapters in "Learning Go" and which ones have you already covered?
- What topics in the book are completely new to you?
- What's the philosophy behind Go's design (simplicity, concurrency, etc.)?

### CODE

- Create a GitHub repository named `go-mastery-jan-2026` (or your preferred name)
- Set up folder structure: `/week1`, `/week2`, `/week3`, `/week4`
- Install necessary tools: `go install honnef.co/go/tools/cmd/staticcheck@latest` and `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
- Create a `day0/main.go` with a simple "Hello World" program
- Write a basic test file `day0/main_test.go` with one test case
- Run `go test -v` to verify your setup works
- Test your schedule and commit everything to git

**Tip:** This setup day ensures you hit the ground running tomorrow with zero friction.

## Week 1: Foundations & Type System Deep Dive

### Day 1

#### READ: Ch 1-2
By the end, can you answer:

- What are Go's built-in types and how do they differ from other languages?
- What are the three ways to declare variables in Go?
- What is zero value and why does it matter?
- How do `:=` and `var` differ in usage?
- What are runes and how do they relate to strings?

#### CODE: Build a CLI tool for data type processing

- Create `week1/day1/datatool/main.go`
- Build a CLI that accepts JSON input (use `flag` package for CLI args)
- Parse the JSON and detect the type of each field (string, number, bool, array, object)
- Print a formatted report of all fields and their types
- Handle errors gracefully (invalid JSON, empty input)
- Test with at least 3 different JSON structures

**Tip:** Use `encoding/json` and `reflect` packages. This will force you to work with multiple types.

### Day 2

#### READ: Ch 3
By the end, can you answer:

- What's the difference between arrays and slices in Go?
- How does slice capacity differ from length?
- What happens when you append to a slice beyond its capacity?
- How are maps implemented internally and what are their performance characteristics?
- When should you use a struct vs a map?
- What is slice sharing and why can it cause bugs?

#### CODE: Implement custom data structures

- Create `week1/day2/datastructures/stack.go`
- Implement a Stack with `Push`, `Pop`, `Peek`, and `IsEmpty` methods using slices
- Create `queue.go` and implement a Queue with `Enqueue`, `Dequeue`, `Front`, `IsEmpty`
- Write a program in `main.go` that demonstrates both structures
- Use `fmt.Printf` with `%v` and `%#v` to inspect internal slice behavior
- Create benchmarks comparing different capacity initializations

**Tip:** Pay attention to how slices grow. Try initializing with `make([]int, 0, 10)` vs `make([]int, 0)` and observe the difference.

### Day 3

#### READ: Ch 4
By the end, can you answer:

- What is shadowing and how can it cause bugs?
- What are the scoping rules for variables in Go?
- How do for loops work differently in Go compared to other languages?
- What is the "comma ok" idiom?
- When should you use `break` vs `continue` vs `goto`?

#### CODE: Build a config file parser

- Create `week1/day3/configparser/main.go`
- Define a config file format (key=value pairs, support comments with `#`)
- Implement `ParseConfig(filename string) (map[string]string, error)`
- Handle various error cases: file not found, malformed lines, empty values
- Use proper control structures: range over lines, continue for comments, break on errors
- Write a main function that reads `config.txt` and prints parsed values
- Test with valid and invalid config files

**Tip:** Use `bufio.Scanner` to read line by line. Pay attention to variable shadowing in error handling.

### Day 4

#### READ: Ch 5
By the end, can you answer:

- What are closures and how do they capture variables?
- What is the order of execution for `defer` statements?
- When should you use `panic` vs returning an error?
- How does `recover` work and where must it be called?
- What are variadic functions and how do you use them?
- How do named return values work?

#### CODE: Create a middleware chain system

- Create `week1/day4/middleware/main.go`
- Define a Handler type: `type Handler func(string) string`
- Implement a `Chain` function that takes multiple handlers and returns a single composed handler
- Create 3 middleware functions: Logger (prints input), Uppercase (transforms to uppercase), Trimmer (removes spaces)
- Use closures to wrap each middleware with additional behavior
- Demonstrate chaining all three together in main
- Add a `defer` statement to measure execution time

**Tip:** This mimics HTTP middleware patterns. Think about how each function wraps the next.

### Day 5

#### READ: Ch 6
By the end, can you answer:

- When should you use pointers vs values?
- What is the zero value of a pointer?
- How does Go prevent pointer arithmetic?
- What happens when you pass a pointer to a function?
- What's the difference between `*T` method receivers and `T` method receivers?
- How does garbage collection work with pointers?

#### CODE: Build a memory cache with pointer analysis

- Create `week1/day5/cache/cache.go`
- Define a Cache struct with a `map[string]*CacheEntry` field
- Implement `Set(key string, value interface{})` and `Get(key string) (interface{}, bool)`
- Create a `CacheEntry` struct that stores value and expiration time
- Implement eviction logic for expired entries
- Write benchmarks comparing pointer vs value storage for large structs
- Use `testing.B.ReportAllocs()` to see allocation differences

**Tip:** Run benchmarks with `go test -bench=. -benchmem`. Watch how pointer usage affects allocations.

### Day 6

#### READ: Ch 7
By the end, can you answer:

- What's the difference between a type declaration and a type alias?
- How do methods work in Go?
- What is an interface and how is it different from other languages?
- What does it mean that interfaces are satisfied implicitly?
- What is the empty interface `interface{}` and when should you use it?
- How do type assertions and type switches work?

#### CODE: Design a plugin system using interfaces

- Create `week1/day6/plugins/main.go`
- Define a Plugin interface with methods: `Name() string`, `Execute(input string) (string, error)`
- Implement three plugins: ReversePlugin, Base64Plugin, HashPlugin
- Create a PluginManager that registers and runs plugins by name
- Implement a `RunAll` method that executes all registered plugins in sequence
- Write tests that verify each plugin satisfies the interface
- Use type assertions to add optional `Validate()` method to some plugins

**Tip:** This demonstrates interface polymorphism. Notice how you never explicitly say a type implements an interface.

### Day 7

#### READ: Review Week 1 + Go Memory Model
By the end, can you answer:

- What are the key concepts you learned this week?
- How does Go's garbage collector work at a high level?
- What is escape analysis and why does it matter?
- What happens to variables that escape to the heap?
- Which areas felt weakest this week?

#### CODE: Write comprehensive tests

- Pick your best project from this week (probably the plugin system or cache)
- Write table-driven tests for all major functions
- Create test cases for: happy path, error cases, edge cases (empty input, nil, etc.)
- Add test helpers to reduce duplication
- Run tests with coverage: `go test -cover`
- Aim for >80% coverage
- Run with race detector: `go test -race`

**Tip:** Table-driven tests format:
```go
tests := []struct {
    name string
    input string
    want string
}{
    {"case1", "input1", "output1"},
}
```

## Week 2: Advanced Patterns & Error Handling

### Day 8

#### READ: Ch 8
By the end, can you answer:

- What's the difference between `errors.New` and `fmt.Errorf`?
- How do you wrap errors with `%w` and why is it useful?
- What is `errors.Is` and `errors.As`?
- When should you create custom error types?
- What's the best practice for error messages (capitalization, punctuation)?
- When is it appropriate to use `panic`?

#### CODE: Build a retry mechanism

- Create `week2/day8/retry/retry.go`
- Implement `Retry(fn func() error, maxAttempts int, delay time.Duration) error`
- Create custom error types: `MaxRetriesError`, `TemporaryError`, `PermanentError`
- Use `errors.Is` to check if an error is temporary and should be retried
- Wrap errors at each retry with context: "attempt 2/5: original error"
- Write a demo in `main.go` that simulates flaky operations
- Add tests that verify retry behavior with different error types

**Tip:** Use `time.Sleep` between retries. Think about exponential backoff as a bonus.

### Day 9

#### READ: Ch 9
By the end, can you answer:

- What is a Go module and how does it differ from a package?
- What is semantic versioning and how does Go use it?
- What's the purpose of `go.mod` and `go.sum`?
- How do you organize code into multiple packages?
- What are the rules for exported vs unexported identifiers?
- How do internal packages work?

#### CODE: Create a multi-package library

- Create `week2/day9/mathlib/` as your root module
- Run `go mod init github.com/yourusername/mathlib`
- Create packages: `statistics/`, `geometry/`, `algebra/`
- Implement 2-3 functions in each package (e.g., Mean, Median in statistics)
- Create a `cmd/demo/main.go` that imports and uses all packages
- Add package-level documentation comments
- Use `go doc` to verify your documentation

**Tip:** Each package should have a clear, single purpose. Notice how import paths work with your module name.

### Day 10

#### READ: Ch 10 (Part 1)
By the end, can you answer:

- What is a goroutine and how is it different from a thread?
- How do you launch a goroutine?
- What are channels and how do they enable communication?
- What's the difference between buffered and unbuffered channels?
- What happens when you send to or receive from a channel?
- How does `close` work on channels?

#### CODE: Build a worker pool

- Create `week2/day10/workerpool/main.go`
- Define a `Job` struct with an ID and some work (e.g., processing a number)
- Create a jobs channel and results channel
- Launch N worker goroutines that read from jobs and write to results
- Send 100 jobs into the jobs channel
- Collect all results from results channel
- Properly close channels and wait for all workers to finish

**Tip:** Use a `WaitGroup` to wait for workers (you'll learn this formally later, but try it). Think about when to close channels.

### Day 11

#### READ: Ch 10 (Part 2)
By the end, can you answer:

- What is the `select` statement and how does it work?
- How do you use `select` with a `default` case?
- What are channel directions (`chan<-` vs `<-chan`) and why are they useful?
- How do you implement timeouts with channels?
- What is the pattern for ranging over a channel?
- What happens if you forget to close a channel being ranged over?

#### CODE: Implement a pub/sub system

- Create `week2/day11/pubsub/pubsub.go`
- Define a `PubSub` struct with methods: `Subscribe(topic string) <-chan Message` and `Publish(topic string, msg Message)`
- Support multiple subscribers per topic
- Use `select` to handle publishing to multiple subscribers without blocking
- Implement `Unsubscribe` functionality
- Create a demo with 3 topics and 5 subscribers in `main.go`
- Add graceful shutdown that closes all channels

**Tip:** Store subscribers in a `map[string][]chan Message`. Use buffered channels to prevent blocking.

### Day 12

#### READ: Ch 11 (Standard Library)
By the end, can you answer:

- What are the key interfaces in the `io` package?
- How do `io.Reader` and `io.Writer` work?
- What's the purpose of `time.Time` vs `time.Duration`?
- How do you parse and format JSON in Go?
- What are the basics of `net/http` - handlers and servers?
- How do you handle different HTTP methods?

#### CODE: Build a simple HTTP server

- Create `week2/day12/httpserver/main.go`
- Implement a REST API with routes: `GET /users`, `POST /users`, `GET /users/:id`
- Use `http.HandleFunc` or `http.NewServeMux`
- Store users in memory (map or slice)
- Parse JSON request bodies and return JSON responses
- Add proper HTTP status codes (200, 201, 404, 400)
- Test with curl or Postman

**Tip:** Use `json.NewDecoder(r.Body).Decode(&user)` for requests and `json.NewEncoder(w).Encode(user)` for responses.

### Day 13

#### READ: Context Package Deep Dive
By the end, can you answer:

- What problem does the `context` package solve?
- What are the main functions: `Background`, `TODO`, `WithCancel`, `WithTimeout`, `WithDeadline`?
- How do you propagate context through a call chain?
- What is `context.Value` and when should you use it?
- How do you listen for context cancellation?
- What are best practices for context usage?

#### CODE: Add context to your HTTP server

- Modify `week2/day12/httpserver/main.go`
- Add a long-running endpoint `GET /process` that simulates 30 seconds of work
- Use context to make it cancellable if client disconnects
- Add request timeouts using `http.TimeoutHandler`
- Implement a middleware that adds request ID to context
- Pass context through all function calls
- Test cancellation by starting a request and cancelling it (Ctrl+C in curl)

**Tip:** Use `ctx := r.Context()` to get request context. Check `ctx.Done()` in long-running operations.

### Day 14

#### READ: Ch 12 (Generics)
By the end, can you answer:

- What are type parameters in Go?
- How do you define a generic function?
- What are type constraints?
- What is the `any` constraint?
- How do you use the `comparable` constraint?
- When should you use generics vs interfaces?

#### CODE: Refactor data structures with generics

- Go back to `week1/day2/datastructures/`
- Create `week2/day14/generics/stack.go`
- Rewrite Stack as `Stack[T any]` with generic methods
- Rewrite Queue as `Queue[T any]`
- Create a `main.go` that demonstrates Stack and Queue with different types: int, string, custom structs
- Write tests that verify type safety
- Compare the generic version with the original - note the differences

**Tip:** Generic syntax: `func Push[T any](s *Stack[T], value T) { ... }`. Much cleaner than using `interface{}`.

## Week 3: Concurrency Mastery & Testing

### Day 15

#### READ: Ch 13 (sync package)
By the end, can you answer:

- What is a `Mutex` and when do you need it?
- What's the difference between `Mutex` and `RWMutex`?
- How does `sync.WaitGroup` work?
- What is `sync.Once` used for?
- What are atomic operations and when should you use `sync/atomic`?
- What is a race condition and how do you detect it?

#### CODE: Build a thread-safe cache

- Create `week3/day15/safecache/cache.go`
- Define `SafeCache` struct with `sync.RWMutex` and a map
- Implement thread-safe `Set`, `Get`, `Delete` methods
- Write a benchmark that compares `Mutex` vs `RWMutex` with 90% reads, 10% writes
- Create a concurrent test with 100 goroutines reading and writing
- Run with race detector: `go test -race`
- Add metrics: track hits, misses, total operations using `sync/atomic`

**Tip:** Use `RLock`/`RUnlock` for reads, `Lock`/`Unlock` for writes. Run benchmark: `go test -bench=. -benchmem`.

### Day 16

#### READ: Concurrency Patterns Research
By the end, can you answer:

- What is the pipeline pattern and how does it work?
- What is fan-out/fan-in?
- How do you implement rate limiting with channels?
- What is the worker pool pattern?
- How do you handle errors in concurrent pipelines?
- What is the done channel pattern?

#### CODE: Implement a data processing pipeline

- Create `week3/day16/pipeline/main.go`
- Stage 1: Generate numbers (1-1000) into a channel
- Stage 2: Square each number (fan-out with 3 workers)
- Stage 3: Filter even numbers only
- Stage 4: Sum all numbers (fan-in)
- Connect all stages with channels
- Add a done channel to handle cancellation
- Measure total processing time

**Tip:** Each stage is a function that takes input channel(s) and returns output channel. Chain them together in main.

### Day 17

#### READ: Ch 14 (Testing)
By the end, can you answer:

- What is table-driven testing and why is it useful?
- How do you structure test files and functions?
- What are subtests and how do you use `t.Run`?
- How do you measure test coverage?
- What are benchmarks and how do you write them?
- How do you use `testing.B` methods like `ResetTimer` and `ReportAllocs`?

#### CODE: Write comprehensive tests for pipeline

- Create `week3/day16/pipeline/pipeline_test.go`
- Write table-driven tests for each pipeline stage
- Use `t.Run` to create subtests for different scenarios
- Test edge cases: empty input, single item, large datasets
- Write benchmarks for the entire pipeline
- Run tests with coverage: `go test -coverprofile=coverage.out`
- View coverage: `go tool cover -html=coverage.out`
- Run with race detector: `go test -race`

**Tip:** Aim for 80%+ coverage. Don't forget to test error paths.

### Day 18

#### READ: Testing - Mocking & Integration
By the end, can you answer:

- Why do we need mocks in testing?
- How do you create a mock implementation of an interface?
- What is dependency injection and why is it important for testing?
- What's the difference between unit and integration tests?
- How do you use test helpers to reduce duplication?
- What is `t.Helper()` and when should you use it?

#### CODE: Build and test a mock HTTP client

- Create `week3/day18/httpclient/client.go`
- Define an `HTTPClient` interface with `Get(url string) ([]byte, error)`
- Implement `RealHTTPClient` using `net/http`
- Create `MockHTTPClient` that returns predefined responses
- Build a `UserService` that depends on `HTTPClient`
- Write tests using the mock - test success and error scenarios
- Add integration test that uses real HTTP (mark with build tag)

**Tip:** Use dependency injection: `func NewUserService(client HTTPClient) *UserService`. This makes testing easy.

### Day 19

#### READ: Ch 15 (Tooling - profiling)
By the end, can you answer:

- What is `go vet` and what does it check for?
- What is `staticcheck` and how is it different from `go vet`?
- How do you use `pprof` to profile CPU and memory?
- What are the different types of profiles (CPU, memory, goroutine, block)?
- How do you interpret a CPU profile?
- What is `go tool trace`?

#### CODE: Profile and optimize a project

- Pick your most complex project (probably the pipeline or pub/sub)
- Add CPU profiling: import `runtime/pprof`, wrap main logic
- Run and generate CPU profile: `go run main.go`
- Analyze with `go tool pprof cpu.prof` and look at `top10`, `list functionName`
- Add memory profiling and check for leaks
- Run `go vet` and `staticcheck` on all your code
- Fix any issues found and re-profile to see improvements

**Tip:** Add this to enable profiling:
```go
f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

### Day 20

#### READ: Go Runtime Internals
By the end, can you answer:

- How does the Go scheduler work (M:N threading)?
- What are G, M, and P in the runtime?
- How does the garbage collector work at a high level?
- What triggers a GC cycle?
- How can you tune GC with `GOGC`?
- What can you observe with `GODEBUG`?

#### CODE: Explore GC behavior

- Create `week3/day20/gcexplore/main.go`
- Write a program that allocates memory in a loop (create large slices)
- Add `runtime.ReadMemStats` to track allocations
- Print heap stats every 1000 iterations
- Run with `GODEBUG=gctrace=1` to see GC output
- Experiment with `GOGC` values (50, 100, 200)
- Create a visualization of heap growth over time

**Tip:** `runtime.MemStats` gives you detailed memory info. Watch how GC kicks in when heap grows.

### Day 21

#### READ: Ch 16 (Reflection & Unsafe)
By the end, can you answer:

- What is reflection and when should you use it?
- What are the key types in the `reflect` package?
- How do you inspect struct fields at runtime?
- What is the `unsafe` package and why is it dangerous?
- When is it acceptable to use `unsafe`?
- How does JSON encoding/decoding use reflection?

#### CODE: Build a simple serialization library

- Create `week3/day21/serialize/serialize.go`
- Implement `Marshal(v interface{}) ([]byte, error)` using reflection
- Support basic types: int, string, bool, slices, structs
- Use struct tags to control field names: `json:"fieldname"`
- Implement `Unmarshal(data []byte, v interface{}) error`
- Write tests with various struct types
- Compare your implementation with `encoding/json`

**Tip:** Use `reflect.TypeOf` and `reflect.ValueOf`. This is complex - focus on understanding, not perfection.

### Day 22 - FINAL DAY

#### READ: Review & Reflect
By the end, can you answer:

- What were the three most important concepts you learned?
- Which topics still feel unclear or need more practice?
- Can you explain goroutines, channels, and sync primitives confidently?
- Do you understand when to use pointers vs values?
- Are you comfortable writing tests and benchmarks?
- What low-level project excites you most for next month?

#### CODE: Mini Capstone - TCP Proxy

- Create `week3/day22/tcpproxy/main.go`
- Accept CLI args: local port and remote address
- Listen on local port for incoming connections
- For each connection, dial the remote address
- Copy data bidirectionally using goroutines
- Use context for graceful shutdown on SIGINT
- Add proper error handling and logging
- Write tests using `net` package test helpers
- Add basic stats: connections handled, bytes transferred

**Tip:** Use `io.Copy` for data transfer. This combines: networking, concurrency, context, error handling, testing.

---

## Resources to Supplement "Learning Go"

- [Official Go Blog](https://go.dev/blog/)
- [Go by Example](https://gobyexample.com)
- [Effective Go](https://go.dev/doc/effective_go)
- Go's race detector - use `go test -race` regularly
- [Dave Cheney's Blog](https://dave.cheney.net) (advanced topics)

---

## Success Metrics

By the end, you should confidently answer YES to:

- ✅ Can I explain how goroutines and channels work?
- ✅ Can I write table-driven tests with good coverage?
- ✅ Do I understand when to use pointers vs values?
- ✅ Can I use context for cancellation and timeouts?
- ✅ Can I profile and optimize Go code?
- ✅ Do I feel ready to build a reverse proxy or database?

---

## Daily Habits

- **Night before:** Review tomorrow's questions so your brain primes overnight
- **Session start:** Setup ready, distractions away, book/code open
- **After session:** Commit your code to GitHub with a meaningful message
- **Track it:** Check off each day - momentum is everything

