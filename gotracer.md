# Go Tracer
> an introduction to go tool trace and describing how you can use it to debug concurrent Go programs.
Go makes it easy to use concurrency with goroutines, which means we do more concurrency, which means we have more concurrency problems.

## Large File Read challenge 
* https://marcellanz.com/post/file-read-challenge/

## How to generate trace.out file?
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
## run your go program using 
> go run main.go

## run following code to open trace.out file in your browser
> go tool trace trace.out

## Reference
* [Intrdoduction Blog](https://about.sourcegraph.com/go/an-introduction-to-go-tool-trace-rhys-hiltner)
* [Presentation Slide](https://speakerdeck.com/rhysh/an-introduction-to-go-tool-trace?slide=9)
* [Trace tool](https://making.pusher.com/go-tool-trace)
* [Youtube video tutorial](https://www.youtube.com/watch?v=Xq5HDH8y0CE&t=4s)
