# Go Inspect

Small object inspection library inspired from Ruby’s `Object#inspect` method for
learning purposes to my self :) Found Go’s `reflect` library to play with!

## Install

```bash
$ go get -u github.com/vigo/go-inspect
```

Run the test:

```go
$ go test -v .
```

## Usage

It works (*as far as I tested*) with numerics, strings, arrays, maps and structs.

```go
package main

import (
	"fmt"

	inspect "github.com/vigo/go-inspect"
)

// User represents basic user data
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// GetFullName returns user's fullname
func (u User) GetFullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

// SetEmail sets user's email with given input
func (u *User) SetEmail(email string) {
	u.Email = email
}

func main() {
	user1 := User{
		FirstName: "Uğur",
		LastName:  "Özyılmazel",
		Email:     "vigo@example.com",
	}
	user2 := &User{
		FirstName: "Ezel",
		LastName:  "Özyılmazel",
		Email:     "ezel@example.com",
	}

	fmt.Println(inspect.Element(&user1))
	fmt.Println(inspect.Element(user2))
}
```

Output:

    Inspection of: [User] object
    
    [*main.User..............] {Uğur Özyılmazel vigo@example.com}
    # of Fields: 3
    
    FirstName........... string.... Uğur.......................... json:"first_name".............
    LastName............ string.... Özyılmazel.................... json:"last_name"..............
    Email............... string.... vigo@example.com.............. json:"email"..................
    
    # of Methods: 2
    
    GetFullName......... func(*main.User) string
    SetEmail............ func(*main.User, string)
    
    --------------------------------------------------------------------------------------------------------------------
    
    Inspection of: [User] object
    
    [*main.User..............] {Ezel Özyılmazel ezel@example.com}
    # of Fields: 3
    
    FirstName........... string.... Ezel.......................... json:"first_name".............
    LastName............ string.... Özyılmazel.................... json:"last_name"..............
    Email............... string.... ezel@example.com.............. json:"email"..................
    
    # of Methods: 2
    
    GetFullName......... func(*main.User) string
    SetEmail............ func(*main.User, string)
    
    --------------------------------------------------------------------------------------------------------------------

Feel free to fix/improve anything you like to!

## Contributer(s)

* [Uğur "vigo" Özyılmazel](https://github.com/vigo) - Creator, maintainer

---

## Contribute

All PR’s are welcome!

1. `fork` (https://github.com/vigo/go-inspect/fork)
1. Create your `branch` (`git checkout -b my-features`)
1. `commit` yours (`git commit -am 'added killer options'`)
1. `push` your `branch` (`git push origin my-features`)
1. Than create a new **Pull Request**!

---

## License

This project is licensed under MIT

---

