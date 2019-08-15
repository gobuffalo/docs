## `pathFor` *helper*

Le *helper* `pathFor` prend en paramètre une `interface{}`, ou une `slice` d'`interface{}`,
et il essaie de le convertir en un chemin d'URL de la forme `/foos/{id}`.

Règles :
* si le paramètre est une `string`, il est retourné tel quel.
* si c'est un type `Pathable`, le résultat de la méthode `ToPath` est retourné.
* si c'est une `slice` ou un `array`, chaque élément est traité en appliquant ces règles, puis le résultat est joint.
* si c'est une `struct`, le nom pluralisé de la struct est utilisé pour définir le nom.
* si c'est un type `Paramable`, la méthode `ToParam` est appelée pour remplir l'`{id}`.
* si dans la struct, un attribut `Slug` est présent, le slug est utilisé pour remplir l'`{id}`.
* si dans la struct, un attribut `ID` est présent, l'ID est utilisé pour remplir l'`{id}`.

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
