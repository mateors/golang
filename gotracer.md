# Go Tracer
> an introduction to go tool trace and describing how you can use it to debug concurrent Go programs.
Go makes it easy to use concurrency with goroutines, which means we do more concurrency, which means we have more concurrency problems.

## Large File Read challenge 
* https://marcellanz.com/post/file-read-challenge/

> copy and paste the following code into your main function
```golang

	//go tool trace
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	//------------------------
  
  //your prgram source code goes here
  
  ```

## Reference
* [Intrdoduction Blog](https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner)
* [Presentation Slide](https://speakerdeck.com/rhysh/an-introduction-to-go-tool-trace?slide=9)
