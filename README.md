# go-intro

When startinga  go project
That is when you do a

```go
go mod init {module_name}
```

when importin go modules no need to create a class or anything in the module you are
importing from because! you can just import the function itestelf somethin like

```go
import "{module_name}/foo"
foo.SomeFunctinFromFoo()
```

## Building

<hr/>
When are done coding and want to build our code we run a

```go
go build {file_name}
```

when this is done a couple things happen

* Compliation
  * The source code is parsed, type checking is done and translated into machine code (complied)
* Dependency Resolution
  * resolves and fetches any dependencies your code relies on. Downloads them if needed
* Executable Output
  * if the build has a main funciton then a binary is created to execute
* SHared LIbraries
  * In some cases, when you're building a package that is meant to be used as a shared library (for example, a package that provides functionality to other programs), go build can generate a shared library file (e.g., a .so file on Linux) instead of an executable binary.

* Error Checking
  * go build performs error checking and reports any syntax or type errors it encounters during compilation. If there are errors in your code, it will not produce a binary.

## Testing

<hr/>
When you are testing go code you should write the test in the same directory of the code
so if we have
.\utils\something.go
the test for something.go should live in the same reictory with _test suffxing the file name so we would have
.\\utls\something_test.go
inside this we would have some tthing looks like this

```go
package utils_tests
//import to have this 
import (
    "testing"
    "myModule/utils"
)

func TestMyLibrary(t *testing.T){
    some_value = utils.foo()
    if some_value != "bar"{
        t.Fail()
    }   
}

```

in order to run JUST THIS TEST
run the command

```go
go test utils

```

in order to run ALL TEESTS in the project run

```go
go run ./...
```
