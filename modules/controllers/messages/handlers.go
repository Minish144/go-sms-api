package messages

import (
	"context"

	"github.com/minish144/go-sms-api/gen/pb"
	"github.com/minish144/go-sms-api/modules/modem"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Send(ctx context.Context, in *pb.Messages_SendRequest) (*pb.Messages_SendResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	comport := viper.GetString("modem.comport")
	baudrate := viper.GetInt("modem.baudrate")
	newModem, err := modem.New(comport, baudrate)
	if err != nil {
		logrus.WithFields(
			logrus.Fields{
				"error": err.Error(),
			},
		).Errorln("failed to initialize modem")
		return nil, err
	}

	if err := newModem.Send(in.Message.Phone, in.Message.Message); err != nil {
		logrus.WithFields(
			logrus.Fields{
				"error": err.Error(),
			},
		).Errorln("failed to send sms")
		return nil, err
	}

	logrus.WithFields(
		logrus.Fields{
			"phone":   in.Message.Phone,
			"message": in.Message.Message,
		},
	).Infoln("sms was sent successfully")

	return &pb.Messages_SendResponse{}, nil
}
