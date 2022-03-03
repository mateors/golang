## Mutex -> sync.Mutex
Mutex is short for mutual exclusion. Mutexes keep track of which thread has access to a variable at any given time.
![Mutex](https://res.cloudinary.com/practicaldev/image/fetch/s---IMLhEFN--/c_limit%2Cf_auto%2Cfl_progressive%2Cq_auto%2Cw_880/https://i1.wp.com/qvault.io/wp-content/uploads/2020/03/download.png%3Fw%3D742%26ssl%3D1)

## Read & Write Lock
* Read lock: multiple concurrent read operations can be performed at the same time, and write operations are not allowed
* Write lock: only one coroutine is allowed to write at the same time, and other write and read operations are not allowed

* Read only mode: multi process can be read but not written
* Write only mode: single co process can be written but not readable

## Resource
* [Mutex Fundamentals](https://www.sohamkamani.com/golang/mutex)
* [rwMutex](https://dev.to/qvault/golang-mutexes-what-is-rwmutex-for-57a0)
* [Read-Write Code Analysis](https://www.mo4tech.com/golang-series-rwmutex-read-write-lock-analysis.html)
