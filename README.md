# basicLog

Another basic non professional loggig for personal daily use program development.

```go
package main

import (
	"errors"
	"fmt"

	l "github.com/ahmetozer/basiclog"
)

func main() {
	l.Init()
	err := errors.New("this is example error")

	l.Debug("message %v", err)
	l.Info("message %v", err)
	l.Error("message %v", err)
	// l.Fatal("message %v", err)

    // instead of if err!=nil { log.Printf(err)}
	l.ErrNil(l.Info, myTestFunctionOne())

	// Checking Mutiple errors
	_, err2 := myTestFunctionTwo()
	_, err3 := myTestFunctionTwo()
	l.ErrNil(l.Error, err2, err3)

	// Do the things,combine errors, if no error execute, otherwise print error
	if l.ErrNil(l.Info, myTestFunctionOne(), myTestFunctionOne()) {
		fmt.Printf("Hello world, no error")
	}


    // Do the things,combine errors, if no error execute, otherwise print error
	if !l.ErrNil(l.Info, myTestFunctionOne(), myTestFunctionOne()) {
		fmt.Printf("Hello world, there is an error and printed before this line")
	}

}

var a,b int
func myTestFunctionOne() error {
    a = a + 1
	return errors.New("myTestFunctionOne")
}



func myTestFunctionTwo() (string, error) {
	b = b + 1
	return "", fmt.Errorf("myTestFunctiontwo %v", b)
}
```


