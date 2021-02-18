package argon2idencoder

import (
	"log"
	"github.com/alexedwards/argon2id"
)

var params = &argon2id.Params{
	Memory:      64 * 1024,
	Iterations:  4,
	Parallelism: 1,
	SaltLength:  16,
	KeyLength:   32,
}

func EncodePassword(password string) (hash string) {
	hash, err := argon2id.CreateHash(password, params)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	return
}


func ComparePasswords(password string, hash string) (match bool) {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	
	if err != nil {
		log.Println(err)
		panic(err)
	}

	return 
}