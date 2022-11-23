# Advance Golang

## FreeOSMemory
Go is a garbage collected language, which means that memory allocated and used by variables is automatically freed by the garbage collector when those variables become unreachable (if you have another pointer to the variable, that still counts as "reachable").

Freed memory does not mean it is returned to the OS. Freed memory means the memory can be reclaimed, reused for another variable if there is a need. So from the operating system you won't see memory decreasing right away just because some variable became unreachable and the garbage collector detected this and freed memory used by it.

The Go runtime will however return memory to the OS if it is not used for some time (which is usually around 5 minutes). If the memory usage increases during this period (and optionally shrinks again), the memory will most likely not be returned to the OS.

If you wait some time and not allocate memory again, freed memory will be returned to the OS eventually (obviously not all, but unused "big chunks" will be). If you can't wait for this to happen, you may call debug.FreeOSMemory() to force this behavior:

### How long should i wait to check how much memory is freed?
you should wait at least 7 minutes to check how much memory is freed. Sometimes it needs two GC passes, so it would be 9 minutes.
If that is not working, or it is too much time, you can add a periodic call to FreeOSMemory (no need to call runtime.GC() before, it is done by debug.FreeOSMemory() )

```go
package main

import (
    "runtime/debug"
    "time"
)

func main() {

    go periodicFree(1 * time.Minute)
    // Your program goes here
}

func periodicFree(d time.Duration) {
    tick := time.Tick(d)
    for _ = range tick {
        debug.FreeOSMemory()
    }
}

```

