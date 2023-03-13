# smack
A simple RPN language

## Install
```bash
$ go install github.com/j0hax/smack
```

## Usage
Example: the input is the equivalent of `(1 + 2) * 6`, resultung in `18`.

```
1 2 + 6 *
18
```

The smack interpreter works with standard input and output streams, meaining you can perform computations with pipes:

```bash
$ echo '1 2 + 6 *' | smack > out.txt
$ smack < out.txt
18
```

## Standard Library
Very much a WIP. Some tokens for mathematical functions have been implemented, such as `log`. Check the source for the lexer.

## Error Handling
Does not exist (yet). Don't try to pop more than the stack can handle, or else you won't be doing many calculations today.

## Name Origin
I just wanted something that rhymes with *stack.*