package domain

import "time"

type Entity struct {
    ID        int
    Name      string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type EntityRepository interface {
    GetByID(id int) (*Entity, error)
    // other repository methods...
}

type EntityUsecase interface {
    GetByID(id int) (*Entity, error)
    // other use case methods...
}
