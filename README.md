
1. https://quii.gitbook.io/learn-go-with-tests/v/master/go-fundamentals/hello-world


### Five suggestions for setting up a Go project

https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project

### Clean Go test Cache


```
go clean -testcache
```
### use godoc

go get -v  golang.org/x/tools/cmd/godoc
sudo apt-get install golang-golang-x-tools


godoc -http=:6060
http://localhost:6060/pkg/



Unchecked errors

go get -u github.com/kisielk/errcheck


errcheck .


constant errors
https://dave.cheney.net/2016/04/07/constant-errors


https://blog.cleancoder.com/uncle-bob/2014/05/10/WhenToMock.html
https://blog.cleancoder.com/uncle-bob/2014/05/14/TheLittleMocker.html


check race condition 
go test -race


## Reflaction

[laws of reflection](https://go.dev/blog/laws-of-reflection)


use `go vet`

Remember to use go vet in your build scripts as it can alert you to some subtle bugs in your code before they hit your poor users.

[Mutex or Channel](https://github.com/golang/go/wiki/MutexOrChannel)

* Use channels when passing ownership of data
* Use mutexes for managing state