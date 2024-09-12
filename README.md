# Text Completion/Editing/Auto-Correction Tool



## Introduction

- The project must be implemented in Go.
- The code should adhere to good programming practices.


The tool you'll develop will take as input a file containing text that requires modifications and another file to store the modified text. The following modifications should be implemented:

1. Replace hexadecimal numbers denoted by `(hex)` with their decimal equivalents.
2. Replace binary numbers denoted by `(bin)` with their decimal equivalents.
3. Convert words to uppercase if preceded by `(up)`.
4. Convert words to lowercase if preceded by `(low)`.
5. Capitalize words if preceded by `(cap)`.
6. If a number follows `(low)`, `(up)`, or `(cap)`, apply the transformation to the specified number of words.
7. Ensure punctuation marks `.`, `,`, `!`, `?`, `:`, and `;` are correctly placed with proper spacing.
8. Handle groups of punctuation marks like `...` or `!?`.
9. Place single quotes `' '` around words enclosed within them.
10. Convert 'a' to 'an' if followed by a word starting with a vowel or 'h'.

## Allowed Packages

Standard Go packages are permitted for this project.

## Usage

```bash
$ cat sample.txt
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.
$ go run . sample.txt result.txt
$ cat result.txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair.
