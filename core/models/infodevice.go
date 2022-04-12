package models

type InfoDeviceClient struct {
	Mac				string
	DeviceName		string
}

type DeviceInfo struct {
	Status			string
	Intensity		float32
}

type InfoDeviceServer struct {
	ResponseCode	int64
	ResponseString	string
	RoomName		string
	Description		string
	Info			map[string]DeviceInfo
}