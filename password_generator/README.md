# PASSWORD GENERATOR

A simple command-line password generator written in **Go**.

The project uses Go's `crypto/rand` package to generate cryptographically secure random values and creates passwords containing lowercase letters, uppercase letters, numbers, and symbols.

---



<br>



## REQUIREMENTS

Before running the project, make sure you have:

* **Go 1.20+**
* A terminal / command prompt
* Git *(optional, if cloning the repository)*

You can check your Go installation with:

```bash
go version
```

---



<br>



## HOW TO RUN

### 1. Clone the repository

```bash
git clone https://github.com/markankrkan/mini-projects/password_generator.git
```

### 2. Navigate into the project

```bash
cd password_generator
```

### 3. Run the program

```bash
go run .
```

The program will ask:

```text
Enter password length (default: 16):
```

Enter the desired password length and press **Enter**.

If you simply press **Enter**, the default length of **16 characters** will be used.

### Example

```text
Enter password length (default: 16): 20

Generated password:

f7@Qm2_8Zp&xK4?La9B+w

Length: 20
```

---



<br>



## FEATURES

### 🔐 Secure Random Generation

The generator uses Go's:

```go
crypto/rand
```

instead of a standard pseudo-random generator.

This provides cryptographically secure random values suitable for password generation.

### 🔢 Custom Password Length

Users can choose the password length.

```text
Enter password length (default: 16): 24
```

If no length is provided, the generator automatically uses:

```text
16
```

### 🔤 Multiple Character Types

Generated passwords can contain:

* Lowercase letters: `a-z`
* Uppercase letters: `A-Z`
* Numbers: `0-9`
* Symbols: `!#$%&/=?*+-_@`

### ✅ Character-Type Guarantee

The generator guarantees that every password contains at least:

* 1 lowercase character
* 1 uppercase character
* 1 number
* 1 symbol

For example:

```text
a
K
7
@
```

will always be represented somewhere in a generated password.

### 🔀 Password Shuffling

The required character types are initially added to the password and then shuffled using a cryptographically secure random value.

This prevents passwords from always following a predictable structure such as:

```text
lowercase
uppercase
number
symbol
```

### 📏 Minimum Length

The minimum supported password length is:

```text
4 characters
```

This is required because the generator guarantees one character from each of the four character categories.

---



<br>



## CHARACTER SETS

The generator currently uses:

```text
Lowercase
abcdefghijklmnopqrstuvwxyz

Uppercase
ABCDEFGHIJKLMNOPQRSTUVWXYZ

Numbers
0123456789

Symbols
@#$%&/=?*+-_
```

These sets can be modified in the Go source code if you want to add or remove characters.

---



<br>



## PROJECT STRUCTURE

The initial version of the project is intentionally simple:

```text
└── password_generator/
    ├── LICENSE
    ├── main.go
    └── README.md
```

### `main.go`

Contains the password generation logic and command-line interface.

### `README.md`

Project documentation, requirements, features, and usage instructions.

### `LICENSE`

Defines how the project can be used, modified, and distributed.

---



<br>



## EXAMPLE OUTPUT

```text
==============================
      PASSWORD GENERATOR
==============================

Enter password length (default: 16): 16

Generated password:

K7@xP2m_Q9&vL4za

Length: 16
```

---



<br>



## SECURITY NOTE

This project uses `crypto/rand` rather than Go's `math/rand` for password generation.

Passwords should still be handled carefully. Avoid storing generated passwords in insecure files or sharing them with other people.

For highly sensitive credentials, consider using a dedicated password manager.

---



<br>



## LICENSE

This project is licensed under the **MIT License**.

See [`LICENSE`](LICENSE) for more information.
