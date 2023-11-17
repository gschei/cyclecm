package rdb

import (
	"database/sql"
	"log"
	"sync"

	"github.com/gschei/cyclecm/database"
	"github.com/gschei/cyclecm/domain/club"
)

type DbRepository struct {
	dbConnection *sql.DB
}

var dbRepository *DbRepository
var once sync.Once

func New() *DbRepository {
	once.Do(func() {
		dbRepository = &DbRepository{
			dbConnection: database.GetDbConnection(),
		}
	})
	return dbRepository
}

func (dbr *DbRepository) Get(id int64) (club.Club, error) {

	rows, err := dbr.dbConnection.Query("select id,name from club where id=$1", id)
	if err != nil {
		log.Fatal("ERROR in query to PG: ", err)
	}
	defer rows.Close()

	var c *club.Club

	for rows.Next() {
		var (
			id   int64
			name string
		)

		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		c, _ = club.NewClub(name)
		c.ID = id
	}

	if c != nil {
		return *c, nil
	} else {
		return club.Club{}, club.ErrClubNotFound
	}

}

func (dbr *DbRepository) Add(c club.Club) (club.Club, error) {
	var newID int
	err := dbr.dbConnection.QueryRow("insert into club (name) values ($1) returning id", c.Name).Scan(&newID)
	if err != nil {
		log.Printf("ERROR inserting row: %v", err)
		return club.Club{}, err
	}
	log.Printf("created club with id %v", newID)
	c.ID = int64(newID)
	return c, nil
}

func (dbr *DbRepository) Update(c club.Club) error {
	_, err := dbr.dbConnection.Exec("update club set name=$1 where id=$2", c.Name, c.ID)
	if err != nil {
		log.Printf("ERROR updating row: %v", err)
		return err
	}
	return nil
}

func (dbr *DbRepository) GetAll() []club.Club {
	rows, err := dbr.dbConnection.Query("select id,name from club")
	if err != nil {
		log.Fatal("ERROR in query to PG: ", err)
	}

	clubs := make([]club.Club, 0, 5)
	for rows.Next() {
		var (
			id   int64
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		c, _ := club.NewClub(name)
		c.ID = id
		clubs = append(clubs, *c)
	}

	return clubs
}
