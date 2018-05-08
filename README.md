# gopartial - Golang Partial Struct Update

This mini library provides a helper function to update go object partially from incoming json data.

## Why?

The challenge with golang is it is hard to take a dynamic data such as json and apply partial update
to a golang struct due to the 'unknown' values that you might receive from a field in a json data.

The challenge becomes even harder when dealing with nullable data because
if you use json null value as a way to determine whether a struct field should be updated or not,
you'll lose ability to update the struct to null value.

A good practical example of this problem is when you want to implement HTTP `PATCH` Request. 

## Installation

```
$ go get github.com/roserocket/gopartial
```

## Dependencies

```
$ go get github.com/guregu/null
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
    DeletedAt   *time.Time      `json:"deleted_at"` // Can be null
}

// Imagine you have a Web API that can partially update an existing User in database
func UpdateUserPartially(user *User, partialDataJSON json.RawMessage) (*User, error) {


    var partialData map[string]interface{}{}
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

## Methods

#### `func PartialUpdate(dest interface{}, partial map[string]interface{}, tagName string, skipConditions []func(reflect.StructField) bool, updaters []func(reflect.Value, reflect.Value) bool) ([]string, error)`

|    Argument    |                    Type                     |                                         Description                                         |
| :------------: | :-----------------------------------------: | :-----------------------------------------------------------------------------------------: |
|      dest      |                `interface{}`                |                      Destination struct (Must be a pointer to struct)                       |
|    partial     |          `map[string]interface{}`           |                     Partial data in the form of map[string]interface{}                      |
|    tagName     |                  `string`                   | The struct tag name that you'll be mapping the struct field to based on the json field name |
| skipConditions |     `[]func(reflect.StructField) bool`      |                              Array of skip condition functions                              |
|    updaters    | `[]func(reflect.Value, reflect.Value) bool` |                                 Array of updater functions                                  |

This function can be easily extended if you have certain skip conditions while updating the struct.
For example you want to skip all the struct field that has tagname `props` with value of `readonly`, then you can create a function as follow:

```go
// SkipReadOnly skips all field that has tag readonly
func SkipReadOnly(field reflect.StructField) bool {
	props := strings.Split(field.Tag.Get("props"), ",")
	return utils.IndexOf(props, "readonly") >= 0
}
```

You can also extend this function to update a certain custom type within your application.
Example:

```go
type MyType string

// MyTypeUpdater update MyType
func MyTypeUpdater(fieldValue reflect.Value, v reflect.Value) bool {
	switch fieldValue.Interface().(type) {
	case MyType:
		// if its null value
		if !v.IsValid() {
			newValue := reflect.ValueOf(MyType{}})
			fieldValue.Set(newValue)
			return true
		}
		// only set if underlying type is a string
		if v.Kind() == reflect.String {
			newValue := reflect.ValueOf(MyType(v.String()))
			fieldValue.Set(newValue)
			return true
		}
	}

	return false
}
```

### Why do we need updatedFields returned?

The idea is using the list of updated fields, you can dynamically build the sql query to update the record in the database.

Hint: use `reflect.Type.FieldByName` function to get the `reflect.StructField` and use `reflect.StructField.Tag.Get("db")`
to get the db field name.

## TODO

Currently this library does not support nested partial update.

## License

This code is free to use under the terms of the MIT license.
