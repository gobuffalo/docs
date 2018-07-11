<%= title("Stocker un type complexe") %>

C'est rarement une bonne idée de stocker des types complexes dans une session. Il y a plein de raisons à ça, mais il est recommandé de conserver l'ID du type à la place de la structure complète.

Si toutefois vous avez vraiment besoin de stocker un type complexe en session (comme une structure `struct`), vous devez enregistrer le type avec le paquet [`encoding/gob`](https://golang.org/pkg/encoding/gob/).

```go
import "encoding/gob"

func init() {
  gob.Register(&models.Person{})
}
```

