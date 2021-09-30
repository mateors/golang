## Custom Pool implementation using golang

```golang
package cpool

type MyModel struct{
  Name string
  Salary float64
  Age int
}

const CR_POOL_SIZE = 10000

var pool chan *MyModel

func init() {
	pool = make(chan *MyModel, CR_POOL_SIZE)
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
