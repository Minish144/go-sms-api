package server

import (
	"fmt"
	"net/http"

	"github.com/minish144/go-sms-api/modules/modem"
	"github.com/sirupsen/logrus"
)

func HandleSendSMS(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	if err := modem.SendSMS("+79653538013", "hello world"); err != nil {
		logrus.WithFields(
			logrus.Fields{
				"error": err.Error(),
			},
		).Errorln("failed to connect to modem")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
