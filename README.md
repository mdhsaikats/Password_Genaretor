# Password Generator (Go)

Simple command-line password generator written in Go.

## Description

`PasswordGenaretorGoLang` is a minimal CLI tool that generates a secure random password using Go's `crypto/rand` package. The generator lives in `main.go` and prompts the user for a password length, then prints a randomly-generated password containing lowercase, uppercase, digits, and special characters.

## Features

- Uses cryptographically secure randomness (`crypto/rand`).
- Supports letters, digits, and a set of special characters.
- Tiny, dependency-free single-file Go program (`main.go`).

## Usage

Run directly with `go run`:

```bash
go run main.go
```

When prompted, enter the desired numeric password length (for example `16`). The program prints the generated password to stdout.

You can build a binary and run it as well:

```bash
go build -o password
./password
```

## Example

Input:

```
Enter password length: 12
```

Output (example):

```
Random Password is: aB3$k9!Lm2Q@
```

Note: Actual output will differ each run since the password is randomly generated.

## Notes & Suggestions

- The program currently reads length via `fmt.Scan` and does not validate the input; negative or non-numeric input may cause unexpected behavior. Consider adding input validation.
- The exported function name `PaswordGenaretor` has a typo; renaming to `PasswordGenerator` or `GeneratePassword` would improve clarity.
- Adding a command-line flag (for example `-length`) would make the tool scriptable and more flexible.

## Security

This tool uses `crypto/rand` to produce cryptographically secure random bytes. The character set includes special characters; if you need passwords compatible with a specific system, adjust the allowed characters in `main.go` accordingly.

## License

This repository is provided without a specific license. If you plan to publish it, consider adding an open-source license (for example, MIT).
# Password_Genaretor
