package modem

import (
	"github.com/spf13/viper"
)

func SendSMS(mobile string, message string) error {
	comport := viper.GetString("modem.comport")
	baudrate := viper.GetInt("modem.baudrate")
	modem := New(comport, baudrate)
	if err := modem.Connect(); err != nil {
		return err
	}
	SendSMS("+79653538013", "hello world") // TODO handle errors
	return nil
}
