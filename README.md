# Golang Multithreading
## Table of content
- [Formation history](#formation-history)
- [Comparison Go vs other languages](#comparison-go-vs-java-javascript-python-and-c)
- [Use cases in programming applications](#how-would-go-be-used-to-programing-applications)
- [Function](#function)
- [For-loop](#for-loop)
- [Package/Import](#package-import)
- [Go Module](#go-module)
- [Naming Convention](#naming-convention)
- [Data Structure](#data-structures)
- [Multithreading](#multithreading)
- [Common packages](#common-packages)

## Formation history
Golang (or Go) was created by Google engineers Robert Griesemer, Rob Pike, and Ken Thompson in 2007 to address challenges in software development, such as slow compilation, dependency management, and concurrency. Officially released as an open-source language in 2009, Go was designed to be simple, efficient, and highly scalable, making it well-suited for cloud computing, networking, and system programming. Its syntax draws inspiration from C, but with modern features like garbage collection, built-in concurrency via goroutines, and a robust standard library.

## Comparison: Go vs. Java, JavaScript, Python, and C++
### 1. Performance
- **Go vs. C++**: Go is generally slower than C++ because it has garbage collection, whereas C++ provides manual memory management for optimized performance.
- **Go vs. Python**: Go is significantly faster than Python due to its compiled nature, whereas Python is interpreted and dynamically typed.
- **Go vs. Java**: Both are compiled languages, but Go's compilation process is faster. However, Java's JVM optimizations can sometimes result in better long-term performance.
- **Go vs. JavaScript**: JavaScript is interpreted and runs in a browser or Node.js, making it slower than Go for backend applications.

### 2. Concurrency
- **Go vs. Java**: Java uses threads with higher overhead, while Go's goroutines are lightweight and more efficient for concurrent programming.
- **Go vs. Python**: Python's Global Interpreter Lock (GIL) limits true parallel execution, whereas Go handles concurrency natively using goroutines.
- **Go vs. JavaScript**: JavaScript relies on an event loop and async programming, while Go provides native goroutines and channels for concurrency.
- **Go vs. C++**: C++ offers low-level thread management with high control but requires manual handling of concurrency, whereas Go simplifies concurrency with built-in support.

### 3. Simplicity & Syntax
- **Go vs. Java**: Java has verbose syntax with OOP principles, whereas Go is more minimalistic with a focus on readability and simplicity.
- **Go vs. Python**: Python is known for its simplicity, but Go enforces strict typing and a more structured approach.
- **Go vs. JavaScript**: JavaScript is flexible but has quirks and inconsistencies, while Go enforces strict rules to reduce errors.
- **Go vs. C++**: C++ has a complex syntax with manual memory management, whereas Go simplifies development with garbage collection and straightforward syntax.

### 4. Memory Management
- **Go vs. C++**: Go has automatic garbage collection, while C++ requires manual memory management, offering more control but with potential risks.
- **Go vs. Java**: Both use garbage collection, but Java's GC is more mature and optimized for larger applications.
- **Go vs. Python**: Both have garbage collection, but Python's reference counting can lead to performance issues in some cases.
- **Go vs. JavaScript**: JavaScript also uses garbage collection, but Go is optimized for backend performance.

### 5. Ecosystem & Use Cases
- **Go vs. Python**: Python excels in data science, machine learning, and scripting, while Go is better for backend systems, cloud computing, and networking.
- **Go vs. JavaScript**: JavaScript dominates web development, whereas Go is mainly used for backend services and infrastructure tools.
- **Go vs. Java**: Java is widely used in enterprise applications, Android development, and large-scale systems, while Go is preferred for microservices and cloud-native applications.
- **Go vs. C++**: C++ is used in system programming, game development, and performance-critical applications, whereas Go is preferred for scalable, cloud-based solutions.  

## How would Go be used to programing applications
- **Web Development**: Go is widely used in backend development, with frameworks like Gin and Fiber making it easy to build RESTful APIs and web services.
- **Cloud Computing**: Go is the preferred language for cloud-native applications, used in Kubernetes, Docker, and Terraform.
- **Networking & Distributed Systems**: With its strong concurrency model, Go is ideal for high-performance networking applications and microservices.
- **Command-Line Tools**: Many CLI applications, including DevOps tools, are built in Go due to its efficiency and ease of deployment.
- **Game Development**: Although not as popular as C++, Go is used in game development through libraries like Ebiten and Pixel.

## Function
### Variadic function
A variadic function is a function that accepts zero, one, or more values as a single argument.
```
func (blockchain *Blockchain) MineBlock(txs ...*Transaction) *Block {
	for _, tx := range txs {
		if blockchain.checkTransactionExists(tx) {
			panic("ERROR: Transaction existed")
		}
	}
	newBlock := NewBlock(txs, blockchain.blocks[len(blockchain.blocks)-1])
	return newBlock
}
```

### Multiple return values
A single return statement of a function can return multiple values. Using blank identifier `_` to use only a subset of the returned value.
```
func (blockchain *Blockchain) FindSpendableUTXO(address string) (int, []int) {
	UTXOs := blockchain.findWalletTXO(address)
	var spendableUTXO []int
	res := 0

	for _, unspentTXO := range UTXOs {
		res += unspentTXO
		spendableUTXO = append(spendableUTXO, unspentTXO)
	}
	return res, spendableUTXO
}
```

### Error/Exception handling
Errors in Go can be communicated via an explicit, separate return value. By convention, errors are the last return value and have type error, a built-in interface.
- `errors`: New constructs a basic error value with the given error message.
```
func (blockchain *Blockchain) AddBlock(newBlock *Block) error {
	var err error
	if newBlock == nil || newBlock.Height < blockchain.blocks[len(blockchain.blocks)-1].Height {
		err = errors.New("ERROR: New block height is too low")
	}
	blockchain.blocks = append(blockchain.blocks, newBlock)
	return err
}
```
- A `nil` value in the error position indicates that there was no error.
- A sentinel error is a predeclared variable that is used to signify a specific error condition.
```
func (wallet *Wallet) Address() string {
	publicKeyHash, err := wallet.HashPublicKey()
	if err != nil {
		fmt.Println(err)
	}
	version := byte(0x00)
	check := checksum(publicKeyHash)

	address := append([]byte{version}, publicKeyHash...)
	address = append(address, check...)
	return base58.Encode(address)
}
```
A panic typically means something went unexpectedly wrong. Mostly it's used to fail fast on errors that shouldn’t occur during normal operation, or an error that isn't handled gracefully.
```
func NewTransaction(from *Wallet, to string, amount int, blockchain *Blockchain) *Transaction {
	var TXIs []*TransactionInput
	var TXOs []*TransactionOutput
	fromAddress := from.Address()
	accumulated, spendableUTXO := blockchain.FindSpendableUTXO(fromAddress)
	if accumulated < amount {
		panic("ERROR: Not enough funds")
	}
	...
}
```
## For-loop
**for** is Go’s only looping construct.
```
func HashAllTransactions(txs []*Transaction) [32]byte {
	var blockData []byte
	for _, tx := range txs {
		blockData = append(blockData, tx.Hash[:]...)
	}
	return sha256.Sum256(blockData)
}
```
The `range` form of the for loop iterates over a slice or map. When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index. 

## Package, Import
Every Go program is made up of packages. Programs start running in package `main`. By convention, the package name is the same as the last element of the import path.
```
import (
	"crypto/sha256"
	"fmt"
	"gobtc/function"
	"gobtc/multithreading"
)
```
## Go module
A module is a collection of Go packages stored in a file tree with a `go.mod` file at its root. The `go.mod` file defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build. Each dependency requirement is written as a module path and a specific semantic version.
```
module gobtc

go 1.24

require (
	github.com/btcsuite/btcutil v1.0.2
	golang.org/x/crypto v0.36.0
)
```
In addition to `go.mod`, the go command maintains a file named `go.sum` containing the expected cryptographic hashes of the content of specific module versions.\
The go command uses the `go.sum` file to ensure that future downloads of these modules retrieve the same bits as the first download, to ensure the modules which the project depends on do not change unexpectedly, whether for malicious, accidental, or other reasons. Both `go.mod` and `go.sum` should be checked into version control.

## Naming convention
Names are as important in Go as in any other language. They even have semantic effect: the visibility of a name outside a package is determined by whether its first character is upper case. It's therefore worth spending a little time talking about naming conventions in Go programs.
### Package name
By convention, packages are given lower case, single-word names; there should be no need for **_underscores** or **mixedCaps**.

### Getter
Go doesn't provide automatic support for getters and setters. It's neither idiomatic nor necessary to put **Get** into the getter's name. The use of upper-case names for export provides the hook to discriminate the field from the method.

### Interface names
By convention, one-method interfaces are named by the method name plus an `-er` suffix or similar modification to construct an agent noun: **Reader, Writer, Formatter, CloseNotifier** etc. There are a number of such names, and it's productive to honor them and the function names they capture. **Read, Write, Close, Flush, String** and so on have canonical signatures and meanings. Giving custom method one of those names will result in confusion unless it has the same signature and meaning. Conversely, if a custom type implements a method with the same meaning as a method on a well-known type, give it the same name and signature; call the string-converter method **String** not **ToString**.

### MixedCaps
Finally, the convention in Go is to use **MixedCaps** or **mixedCaps** rather than underscores to write multiword names.

## Data structures
### Boolean
`bool` values are `true` and `false`

### Numeric
- Integers: `int`, `int8`, `int16`, `int32`, `int64`
- Unsigned integers: `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- Floating points: `float32`, `float64`
- Complex numbers: `complex64`, `complex128`
- Byte: `byte` is an alias for `uint8`
- Rune: `rune` is an alias for `int32`
```
version := byte(0x00)
const decimal int = 100000000
```
### String
A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. The number of bytes is called the length of the string and is never negative. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is `string`.\
The length of a string s can be discovered using the built-in function `len`. The length is a compile-time constant if the string is a constant. A string's bytes can be accessed by integer indices 0 through `len`-1. It is illegal to take the address of such an element; if s[i] is the i-th byte of a string, &s[i] is invalid.

### Array
An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the array and is never negative.\
The length of an array must evaluate to a non-negative constant representable by a value of type `int`, and can be discovered using the built-in function len. The elements can be addressed by integer indices 0 through `len`-1. Array types are always one-dimensional but may be composed to form multi-dimensional types.\
An array type T may not have an element of type T, or of a type containing T as a component, if those containing types are only array or struct types.
```
genesisHash [32]byte
```

### Slice
A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array. The number of elements is called the length of the slice and is never negative. The value of an uninitialized slice is `nil`.\
The length of a slice s can be discovered by the built-in function `len`; unlike with arrays it may change during execution. The elements can be addressed by integer indices 0 through `len`-1. The slice index of a given element may be less than the index of the same element in the underlying array.\
A slice, once initialized, is always associated with an underlying array that holds its elements. A slice therefore shares storage with its array and with other slices of the same array; by contrast, distinct arrays always represent distinct storage.\
The array underlying a slice may extend past the end of the slice. The capacity is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by slicing a new one from the original slice. The capacity of a slice a can be discovered using the built-in function `cap`.\
A new, initialized slice value for a given element type T may be made using the built-in function `make`.
```
var blockData []byte
var spendableUTXO []int
```

### Interface
An interface type defines a type set. A variable of interface type can store a value of any type that is in the type set of the interface. Such a type is said to implement the interface. The value of an uninitialized variable of interface type is `nil`.\
Error: `error` is an interface type
```
var err error
```

### Map
A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type. The value of an uninitialized map is `nil`.\
The comparison operators `== `and `!=` must be fully defined for operands of the key type; thus the key type must not be a function, map, or slice. If the key type is an interface type, these comparison operators must be defined for the dynamic key values; failure will cause a run-time panic.\
The number of map elements is called its length. For a map m, it can be discovered using the built-in function `len` and may change during execution. Elements may be added during execution using assignments and retrieved with index expressions; they may be removed with the `delete` and `clear` built-in function.\
A new, empty map value is made using the built-in function `make`.
```
	spentUTXOs := make(map[[32]byte][]int)
```
The initial capacity does not bound its size: maps grow to accommodate the number of items stored in them, except `nil` maps. A `nil` map is equivalent to an empty map except that no elements may be added.

### Structures
A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type. This concept is generally compared with the classes in object-oriented programming. It can be termed as a lightweight class that does not support inheritance but supports composition.
The `type` keyword is used to declare a new structure type.
```
type Block struct {
	Timestamp int64
	Hash      [32]byte
	Data      []*Transaction
	PrevHash  [32]byte
	Height    int
}
```

## Multithreading
### Goroutines
Goroutine is a lightweight thread managed by the Go runtime. It is created using the `go` keyword before a function call:
```
go func() {
    defer wg.Done()
    _, err := http.Get(url)
    if err != nil {
        channel <- fmt.Sprintf("Error fetching %s: %v", url, err)
        return
    }
    channel <- fmt.Sprintf("Fetched %s succesfully", url)
}()
```

### Channel
Channels are communication mechanisms that allow goroutines to exchange data safely. Channels act as FIFO queues, so they ensure synchronization between concurrent tasks.\
A single channel may be used in send statements, receive operations.\
A channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type. The value of an uninitialized channel is `nil`.\
The optional `<-` operator specifies the channel direction, send or receive. If a direction is given, the channel is directional, otherwise it is bidirectional. A channel may be constrained only to send or only to receive by assignment or explicit conversion. The `<-` operator associates with the leftmost chan possible.
```
channel <- fmt.Sprintf("Fetched %s succesfully", url)
```
A new, initialized channel value can be made using the built-in function `make`.
```
channel := make(chan string, len(urls))
```
The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives). A `nil` channel is never ready for communication.\
A channel may be closed with the built-in function `close`. The multi-valued assignment form of the receive operator reports whether a received value was sent before the channel was closed.
```
close(channel)
```
The length and capacity of a channel can be discovered by the built-in functions `cap` and `len` by any number of goroutines without further synchronization

### Buffered channel
By default, channels are unbuffered, meaning that they will only accept sends `chan <-`, and if there is a corresponding receive, `<- chan` ready to receive the sent value. Buffered channels accept a limited number of values without a corresponding receiver for those values.
```
channel := make(chan string, len(urls))
```

### WaitGroups
A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls `WaitGroup.Add` to set the number of goroutines to wait for. Then each of the goroutines runs and calls `WaitGroup.Done` when finished. At the same time, `WaitGroup.Wait` can be used to block until all goroutines have finished.\
A WaitGroup must not be copied after first use.
In the terminology of the Go memory model, a call to `WaitGroup.Done` “synchronizes before” the return of any `Wait` call that it unblocks.

## Common packages
- `fmt`: provides formatting functions for I/O, is widely used for debugging and logging.
- `time`: provides functions for handling time, sleeping, and parsing dates
- `net/http`: is used for building HTTP servers and clients in Go
- `sync`: provides synchronization primitives (such as WaitGroups)
- `crypto`: contains various cryptographic implementations
- `errors`: provides functions for defining, wrapping, and unwrapping errors, is used for creating and handling errors manually.

## Example code
- The **Multithreading** example codes are in https://github.com/vnFuhung2903/golang-multithreading/tree/master/multithreading
- The **Function** example codes, the **For-loop** example codes and the **Data Structure** example codes are in https://github.com/vnFuhung2903/golang-multithreading/tree/master/entities
- The **Package Import** example codes, and the **Naming Convention** example codes are in all `.go` files
- The **Go Module** example codes are in https://github.com/vnFuhung2903/golang-multithreading/blob/master/go.mod
- The **Common packages** example codes:
  * The `fmt`, `net/http`, `sync` package are in https://github.com/vnFuhung2903/golang-multithreading/tree/master/multithreading/fetch.go
  * The `time`, `errors` package are in https://github.com/vnFuhung2903/golang-multithreading/blob/master/entities/blockchain.go
  * The `crypto` package is in https://github.com/vnFuhung2903/golang-multithreading/blob/master/entities/wallet.go