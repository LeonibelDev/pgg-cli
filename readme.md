
# pgg — Password Generator in Golang

![pgg cover](assets/cover.png)

**pgg** is a secure, terminal-based password generator written in **Go**, built on top of `crypto/rand`.  
It allows you to generate strong passwords, save them locally, copy them to the clipboard, and manage them directly from the command line.

This project is designed as a lightweight and secure alternative to GUI-based password managers, focused on developer workflows.

---

## Features

- Cryptographically secure password generation using `crypto/rand`
- Configurable password length and character sets
- Save generated passwords with a service identifier
- Copy passwords to clipboard
- List, delete, and export saved passwords
- Built using Cobra for a clean CLI experience

---

## Installation

### Using Go


go install github.com/leonibeldev/pgg@latest


### Build from source

```bash
git clone https://github.com/leonibeldev/pgg.git
cd pgg
go build -o pgg
```

---

## Usage

Generate a password with a custom length and character types:

```bash
pgg -l 10 -t lower,numbers,special -s facebook.com
```

Example output:

```bash
Password 4sd56456sa3?w2
```

---

## Flags

### Password Generation

| Flag       | Alias | Description                       | Default                       |
| ---------- | ----- | --------------------------------- | ----------------------------- |
| `--length` | `-l`  | Length of the password            | `16`                          |
| `--types`  | `-t`  | Character types to include        | `numbers,upper,lower,special` |
| `--save`   | `-s`  | Save password with a service name | `""`                          |
| `--copy`   | `-c`  | Copy password to clipboard        | `false`                       |

---

### Character Types

| Code | Description        |
| ---- | ------------------ |
| `lower`  | Lowercase letters  |
| `upper`  | Uppercase letters  |
| `numbers`  | Numbers (0–9)      |
| `special`  | Special characters |

Example:

```bash
pgg -l 12 -t lower,numbers
```

---

## Password Management

### List Saved Passwords

```bash
pgg -list
```

Output example:

```text
-----------------------------------------------
|   ID   |     Service     |      Password    |
|---------------------------------------------|
|  5132  |  facebook.com  |  4sd56456sa3?w2   |
-----------------------------------------------
```

---

### Delete a Password

```bash
pgg -delete 5132
```

Confirmation:

```text
You want to delete the password with ID 5132? Y/n
```

---

### Export Passwords

Export all stored passwords to a JSON file:

```bash
pgg -export -filename keys
```

Example output file:

```json
{
  "keys": [
    {
      "ID": 5132,
      "Service": "facebook.com",
      "Password": "fd538245d2z105104:'s"
    }
  ]
}
```

---

## Version

```bash
pgg version
```

---

## Project Structure

* `cmd/` – CLI commands and flags
* `crypt/` – Secure password generation logic
* `flags/` – Flag parsing and validation
* `utils/` – Persistence and helper utilities

---

## Security Notes

* Passwords are generated using `crypto/rand`, not `math/rand`
* No passwords are sent over the network
* Local storage responsibility belongs to the user

---

## Roadmap

* Encrypted local storage
* Clipboard auto-clear timeout
* Config file support
* Cross-platform clipboard improvements

---

## License

Copyright © 2026
LeonibelDev [leonibel.ramirez@gmail.com](mailto:leonibel.ramirez@gmail.com)
