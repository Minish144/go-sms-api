package server

import (
	"encoding/json"
	"net/http"

	"github.com/minish144/go-sms-api/modules/modem"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type sms struct {
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func HandleSendSMS(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	newSms := sms{}
	if err := decoder.Decode(&newSms); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	comport := viper.GetString("modem.comport")
	baudrate := viper.GetInt("modem.baudrate")
	newModem := modem.New(comport, baudrate)

	// TODO add timer
	if err := newModem.Send(newSms.Phone, newSms.Message); err != nil {
		logrus.WithFields(
			logrus.Fields{
				"error": err.Error(),
			},
		).Errorln("failed to send sms")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	logrus.WithFields(
		logrus.Fields{
			"phone": newSms.Phone,
			"msg":   newSms.Message,
		},
	).Infoln("sms was sent successfully")
}
