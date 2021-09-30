## Custom Pool implementation using golang

```golang
package cpool

type MyModel struct{
  Name string
  Salary float64
  Age int
}

//you may increase accodring to your project need
const POOL_SIZE = 1000

//Global pool variable | resource pool
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

## Implementation
```golang

//import "sync/atomic"
//import "encoding/json"
//import "time"
//import "net/http"

type Data struct {
	numReqs uint64
	name    string
	value     log.Logger
}

func (d *Data) Add(){

 //Any business login implementation
 // Your code
}

func (d *Data) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	atomic.AddUint64(&d.numReqs, 1)

	go func(){
		
		dec := json.NewDecoder(r.Body)
		defer r.Body.Close()

		myModel := cp.Get()
		dec.Decode(myModel)
		// INFO - pretent we do some work on with the msg
		time.Sleep(10 * time.Millisecond)
		co.Put(myModel)
	}()
	
}
