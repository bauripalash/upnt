# UPNT

A simple to program to read a string from STDIN and convert each character to it its Unicode Notation.

## Example

* Create a input file
`$ echo "Hello World!" > hw.txt`

* Pipe that to **upnt**

`$ cat hw.txt | upnt`

* Output ->
```
U+0048
U+0065
U+006C
U+006C
U+006F
U+0020
U+0057
U+006F
U+0072
U+006C
U+0064
U+0021
```


# Install

Download the compiled binary for your OS from Release page; and put that somewhere in your `$PATH`

# Build

> No external library needed.

`make build`

# Test

`make test`
