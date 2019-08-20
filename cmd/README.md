This is where your executables will go. Each executable should be put into a separate package, corresponsing to the final artifact's name (e.g. `webservice`, `adminservice`, `cli`, etc) . Inside the package, create a main execution point: typically, this will be a `main.go` file containing a `main` function. You can have as many executables in your app as you like.

```
├── cmd
│   ├── adminservice
│   │   └── main.go
│   ├── cli
│   │   └── main.go
│   ├── cronjobs
│   │   └── main.go
│   └── webservice
│       └── main.go
```
