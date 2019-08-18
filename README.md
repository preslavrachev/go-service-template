# Why another template?
Developing a new application in Go is a wonderful experience, but scaling it is a rather tricky matter. The Go community has been relentless about code simplicity, and avoiding abstraction layers unless necessary. This works fine when keeping applications small, but there is little consensus of how to organize apps with the intention of scaling them beyond a certain extent.

I have looked at various approaches of organizing a Go application. From the bare-bones, proudly setting a database instance directly in the HTTP handler, to the ones resembling an enterprise Java app - interfaces with tens of methods that need to be implemented on every step. Neither really worked for me. The former makes the entire application too dependent on the infrastructure (DB, message broker, etc), whose replacement requires touching the entire application. The latter on the other hand, requires spending lots of time before introducing the slightest addition to the code.

# How is this one different?
In a way, my approach tries to learn from both the aforementioned approaches. It does create a separation of responsibilities, but does it in way that benefits from Go's unique features. 

## Declare interfaces at the consuming side not where you implement them
TBA
## Chunk logic into actionable items
TBA

# Learning Resources

- [How I write Go HTTP services after seven years - Statuscode - Medium](https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831)

- [Standard Package Layout - Ben Johnson - Medium](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1)
