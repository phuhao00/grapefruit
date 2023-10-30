package adapter

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"grapefruit/config"
	"net/smtp"
	"strings"
)

func SendEmail(subject string, receiver string, content string) error {
	if config.SMTPFrom == "" { // for compatibility
		config.SMTPFrom = config.SMTPAccount
	}
	encodedSubject := fmt.Sprintf("=?UTF-8?B?%s?=", base64.StdEncoding.EncodeToString([]byte(subject)))
	mail := []byte(fmt.Sprintf("To: %s\r\n"+
		"From: %s<%s>\r\n"+
		"Subject: %s\r\n"+
		"Content-Type: text/html; charset=UTF-8\r\n\r\n%s\r\n",
		receiver, config.SystemName, config.SMTPFrom, encodedSubject, content))
	auth := smtp.PlainAuth("", config.SMTPAccount, config.SMTPToken, config.SMTPServer)
	addr := fmt.Sprintf("%s:%d", config.SMTPServer, config.SMTPPort)
	to := strings.Split(receiver, ";")
	var err error
	if config.SMTPPort == 465 {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         config.SMTPServer,
		}
		conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", config.SMTPServer, config.SMTPPort), tlsConfig)
		if err != nil {
			return err
		}
		client, err := smtp.NewClient(conn, config.SMTPServer)
		if err != nil {
			return err
		}
		defer client.Close()
		if err = client.Auth(auth); err != nil {
			return err
		}
		if err = client.Mail(config.SMTPFrom); err != nil {
			return err
		}
		receiverEmails := strings.Split(receiver, ";")
		for _, receiver := range receiverEmails {
			if err = client.Rcpt(receiver); err != nil {
				return err
			}
		}
		w, err := client.Data()
		if err != nil {
			return err
		}
		_, err = w.Write(mail)
		if err != nil {
			return err
		}
		err = w.Close()
		if err != nil {
			return err
		}
	} else {
		err = smtp.SendMail(addr, auth, config.SMTPAccount, to, mail)
	}
	return err
}
