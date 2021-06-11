# math-interpreter-go
Simple infix math interpreter in Go

## Usage

```bash
go build . [-h] [--postfix] [--backend={recursive, shunting}] [-a] [-t]
```

Flags:

- `-h`: displays a help message
- `--postfix`: sets the calculator to postfix evaluation mode (i.e. 3 + 4 * 5 -> 3 4 5 * + \))
- `--backend {recursive, shunting}`: choose between a direct recursive backend or shunting-yard to postfix backend for parsing infix expressions
- `-t`: print out the result of the lexer
- `-a`: print out the result of the parser, ignored in postfix mode
