# Go_linked_list

## 1. Create a new struct called StackSlice that implements Stacker, and uses an int slice as its internal representation.

## 2. Create a new struct called StackLinked that implements Stacker, and uses a singly (or doubly) linked list as its internal representation.

## 3. Go’s fmt package has a useful interface called Stringer that looks like this:
```go
type Stringer interface {   // already defined in fmt
    String() string
}
```
Anything that implements Stringer can be printed with fmt print functions, such as Println and Printf.

StackSlice and StackLinked both implement the Stringer interface.

## 4. Put all of your Stacker testing into a single testing function with this header:
```go
// s is assumed to be empty
func stackerTest(s Stacker) {
    // ... Stacker test code goes here ...
}
```

## 5. Implement functions:
```go
// Pre-condition:
//    none
// Post-condition:
//    s.isEmpty()
func popAll(s Stacker) {
    // ...
}
```
```go
// Pre-condition:
//    none
// Post-condition:
//    returns true if s and t have the same elements in the same order;
//    both s and t have the same value after calling stackEquals as before
// Annoying constraint:
//    Use only Stackers in the body of this functions: don't use arrays,
//    slices, or any container other than a Stacker.
func stackEquals(s, t Stacker) bool {
    // ...
}
```

