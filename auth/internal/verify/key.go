package verify

import (
	"live-auth/internal/database"
	"log"
)

type Verify struct {
	Key string
	AuthKey string
}

type Key struct {
	Streaming_key string
}

func New(recivedkey string) *Verify {
	return &Verify{
		Key: recivedkey,
	}
}


func (v *Verify) Ok() bool {

	conn, err := database.Connection()
	
	if (err != nil) {
		log.Println(err)
	}

	var DbKey *Key

	conn.First(&DbKey, "streaming_key = ?", v.Key)

	v.AuthKey = DbKey.Streaming_key

	return v.Key == v.AuthKey
}