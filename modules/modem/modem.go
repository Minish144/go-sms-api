package modem

import (
	"errors"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/tarm/serial"
)

type Modem struct {
	comport  string
	boudrate int
	instance *serial.Port
}

// TODO add checking if device is connected for windows
func New(comport string, boudrate int) (*Modem, error) {
	if _, err := os.Stat(comport); runtime.GOOS != "windows" && os.IsNotExist(err) {
		return nil, errors.New("comport was not found, try checking config.yaml or reconnecting your device")
	}

	m := &Modem{comport: comport, boudrate: boudrate}
	c := &serial.Config{Name: comport, Baud: boudrate, ReadTimeout: time.Second}
	var err error
	m.instance, err = serial.OpenPort(c)
	if err != nil {
		return nil, nil
	}

	return m, nil
}

func (m *Modem) Send(number string, message string) error {
	c := make(chan error)

	go func() {
		m.sendCommand("AT+CMGF=1\r", false)
		m.sendCommand("AT+CMGS=\""+number+"\"\r", false)
		_, err := m.sendCommand(message+string(rune(26)), true) // 26 = CTRL+Z
		if err != nil {
			c <- err
		}
		c <- nil
	}()

	select {
	case err := <-c:
		return err
	case <-time.After(viper.GetDuration("modem.sms_timeout") * time.Second):
		return errors.New("timeout exceed")
	}
}

/*
func (m *Modem) ReadAll() {
	m.sendCommand("AT+CMGF=1\r", false)
	x := m.sendCommand("AT+CMGL=\"ALL\"\r", true)
	log.Println("XXX ", x)
	tmp := parseMessage(x)
	log.Println("MESSAGE ", tmp)
}

type Message struct {
	Id    string
	Phone string
	Date  string
	Msg   string
}

func parseMessage(text string) []Message {
	var list []Message
	listLines := strings.Split(text, "\r\n")
	for i := 0; i < len(listLines)-3; i = i + 2 {
		tmp := strings.Split(listLines[i], ",")
		tmp[2] = strings.Replace(tmp[2], `"`, ``, -1)
		id := tmp[0][7:]
		phone := make([]byte, len(tmp[2]))
		hex.Decode(phone, []byte(tmp[2]))
		msg := make([]byte, len(listLines[i+1]))
		hex.Decode(msg, []byte(listLines[i+1]))
		list = append(list, Message{Id: id, Phone: string(phone), Date: tmp[4], Msg: string(msg)})
	}
	return list
}

func (m *Modem) Delete(id string) {
	m.sendCommand("AT+CMGF=1\r", false)
	x := m.sendCommand("AT+CMGD="+id+"\r", true)
	log.Println("MESSAGE ", x)
}
*/
func (m *Modem) sendCommand(message string, wait bool) (string, error) {
	m.instance.Flush()
	_, err := m.instance.Write([]byte(message))
	if err != nil {
		return "", err
	}
	buf := make([]byte, 1024)
	var loop int = 1
	if wait {
		loop = 10
	}
	var msg string
	var status string
	for i := 0; i < loop; i++ {
		n, _ := m.instance.Read(buf)
		if n > 0 {
			status = string(buf[:n])
			msg += status
			if strings.HasSuffix(status, "OK\r\n") || strings.HasSuffix(status, "ERROR\r\n") {
				break
			}
		}
	}

	return msg, nil
}
