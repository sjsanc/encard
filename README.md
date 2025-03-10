<p align="center">
  <img src="docs/logo.png" alt="Encard Logo" width="86">
</p>

<h1 align="center">Encard</h1>

<div align="center">

[![Build Status](https://img.shields.io/github/actions/workflow/status/sjsanc/encard/.github/workflows/go.yml?branch=master)](https://github.com/sjsanc/encard/actions)
</div>

Encard is a filesystem-based flashcard TUI. It stores your cards inside regular files, allowing you to manipulate them with existing shell tooling. 

## Cards

### Basic
```md
# What is the capital of Laos?
Vientiane
```

### Cloze
```md
# Complete the quote
It always seems {{impossible}} until it's {{done}}
```

### Reversible
Reversible cards are loaded twice, filling in the curly braces.
```md
# Translate {}
House
Haus
```
This generates two cards:

Front: "Translate House" → Back: "Haus"\
Front: "Translate Haus" → Back: "House"

### Multiple Choice
```md
# Which of these programming languages is best?
- C#
* Go
- Javascript
- Rust
```

### Multiple Answer
```md
# Which of these countries are in South America?  
[*] Brazil  
[*] Argentina  
[ ] Mexico  
[ ] Spain 
```

## Credits

Largely inspired by [hascard](https://github.com/Yvee1/hascard?tab=readme-ov-file).