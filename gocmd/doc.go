// package gocmd contains methods to make the usage of go commands in magefiles easier.
//
// Example:
//		// Execute the command `go get github.com/pvormste/magext@v1.0.0` and return the output
//		output, err := gocmd.GetByVersionWithOutput("github.com/pvormste/magext", "v1.0.0")
package gocmd
