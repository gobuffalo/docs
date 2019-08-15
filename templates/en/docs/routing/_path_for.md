## The `pathFor` Helper

The `pathFor` helper takes an `interface{}`, or a `slice` of them,
and tries to convert it to a `/foos/{id}` style URL path.

Rules:
* if `string` it is returned as is
* if `Pathable` the `ToPath` method is returned
* if `slice` or an `array` each element is run through the helper then joined
* if `struct` the name of the struct, pluralized is used for the name
* if `Paramable` the `ToParam` method is used to fill the `{id}` slot
* if `struct.Slug` the slug is used to fill the `{id}` slot of the URL
* if `struct.ID` the ID is used to fill the `{id}` slot of the URL

```go
// Car{1} => "/cars/1"
// Car{} => "/cars"
// &Car{} => "/cars"
type Car struct {
  ID int
}

// Boat{"titanic"} => "/boats/titanic"
type Boat struct {
  Slug string
}

// Plane{} => "/planes/aeroPlane"
type Plane struct{}

func (Plane) ToParam() string {
  return "aeroPlane"
}

// Truck{} => "/a/Truck"
// {[]interface{}{Truck{}, Plane{}} => "/a/Truck/planes/aeroPlane"
type Truck struct{}

func (Truck) ToPath() string {
  return "/a/Truck"
}
```
