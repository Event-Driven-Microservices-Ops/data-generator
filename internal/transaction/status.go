package transaction

import "math/rand"

type Status string

const (
	PENDING  Status = "PENDING"
	SUCCESS  Status = "SUCCESS"
	RUNNING  Status = "RUNNING"
	FAILED   Status = "FAILED"
	FINISHED Status = "FINISHED"
)

var allStatuses = []Status{PENDING, SUCCESS, RUNNING, FAILED, FINISHED}

func RandomStatus() Status {
	return allStatuses[rand.Intn(len(allStatuses))]
}
