# gde
Set environment variables from file for Go debug.  

# Example
env file `dev_env`:
```
export MYENV=123
```
go file:
```
package main

func main() {
    gde.Setenv("./dev_env")
    // do something
}
```
then we can start debug with tools like delve.

