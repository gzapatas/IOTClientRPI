package gpio

import (
	"strconv"
)

type GPIOData struct{
	Id		int16
	Pin		int16
}

type GPIOInfo struct{
	GPIO	map[string]GPIOData
}

func (self *GPIOInfo) InitializeGPIOInfo(){
	self.GPIO = map[string]GPIOData{}
	ids := []int16{0,1,2,3,4,5,6,7,8,9,10,11,12,13,16,17,18,19,20,21,22,23,24,25,26,27}
	pinouts := []int16{27,28,3,5,7,29,31,26,24,21,19,23,32,33,36,11,12,35,38,40,15,16,18,22,37,13}


	for i,idvalue:= range ids {
		pinvalue := pinouts[i]

		obj := GPIOData{Id:idvalue,Pin:pinvalue}
		key := "GPIO" + strconv.FormatInt(int64(idvalue),10)

		self.GPIO[key] = obj
	}
}

func (self *GPIOInfo) GetPin(gpioname string) (bool,int16){
	value,exists := self.GPIO[gpioname]

	if !exists {
		return false,0
	}

	return true, value.Id
}