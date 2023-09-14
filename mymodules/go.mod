module github.com/Abhijeet6387/golang-tutorial

go 1.21.1

// go mod tidy is a crucial command for dependency management in Go, helping you keep your project's dependencies organized, up-to-date, and free from unnecessary or unused packages.
// go.sum plays a pivotal role in ensuring that your project uses the correct and unaltered versions of dependencies.
// We can run go mod verify to check if all the dependencies are verified or not
// We can use go list -m all to view all packages that our project is dependent on
// We can use go mod vendor (like node modules) for managing dependencies
// We can run go run -mod=vendor main.go to fetch dependencies from vendor folder
require github.com/gorilla/mux v1.8.0
