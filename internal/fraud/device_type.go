package fraud

import "math/rand"

type DeviceType string

const (
	DESKTOP DeviceType = "DESKTOP"
	MOBILE  DeviceType = "MOBILE"
	TABLET  DeviceType = "TABLET"
)

var allDevices = []DeviceType{DESKTOP, MOBILE, TABLET}

func RandomDevice() DeviceType {
	return allDevices[rand.Intn(len(allDevices))]
}
