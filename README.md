# gopartial

This mini library provides a helper function to update go object partially from incoming json data.

## Why?

The challenge with golang is it is hard to take a dynamic data such as json and apply partial update
to a golang struct due to the 'unknown' values that you might receive from a field in a json data.

The challenge becomes even harder when dealing with nullable data because
if you use json null value as a way to determine whether a struct field should be updated or not,
you'll lose ability to update the struct to null value.

## Installation

```
$ go get github.com/roserocket/gopartial
```

## Dependencies

```
$go get github.com/guregu/null
```

## Example

```go
import (
    "time"
    "log"

    "github.com/roserocket/gopartial"
)

// User struct
type User struct {
    ID          string          `json:"id"`
    Name        string          `json:"name"`
    Age         *int            `json:"age"` // Can be null
    DeletedAt   *time.Time      `json:"verified_at"` // Can be null
}

// Imagine you have a Web API that can partially update an existing User in database
func UpdateUserPartially(user *User, partialDataJSON json.RawMessage) (*User, error) {


    var partialData map[string]interface{}
    if err := json.Unmarshal(partialDataJSON, &partialData); err != nil {
        log.Fatal(err)
        return
    }

    updatedFields, err := gopartial.PartialUpdate(user, partialData, "json", gopartial.SkipConditions, gopartial.Updaters)
    log.Println("Updated fields: ", updatedFields)

    return user, err
}

func main() {
    t := time.Now()
    // Existing user data
    var user := &User{
        ID:         "1",
        Name:       "John",
        Age:        nil,
        DeletedAt:  *t,
    }

    log.Printf("Initial user object: %+v", user)

    // You want to update just name, age and deleted_at
    partialDataJSON := json.RawMessage(`{"name": "Johnson", "age": 21, "deleted_at": null}`)

    var err error
    user, err = UpdateUserPartially(user, partialDataJSON)
    if err != nil {
        log.Fatal(err)
        return
    }

    // Updated user data should now be:
    // User{
    //     ID:         "1",
    //     Name:       "Johnny",
    //     Age:        21,
    //     DeletedAt:  nil,
    // }
    log.Printf("Updated user object: %+v", user)
}
```

## TODO

Currently this library does not support nested partial update.

## License

This code is free to use under the terms of the MIT license.
