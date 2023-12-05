package email

import (
	"errors"
	"github.com/RianIhsan/go-code-nexus/config"
	"github.com/wneessen/go-mail"
	"html/template"
	"path/filepath"
	"runtime"
	"strings"
)

func EmailService(username, email, token string) error {
	hostEmail := config.BootConfig().EmailConfig.Host
	senderEmail := config.BootConfig().EmailConfig.Sender
	senderPass := config.BootConfig().EmailConfig.Password
	senderMailPort := config.BootConfig().EmailConfig.MailPort

	m := mail.NewMsg()
	if err := m.From(senderEmail); err != nil {
		return err
	}
	if err := m.To(email); err != nil {
		return err
	}

	m.Subject("Important: Your Token for CodeNexus Account")
	emailTemplate := struct {
		Token    string
		Username string
	}{
		Token:    token,
		Username: username,
	}
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return errors.New("Failed to get the current file path")
	}
	templatePath := filepath.Join(filepath.Dir(filename), "email_template.html")
	tmpl, err := template.New("emailTemplate").ParseFiles(templatePath)
	if err != nil {
		return err
	}
	var bodyContent strings.Builder
	if err := tmpl.ExecuteTemplate(&bodyContent, "email_template.html", emailTemplate); err != nil {
		return err
	}
	m.SetBodyString(mail.TypeTextHTML, bodyContent.String())

	c, err := mail.NewClient(hostEmail, mail.WithPort(senderMailPort), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(senderEmail), mail.WithPassword(senderPass))
	if err != nil {
		return err
	}
	if err := c.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
