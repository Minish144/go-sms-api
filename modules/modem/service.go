package modem

import (
	"github.com/spf13/viper"
)

func SendSMS(mobile string, message string) error {
	comport := viper.GetString("modem.comport")
	baudrate := viper.GetInt("modem.baudrate")
	modem := New(comport, baudrate)
	return modem.Send("+79653538013", "hello world") // TODO handle errors
}
