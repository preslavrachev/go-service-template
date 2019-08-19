# Why another template?

Developing a new application in Go is a wonderful experience, but scaling it is a rather tricky matter. The Go community has been relentless about code simplicity, and avoiding abstraction layers unless necessary. This works fine when keeping applications small, but there is little consensus of how to organize apps with the intention of scaling them beyond a certain extent.

I have looked at various approaches of organizing a Go application. From the bare-bones, proudly setting a database instance directly in the HTTP handler, to the ones resembling an enterprise Java app - interfaces with tens of methods that need to be implemented on every step. Neither really worked for me. The former makes the entire application too dependent on the infrastructure (DB, message broker, etc), whose replacement requires touching the entire application. The latter on the other hand, requires spending lots of time before introducing the slightest addition to the code.

# How is this one different?

In a way, my approach tries to learn from both the aforementioned approaches. It does create a separation of responsibilities, but does it in way that benefits from Go's unique features.

## Declare interfaces at the consuming side, not where you implement them

Interfaces in Go are unique, in that you don't have to explicitly bind a type to an interface, in order to implement it. It is enough that the "implemntor" conforms to the contract of the interface. Where the interface is declared is totally irrelevant for the potential implementor. This opens up a whole new world of possibilities, where one can define an interface where one is needed, and start using it right away. Those can even be anonymous, so one won't need a separate type declaration.

```go
func (s *Server) HandleTodoListGet(todoHandler interface {
	FindAllTodos() []*myservice.Todo
}) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(todoHandler.FindAllTodos())
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
}
```

The great thing is that anything can be a `todoHandler`, as long as it conforms to the contract (i.e. has an exported `FindAllTodos` method). This makes it incredibly to swap one handler for another, especially in tests. It also brings me to the next point:

### Keep interfaces as small as possible

Too many times, have I seen interfaces with close to 20 methods or more. This is a direct heritage of the limitations of interfaces in other programmign languages. Typically, onewould avoid making a class explicitly implement 20 interfaces, but rather, shove all fo tehir methods into one or two. IMHO, those are not helpful, because even if an interface consumer needs only one of those, one needs to implement all of them. This is especially cumbersome in testing. As we saw above, Go takes the concept of interfaces upside down, and makes it possible for any type to implement an interface without implicitly knowing about it. Unless an interface's methods are really bound together and always used in a particular workflow, it makes little sense to cram them all up in the same interface.

## Chunk logic into actionable items

TBA

# Learning Resources

- [How I write Go HTTP services after seven years - Statuscode - Medium](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831)

- [Standard Package Layout - Ben Johnson - Medium](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
