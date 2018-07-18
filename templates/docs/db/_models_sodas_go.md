```go
package models

import (
	"time"

	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
)

type Soda struct {
	ID                   uuid.UUID    `db:"id"`
	CreatedAt            time.Time    `db:"created_at"`
	UpdatedAt            time.Time    `db:"updated_at"`
	Label                nulls.String `db:"label"`
}
```