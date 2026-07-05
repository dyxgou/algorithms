1. (multi) Which of these cause a variable to escape to the heap?

A) Returning a pointer to a local variable

B) Storing a value in an interface

C) Passing a small struct by value to a function

D) Taking the address of a variable and passing it to a goroutine

E) Using a value receiver method that doesn't escape

1: A, B, D

2. (single) What does the Go GC primarily optimize for?

A) Minimizing heap size

B) Minimizing pause times (low latency)

C) Maximizing throughput at all costs

D) Eliminating allocations entirely

2: A

3. (multi) Which statements about slices are true?

A) append always allocates a new backing array

B) Slicing a slice shares the same backing array

C) cap() can be larger than len()

D) Two slices can share memory and modifying one can affect the other

E) The zero value of a slice is nil and is safe to read from (len 0)

3: B, C

4. (multi) Which are valid reasons an interface method call has runtime cost compared to a direct call?

A) Pointer dereferencing speed of the interface value

B) Dynamic dispatch through the itable

C) Inability to inline across interface boundaries (usually)

D) GC pressure from boxing values into interfaces

E) Interface comparisons being O(1) always

4: A, C, D

5. (single) What happens when you compare two interface values holding the same dynamic type but uncomparable underlying values (e.g., a slice)?

A) Returns false

B) Returns true

C) Panics at runtime

D) Compile error

5: A

6. (multi) Which of the following can cause a goroutine leak?

A) A goroutine blocked sending on an unbuffered channel with no receiver

B) Forgetting to call cancel() on a context derived from context.WithCancel

C) A select with a default case in a loop

D) A goroutine waiting on a sync.WaitGroup that never reaches zero

E) Using context.WithTimeout correctly with deferred cancel

6: A, B, D

7. (single) What's the difference between sync.Mutex and sync.RWMutex in terms of write starvation?

A) RWMutex can starve writers if there's continuous read traffic

B) Mutex can starve readers

C) They behave identically

D) RWMutex prevents starvation by design

7: A

8. (multi) Which statements about Go's memory model are correct?

A) Without synchronization, there's no guarantee one goroutine sees another's writes

B) Channel send/receive establishes a happens-before relationship

C) sync/atomic operations alone are sufficient to build complex lock-free data structures safely in all cases

D) sync.WaitGroup.Wait() happens-before the goroutines it waits on complete

E) Closing a channel happens-before a receive that returns because the channel is closed

8: A, B, E

9. (single) In PGO (Profile-Guided Optimization), what does the compiler primarily use the profile for?

A) Garbage collection tuning

B) Better inlining decisions based on hot call paths

C) Reducing binary size

D) Choosing GOMAXPROCS at runtime

9: B

10. (multi) Which are true about Go generics?

A) Type parameters can have method constraints via interfaces

B) Generics use full monomorphization like C++ templates, with zero runtime dispatch ever

C) The Go compiler may use GC shape stenciling, sharing code across types with identical underlying representations

D) You can't use generic functions with type inference in all cases — explicit instantiation is sometimes required

E) Generic methods (not just functions) are fully supported on generic types

10: A, C, D

11. (single) What's the main risk of using context.Context values (via WithValue) for passing data?

A) Performance overhead on every request

B) Loss of type safety and implicit/hidden dependencies

C) Context values aren't goroutine-safe

D) Contexts can't carry more than one value

11: B

12. (multi) Which benchmarking practices are correct for go test -bench?

A) Always reset the timer with b.ResetTimer() after setup that shouldn't be measured

B) Use b.ReportAllocs() to track allocations per operation

C) Run with -benchmem to see B/op and allocs/op

D) A single run is statistically sufficient; no need for -count

E) Avoid compiler dead-code elimination by assigning results to a package-level variable

12: A, B, C

13. (single) What's the primary cause of false sharing in concurrent Go code?

A) Two goroutines writing to different cache lines

B) Two goroutines writing to variables that land on the same CPU cache line

C) Using too many goroutines relative to GOMAXPROCS

D) Channel buffering being too small

13: B

14. (multi) Which statements about error handling idioms are correct in modern Go?

A) errors.Is checks for equality or matches via an Unwrap chain

B) errors.As is used to extract a specific error type from a chain

C) Wrapping with fmt.Errorf("%w", err) preserves the chain for errors.Is/As

D) Sentinel errors should always be compared with == instead of errors.Is

E) panic/recover is the idiomatic way to handle expected error conditions

14: A, B, C

15. (single) Why might a for range loop over a map produce different iteration orders on different runs?

A) Go maps preserve insertion order but shuffle on resize

B) Go intentionally randomizes map iteration order to prevent reliance on ordering

C) It's a bug-prone implementation detail that varies by platform

D) Maps are sorted by hash value, which changes per run

15: D