* [FreeOSMemory](https://pkg.go.dev/runtime/debug#FreeOSMemory) forces a garbage collection followed by an attempt to return as much memory to the operating system as possible.
* Keep in mind that it won't do anything unless the GC has been run
* The amount of memory released shown by runtime.ReadMemStats
 
## Topics
* [Interface](https://github.com/mateors/golang/tree/interface)
* Pointer
* Singleflight | Mutex | Sync
* Bitwise operator | Left shift | right shift
* Goroutine
* Channel | Select
* Race condition | [Race tracer](https://pkg.go.dev/runtime/trace#Start)
* Context

> Package singleflight provides a duplicate function call suppression

> go doc & go tool usage details

## How do we call function by its name (string)

```golang

package main

import "fmt"

func someFunction1(a, b int) int {
        return a + b
}

func someFunction2(a, b int) int {
        return a - b
}

func someOtherFunction(a, b int, f func(int, int) int) int {
        return f(a, b)
}

func main() {
      
      //Method -1
        fmt.Println(someOtherFunction(111, 12, someFunction1))
        fmt.Println(someOtherFunction(111, 12, someFunction2))
        
      //Method-2  
      m := map[string]func(int, int) int {
              "someFunction1": someFunction1,
              "someFunction2": someFunction2,
      }
      z := someOtherFunction(x, y, m[key])
      fmt.Println(z)

}
```
## Struct to Fields (Embedded type included)

```go

var tagGetName string = "json" //bson|msgpack

func StructToFields(anyStruct interface{}) map[string]map[string]interface{} {

	stype := reflect.TypeOf(anyStruct)
	sf := structFields(stype)
	return structFieldMap(sf)
}

func structFieldMap(sf []reflect.StructField) map[string]map[string]interface{} {

	rmap := make(map[string]map[string]interface{})
	for i, field := range sf {
		valm := make(map[string]interface{})
		valm["name"] = field.Name
		valm["rtype"] = field.Type.String() //reflect.Type
		valm["tag"] = field.Tag.Get(tagGetName)
		valm["isexported"] = field.IsExported()
		valm["position"] = i
		//valm["anonymous"] = field.Anonymous //if embedded type then true
		rmap[field.Name] = valm
	}
	return rmap
}

func structFields(stype reflect.Type) []reflect.StructField {

	sf := make([]reflect.StructField, 0)
	if stype.Kind() == reflect.Pointer {
		stype = stype.Elem()
	}
	if stype.Kind() == reflect.Struct {
		for i := 0; i < stype.NumField(); i++ {
			f := stype.Field(i)
			if f.Anonymous {
				rtype := f.Type
				rs := structFields(rtype)
				sf = append(sf, rs...)
			} else {
				sf = append(sf, f)
			}
		}
	}
	return sf
}
```

## Usage example
```go
type People struct{
    Name string `json:"name"`
    Dob string `json:"dob"`
    Height float64 `json:"height"`
}

type Employee struct{
    People
    Salary float64
    Status int
}

//struct
emp := Employee{}
rmap := StructToFields(emp)
for key, val := range rmap {
    fmt.Println(key, val)
}

//pointer to struct
emp2 := &Employee{}
rmap = StructToFields(emp2)
for key, val := range rmap {
    fmt.Println(key, val)
}
```

## Form filler | Form + StructReference = StructPropertyWithValue
```go
func structFiller(form url.Values, anyStructToPointer interface{}) error {

	empv := reflect.ValueOf(anyStructToPointer) //must be a pointer
	empt := reflect.TypeOf(anyStructToPointer)  //

	if empv.Kind() != reflect.Pointer {
		return errors.New("anyStructToPointer must be a pointer")
	}

	for i := 0; i < empv.Elem().NumField(); i++ {

		var vField reflect.Value
		var sField reflect.StructField

		vField = empv.Elem().Field(i)
		if empt.Kind() == reflect.Struct {
			sField = empt.Field(i)
		}

		if empt.Kind() == reflect.Pointer {
			sField = empt.Elem().Field(i) //?
		}

		fieldValue := structFieldValue(sField, "json")
		fvalue := form.Get(fieldValue)
		valueSet(vField, fvalue) //

		if sField.Type.Kind() == reflect.Struct {
			for j := 0; j < vField.NumField(); j++ {
				ssField := sField.Type.Field(j)
				fieldValue := structFieldValue(ssField, "json")
				fvalue := form.Get(fieldValue)
				valueSet(vField.Field(j), fvalue)
			}
		}
	}
	return nil
}

func valueSet(vField reflect.Value, fvalue interface{}) {

	if vField.Kind() == reflect.String {
		vField.SetString(fmt.Sprint(fvalue))

	} else if vField.Kind() == reflect.Int64 {
		fvalueInt, _ := strconv.ParseInt(fmt.Sprint(fvalue), 10, 64)
		vField.SetInt(fvalueInt)

	} else if vField.Kind() == reflect.Int {
		fvalueInt, _ := strconv.ParseInt(fmt.Sprint(fvalue), 10, 64)
		vField.SetInt(fvalueInt)

	} else if vField.Kind() == reflect.Float64 {
		fvalueInt, _ := strconv.ParseFloat(fmt.Sprint(fvalue), 64)
		vField.SetFloat(fvalueInt)

	} else if vField.Kind() == reflect.Bool {
		boolVal, _ := strconv.ParseBool(fmt.Sprint(fvalue))
		vField.SetBool(boolVal)

	} else if vField.Kind() == reflect.Slice {
		vslc := strings.Split(fmt.Sprint(fvalue), ",")
		var strValue reflect.Value = reflect.ValueOf(vslc)
		vField.Set(strValue)
	}
}

func structFieldValue(field reflect.StructField, tagField string) string {
	var fieldValue string = field.Tag.Get(tagField)
	if fieldValue == "" {
		fieldValue = strings.ToLower(field.Name)
	}
	return fieldValue
}
```

## Form filler usage example

```go

type BasicInfo struct {
	Name   string `json:"name"`
	Age    int64
	Gender string
}

type Employee struct {
	BasicInfo
	Designation string
	Salary      float64
	Adult       bool
	Hobby       []string
	Status      int
}

form := make(url.Values)
form.Set("name", "Wania")
form.Set("age", "2")
form.Set("gender", "female")
form.Set("designation", "Engineer")
form.Set("salary", "5000.50")
form.Set("status", "1")
form.Set("adult", "1")
form.Set("hobby", "quran,reading,swimming")

var emp Employee
err := structFiller(form, &emp) //form + pointer to struct | struct reference
fmt.Println(err)

```


## Reference
* https://stackoverflow.com/questions/18017979/golang-pointer-to-function-from-string-functions-name
* https://mikespook.com/2012/07/function-call-by-name-in-golang/

## Resource
* [Bitshifting Basic video](https://www.youtube.com/watch?v=qq64FrA2UXQ)
* https://github.com/yangchenxing/go-map2struct
* https://github.com/yangchenxing/go-singleflight
* [Map to Struct](https://github.com/mateors/go-map2struct)
* [free-memory-once-occupied](https://stackoverflow.com/questions/37382600/cannot-free-memory-once-occupied-by-bytes-buffer)
* [Full-text-search-engine](https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine)
