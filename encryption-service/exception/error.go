package exception

import "log"

func PanicIfNeeded(err interface{}) {
	if err != nil {
		log.Println(err)
	}
}
