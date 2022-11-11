package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

const (
	host     = "host.docker.internal"
	port     = 5432
	user     = "admin"
	password = "123"
	dbname   = "marketplace"
)

func Update(db *sql.DB, users *Users, mx *sync.Mutex) {
	for {
		mx.Lock()
		for k, v := range users.Users {
			_, errq := db.Exec(`UPDATE marketplace SET money=$1 WHERE userid=$2`, v.Money, k)
			if errq != nil {
				log.Panicln(errq)
			}
		}
		mx.Unlock()
	}
}
func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()
	mux := &sync.Mutex{}
	users := Users{}
	users.InitUsers()
	sqlStatement := `
INSERT INTO marketplace (money,userid)
VALUES ($1, $2)`
	for i := 0; i < 5; i++ {
		users.AddUser()
	}
	for _, v := range users.Users {
		errq := db.QueryRow(sqlStatement, v.Money, v.Id)
		if errq == nil {
			log.Panicln(errq)
		}
	}
	fmt.Println(users.Users)
	go Update(db, &users, mux)
	r := gin.Default()
	h := Handler{ref: &users}
	r.GET("/all", h.All)
	r.POST("/donate", h.Donate)
	r.POST("/trade", h.Trade)
	r.POST("/info", h.Info)
	r.POST("/buy", h.Buy)
	r.POST("/accept", h.Accept)
	if err := r.Run(); err != nil {
		log.Panicln("Cant start router!")
	}
}
