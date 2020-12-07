package main

import (
	"fmt"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/mongo"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

type Person struct {
	ID     uint64 `db:"id,omitempty"`
	Name   string `db:"name"`
	Age    int    `db:"age"`
	Email  string `db:"email"`
	Gender int8   `db:"gender"`
}

func openPg() db.Session {
	settings := postgresql.ConnectionURL{
		Database: "upper_db",
		Host:     "localhost",
		User:     "postgres",
		Password: "postgres",
	}
	sess, err := postgresql.Open(settings)
	check(err)
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

	sess.SQL().Exec(`
DROP TABLE IF EXISTS persons;
CREATE TABLE "persons" (
	"id" SERIAL,
	"name" VARCHAR(255) NULL DEFAULT 'NULL::character varying',
	"age" INTEGER NOT NULL DEFAULT '0',
	"email" VARCHAR(255) NULL DEFAULT 'NULL::character varying',
	"gender" SMALLINT NULL,
	PRIMARY KEY ("id")
)
;
COMMENT ON COLUMN "persons"."id" IS '';
COMMENT ON COLUMN "persons"."name" IS '';
COMMENT ON COLUMN "persons"."age" IS '';
`)
	return sess
}

func openMongo() db.Session {
	settings := mongo.ConnectionURL{
		Database: "upper_db",
		Host:     "localhost",
		User:     "",
		Password: "",
	}
	sess, err := mongo.Open(settings)
	check(err)
	fmt.Printf("Connected to %q with DSN:\n\t%q\n", sess.Name(), settings)

	return sess
}

func main() {
	sess := openPg()
	defer sess.Close()

	col := sess.Collection("persons")

	// curd
	check(col.Insert(&Person{Name: "John1", Age: 18}))
	check(col.Insert(&Person{Name: "John2", Age: 19}))
	check(col.Insert(&Person{Name: "John3", Age: 20}))
	check(col.Insert(&Person{Name: "John4", Age: 21}))

	var persons []*Person
	check(col.Find(db.Cond{"age >=": 20}).OrderBy("-age").All(&persons))
	for _, person := range persons {
		fmt.Println(person.ID, person.Name, person.Age)
	}

	// list collections
	collections, err := sess.Collections()
	if err != nil {
		log.Fatal("Collections: ", err)
	}

	for i := range collections {
		// Name returns the name of the collection.
		fmt.Printf("-> %q\n", collections[i].Name())
	}
}

func check(a ...interface{}) {
	for _, e := range a {
		if err, ok := e.(error); ok && err != nil {
			panic(err)
		}
	}
}
