package greet

import "fmt"

// Hello returns a greeting for the supplied name.
// If name is empty it returns the default greeting.
func Hello(name string) string {
    if name == "" {
        return "Hello, World!"
    }
    return fmt.Sprintf("Hello, %s!", name)
}
