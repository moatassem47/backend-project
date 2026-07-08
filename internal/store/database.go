package store

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)
func Open()(*sql.DB, error){
db,err:= sql.Open("pgx","host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable,db.SetMaxOpenConns(), db.SetMaxIdleConns(), and db.SetConnMaxIdleTime()")
if err!=nil{
	return nil,fmt.Errorf("db:open %w",err)
}
fmt.Println("connected to Database")
return db,nil
}