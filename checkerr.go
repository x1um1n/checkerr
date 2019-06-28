// Package checkerr provides error handlers I use in pretty much every app I write
package checkerr

import (
	"log"
	"net/http"
)

// Check if an error is an error and log the supplied strings if so
func Check(e error, s ...string) bool {
	if e != nil {
		for _, str := range s {
			log.Println(str)
		}
		log.Println(e.Error())
		return true
	}
	return false
}

// CheckFatal checks if an error is an error, log the supplied strings and exit if so
func CheckFatal(e error, s ...string) bool {
	if e != nil {
		for _, str := range s {
			log.Println(str)
		}
		log.Fatalln(e.Error())
		return true
	}
	return false
}

// Check401 checks if an error is an error and throws a 401 Unauthorised if so
// this is intended to be used when checking authorisation headers
func Check401(e error, w http.ResponseWriter, s ...string) bool {
	if e != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Println("ERROR 401 Unauthorized")
		for _, str := range s {
			log.Println(str)
		}
		log.Println(e.Error())
		return true
	}
	return false
}

// Check500 checks if an error is an error and throws a 500 error if so
func Check500(e error, w http.ResponseWriter, s ...string) bool {
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERROR 500 Internal Server Error")
		for _, str := range s {
			log.Println(str)
		}
		log.Println(e.Error())
		return true
	}
	return false
}
