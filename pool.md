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
```

## Resource
* [Go Concurrency Worker Pool Pattern like I’m five](https://itnext.io/explain-to-me-go-concurrency-worker-pool-pattern-like-im-five-e5f1be71e2b0)
* [Golang package](https://github.com/mateors/pond)
* [Worker pool video tutorial](https://www.youtube.com/watch?v=1iBj5qVyfQA)
* [Implemention Resource Pool Part-1](https://www.youtube.com/watch?v=G33OlABzxW8)
* [Part-2](https://www.youtube.com/watch?v=HNZTJZ7Buk0)
* [Part-3](https://www.youtube.com/watch?v=h6ZwhM20_Yo)
* [Resouce-pool code](https://github.com/striversity/gotr/tree/master/021-resouce-pool)
* [Detailed-explanation](https://developpaper.com/detailed-explanation-of-the-use-of-sync-pool-in-golang/)
* [Handling-1-million-requests-per-minute](http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/)
* [Mysql 1 billion record processing](https://github.com/rcbensley/brimming)
* [Sqlite data generator](https://github.com/sqlitebrowser/sqlitedatagen)
* [multithreading-to-read-large-files](https://hackernoon.com/leveraging-multithreading-to-read-large-files-faster-in-go-lmn32t7)
* [Build a TCP Connection Pool From Scratch](https://betterprogramming.pub/build-a-tcp-connection-pool-from-scratch-with-go-d7747023fe14?gi=b3978a42d21d)
* [Simple code](https://chowdera.com/2021/07/20210715165328495l.html)
* [how-to-make-1-million-get-request](https://forum.golangbridge.org/t/how-to-make-1-million-get-request-with-golang-concurrency-in-low-end-windows-pc/16628)
