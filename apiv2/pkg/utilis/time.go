package utilis

import "time"

func DBTime() string {
	return time.Now().In(time.FixedZone("GMT+3", 3*60*60)).Format("2 Jan 2006 15.04") + " GMT+3"
}
