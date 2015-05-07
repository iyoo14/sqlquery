package sqlquery

import (
    "fmt"
    "testing"
)

type Member struct {
    Id int
    Name string
    Gender int
    Gold  bool
}

func TestNewQueryBuilder(t *testing.T) {
    var mem interface{}
    mem = Member{1, "bill", 1, true}
    qb := NewQueryBuilder(mem)
    fmt.Println(qb)
}

