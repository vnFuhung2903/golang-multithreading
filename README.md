# Golang Multithreading

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
### For-loop
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

### Package, Import
Every Go program is made up of packages. Programs start running in package `main`. This program is using the packages with import paths `"fmt"` and `"math/rand"`. By convention, the package name is the same as the last element of the import path.
```
import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)
```
### Go module
A module is a collection of Go packages stored in a file tree with a `go.mod` file at its root. The `go.mod` file defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build. Each dependency requirement is written as a module path and a specific semantic version.
```
module go_btc

go 1.24

require (
	github.com/btcsuite/btcutil v1.0.2
	golang.org/x/crypto v0.36.0
)
```
In addition to `go.mod`, the go command maintains a file named `go.sum` containing the expected cryptographic hashes of the content of specific module versions.\
The go command uses the `go.sum` file to ensure that future downloads of these modules retrieve the same bits as the first download, to ensure the modules which the project depends on do not change unexpectedly, whether for malicious, accidental, or other reasons. Both `go.mod` and `go.sum` should be checked into version control.

### Naming convention
Names are as important in Go as in any other language. They even have semantic effect: the visibility of a name outside a package is determined by whether its first character is upper case. It's therefore worth spending a little time talking about naming conventions in Go programs.
#### Package name
By convention, packages are given lower case, single-word names; there should be no need for **_underscores** or **mixedCaps**.
#### Getter
Go doesn't provide automatic support for getters and setters. It's neither idiomatic nor necessary to put **Get** into the getter's name. The use of upper-case names for export provides the hook to discriminate the field from the method.
#### Interface names
By convention, one-method interfaces are named by the method name plus an `-er` suffix or similar modification to construct an agent noun: **Reader, Writer, Formatter, CloseNotifier** etc. There are a number of such names, and it's productive to honor them and the function names they capture. **Read, Write, Close, Flush, String** and so on have canonical signatures and meanings. Giving custom method one of those names will result in confusion unless it has the same signature and meaning. Conversely, if a custom type implements a method with the same meaning as a method on a well-known type, give it the same name and signature; call the string-converter method **String** not **ToString**.
#### MixedCaps
Finally, the convention in Go is to use **MixedCaps** or **mixedCaps** rather than underscores to write multiword names.

### Data structures
