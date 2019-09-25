package communicator

import (
	"fmt"
	"math"
)

type configItem interface {
	setVal (interface{}) error
	getVal () interface{}
}

type capturePacketRateConfig struct {
	name string
	tcpTPR float64
	mysqlTPR float64
}

func newCapturePacketRateConfig() (cprc *capturePacketRateConfig) {
	cprc = &capturePacketRateConfig{
		name:     CAPTURE_PACKET_RATE,
		tcpTPR:   1.0,
		mysqlTPR: 1.0,
	}
	return
}

func (cprc *capturePacketRateConfig) setVal (val interface{}) (err error){
	realVal, ok := val.(float64)
	if !ok {
		err = fmt.Errorf("cannot reansform val: %v to float64", val)
		return
	}

	cprc.mysqlTPR = realVal
	cprc.tcpTPR = math.Sqrt(realVal)
	return
}

func (cprc *capturePacketRateConfig) getVal () (val interface{}){
	return cprc.mysqlTPR
}