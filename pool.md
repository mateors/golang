## Custom Pool implementation using golang

```golang
package cpool

type MyModel struct{
  Name string
  Salary float64
  Age int
}

const POOL_SIZE = 1000

var pool chan *MyModel

func init() {
	pool = make(chan *MyModel, POOL_SIZE)
}

//Add/Allocate resource to the pool or pick from the existing pool
// Get()
func Get() *MyModel {
	select {
	case cr := <-pool:
		return cr
	default:
		cr := &MyModel{}
		return cr
	}
}

//Release resource from the pool
//Garbage Collector (GC) is relaxed now
//Put() | put resource back
func Put(cr *MyModel) {
	select {
	case pool <- cr:
	default:
	}
}

```
