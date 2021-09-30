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

func Alloc() *MyModel {
	select {
	case cr := <-pool:
		return cr
	default:
		cr := &MyModel{}
		return cr
	}
}

func Release(cr *MyModel) {
	select {
	case pool <- cr:
	default:
	}
}

```
