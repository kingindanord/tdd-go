
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