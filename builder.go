package sqlquery

import (
    "log"
    "reflect"
    "strings"
)

type QueryBuilder struct {
    table string
    columns []string
}

func NewQueryBuilder(s interface{}) *QueryBuilder {
    v := reflect.ValueOf(s)
    t := v.Type()
    log.Println("Name of h: v=reflect.ValueOf(h), t=v.Type(), t.Name()")
    log.Printf("%s\n", t.Name())
    qb := new(QueryBuilder)
    qb.table = strings.ToLower(t.Name())
    log.Println("Num of Field: v=reflect.ValueOf(h), v.NumField()")
    log.Println(v.NumField())
    for i := 0; i < v.NumField(); i++ {
        log.Printf("Field Name: v=reflect.ValueOf(h), t=v.Type(), t.Field(%d).Name", i)
        name := t.Field(i).Name
        log.Println(name)
        qb.columns = append(qb.columns, strings.ToLower(name))
    }
    return qb
}

