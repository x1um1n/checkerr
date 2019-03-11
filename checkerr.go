// package providing error handlers I use in pretty much every app I write
package checkerr

import "log"

// check if an error is an error and log the supplied strings if so
func Check(e error, s ...string) bool {
    if e != nil {
        for _, str := range s {
            log.Println(str)
        }
        log.Println(e.Error())
        return true
    } else {
        return false
    }
}

// check if an error is an error, log the supplied strings and exit if so
func CheckFatal(e error, s ...string) bool {
    if e != nil {
        for _, str := range s {
            log.Println(str)
        }
        log.Fatalln(e.Error())
        return true
    } else {
        return false
    }
}
