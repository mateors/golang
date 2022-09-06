# Interface

## Hexagonal Architecture
Make sure we devide our software in such a way that each peice of the software maintain its separation of concerns
by doing this make it extreamly modular. 

1. App and domain logic in the middle
2. Ports & Adapters
3. User interface
4. Repo
5. External API
6. Message Queue
7. REST | GraphQL

### Package used in this project
* github.com/teris-io/shortid
* github.com/vmihailenco/msgpack
* gopkg.in/dealancer/validate.v2

### How do I change main branch to a new branch
> `git branch -m master interface`

> `git push -d origin interface`

Inheritence can be coped with by using interface and composition

* Composition -> using type embedding
* Inheritance -> can be achived using interface


## Resource

* [Interface-pattern-in-golang](https://medium.com/swlh/what-is-the-extension-interface-pattern-in-golang-ce852dcecaec)
* [Type Embedding](https://go101.org/article/type-embedding.html)
* [Interface](https://www.youtube.com/watch?v=qJKQZKGZgf0)
* [Hexagonal Microservices with Go](https://www.youtube.com/watch?v=rQnTtQZGpg8)
* [MessagePack](https://msgpack.org/index.html)
* [Golang Tensor Programming](https://www.youtube.com/watch?v=QyBXz9SpPqE&list=PLJbE2Yu2zumCe9cO3SIyragJ8pLmVv0z9)
* [Part-3](https://github.com/tensor-programming/hex-microservice/tree/part-3)