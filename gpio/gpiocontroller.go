package gpio

import (
	"fmt"
	"sync"
	"time"
	rpi "github.com/stianeikeland/go-rpio/v4"
)

type GPIOController struct {
	GPIOMap			map[string]GPIOState
	Info			GPIOInfo
	Status			bool
	mutex			sync.RWMutex
}

type GPIOState struct{
	Intensity		float32
	State			string
	IsRunning		bool
}

func (self *GPIOController) Process(gpio string) {
	dutytime := 20000
	
	ok, idgpio := self.Info.GetPin(gpio)

	if !ok {
		fmt.Println("No se puede crear instancia para " + gpio + " - Desconocido")
		return
	}

	pinout := rpi.Pin(idgpio)
	pinout.Mode(rpi.Output)

	for {
		state, intensity := self.GetState(gpio)
		
		if state == "Off" {
			pinout.Write(rpi.Low)
			time.Sleep(time.Duration(dutytime) * time.Microsecond)
		} else {
			ontime := int(float32(dutytime) * intensity)
			offtime := dutytime - ontime
			
			pinout.Write(rpi.High)
			time.Sleep(time.Duration(ontime) * time.Microsecond)
			pinout.Write(rpi.Low)
			time.Sleep(time.Duration(offtime) * time.Microsecond)
		}
	}
}

func (self *GPIOController) Initialize() bool {
	fmt.Println("Iniciando GPIO Controller")
	
	err := rpi.Open()

	if err != nil {
		fmt.Println("Hubo un error al inicializar el GPIO - " + err.Error())
		self.Status = false
		return false
	}

	self.Status = true

	self.GPIOMap = make(map[string]GPIOState)
	self.Info.InitializeGPIOInfo()

	fmt.Println("Controlador GPIO iniciado correctamente")

	return true
}

func (self *GPIOController) GetState(key string) (string, float32) {
	var state string
	var intensity float32

	state = "Off"

	self.mutex.RLock()
	value,ok := self.GPIOMap[key]

	if ok {
		state = value.State
		intensity = value.Intensity
	}

	self.mutex.RUnlock()

	return state, intensity
}

func (self *GPIOController) SetState(gpio string, state string, intensity float32){
	self.mutex.Lock()
	value,ok := self.GPIOMap[gpio]

	if !ok {
		//fmt.Println("Creando controlador de PWM para el pin " + gpio)


		self.GPIOMap[gpio] = GPIOState{
			Intensity: intensity,
			State: state,
			IsRunning : true,
		}
		go self.Process(gpio)
	} else{
		//fmt.Println("Actualizando controlador de PWM " + gpio + " " + state)
		value.State = state
		value.Intensity = intensity
		self.GPIOMap[gpio] = value
	}

	self.mutex.Unlock()
}