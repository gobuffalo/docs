<%= title("Le type Session") %>

Le type `buffalo.Session` contient tout le nécessaire pour travailler avec une session attachée à une requête. Buffalo utilise le paquet [github.com/gorilla/sessions](http://www.gorillatoolkit.org/pkg/sessions) derrière le décors pour gérer la session.

```go
type Session
  // Clear vide les données de la session
  func (s *Session) Clear()
  // Delete supprime une donnée en particulier
  func (s *Session) Delete(name interface{})
  // Get récupère une valeur spécifique
  func (s *Session) Get(name interface{}) interface{}
  // GetOnce récupère une valeur spécifique et la supprime
  func (s *Session) GetOnce(name interface{}) interface{}
  // Save sauvegarde la session
  func (s *Session) Save() error
  // Set place une valeur dans la session
  func (s *Session) Set(name, value interface{})
```
