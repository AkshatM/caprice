# About

A Go client for RANDOM.org. Generate truly random numbers in your code!

# Example

```
package main

import (
	"github.com/AkshatM/caprice"
	"fmt"
)

func main() {
	rng := caprice.TrueRNG(<api key>)

	// generate 10 integers in base 10 between 1 and 100 with no duplicates 
	intArray, err := rng.GenerateIntegers(10, 1, 100, false)
	fmt.Printf("%v, intArray)

	// same as last time, but get back the entirety of the JSON response as a Result struct
	result, err := rng.GenerateIntegersRaw(10, 1, 100, false)
	fmt.Printf("%+v", result)

	// same as last time, but get back a structure that contains all the information you need
	// to verify the data you received along with the data itself in convenient format.
	verifiableResult, err := rng.GenerateSignedIntegers(10, 1, 100, false)
	fmt.Printf("%v", verifiableResult.Data) // actual data
	fmt.Printf("%+v", verifiableResult)
}
```

# Installation

For now, 

```
go get -u github.com/AkshatM/caprice
```

# Documentation

Working on getting this into Godoc. For now, consult the source.

With the exception of `verifySignature`, all API calls are supported as listed [here](https://api.random.org/json-rpc/1/basic) and [here](https://api.random.org/json-rpc/1/signing).

- Per Go convention, all API calls begin with capitalised letters.
- Every basic API call `x` has a corresponding method called `xRaw` that will return a `Response` object. e.g. `GenerateIntegers` has `GenerateIntegersRaw`. This is useful if you need access to any of the other response items RANDOM.org returns. Signed methods already return the raw JSONified data as well as the actual data supplied, so no equivalent exists for signed methods.
- `verifySignature` currently has [issues](https://stackoverflow.com/questions/48052917/preserve-json-rawmessage-through-multiple-marshallings?noredirect=1#comment83078240_48052917) :( however, you can still verify the integrity of your data by taking the signature and raw fields of the result struct from a signed method manually.

# Road Map

## 1.0: 
- Add support for using `advisoryDelay`, making this code truly thread-safe. 
- Get documentation up. 
- Add mocked network calls for decoupled testing.
