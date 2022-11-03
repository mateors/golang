## Mutex -> sync.Mutex
Mutex is short for mutual exclusion. Mutexes keep track of which thread has access to a variable at any given time.
![Mutex](https://res.cloudinary.com/practicaldev/image/fetch/s---IMLhEFN--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://i1.wp.com/qvault.io/wp-content/uploads/2020/03/download.png%3Fw%3D742%26ssl%3D1)

## Read & Write Lock
* Read lock: multiple concurrent read operations can be performed at the same time, and write operations are not allowed
* Write lock: only one coroutine is allowed to write at the same time, and other write and read operations are not allowed

* Read only mode: multi process can be read but not written
* Write only mode: single co process can be written but not readable

```go
	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}
  
	var result []float64 //shared resource

	wg.Add(1)
	go func() {
		mut.Lock()
		fmt.Println("worker 1")
		result = append(result, 50.50)
		mut.Unlock()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		mut.Lock()
		fmt.Println("worker 2")
		result = append(result, 78.50)
		mut.Unlock()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(result)
  ```
  

## Semaphore
The semaphore is the concept that allows In-N-Out to receive 4 orders in a restaurent concurrently (actually in parallel), causing everyone else to sit and wait.

### Let's compare a mutex to a sempahore
* If a mutex is concerned with ensuring a single thread ever accesses code exclusively 
* a semaphore is concerned with ensuring at most N threads can ever access code exclusively.

> a semaphore is a more generalized version of a mutex 

### What's the point of giving exclusive access to N threads? 
The point is that in this scenario you are purposefully constraining access to a resource therefore protecting that resource from overuse.

> **mutex**: constrains access to a 1 single thread, to guard a critical section of code.

> **semaphore**: constrains access to at most N threads, to control/limit concurrent access to a shared resource

### When using a semaphore how do you figure out the value of N for how many threads to limit?
Unfortunately there is no hard and fast rule and the final number of N is going to depend on many factors. A good place to start is by benchmarking and actually hitting your shared resource to see where it starts to fall over in terms of performance and latency.

## Resource
* [Mutex Fundamentals](https://www.sohamkamani.com/golang/mutex)
* [rwMutex](https://dev.to/qvault/golang-mutexes-what-is-rwmutex-for-57a0)
* [Read-Write Code Analysis](https://www.mo4tech.com/golang-series-rwmutex-read-write-lock-analysis.html)
* [sync.Mutex Video Explanation](https://www.youtube.com/watch?v=JlmYLPxwVzQ)
* [Mutex vs Synchronization](https://www.youtube.com/watch?v=jkRN9zcLH1s) 
* [Semaphore and Mutex](https://www.youtube.com/watch?v=DvF3AsTglUU)
* [concurrency-semaphores](https://medium.com/@deckarep/gos-extended-concurrency-semaphores-part-1-5eeabfa351ce)
