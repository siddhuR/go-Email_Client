package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func sendMailSimple(subject string, body string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"routhusiddhartha@gmail.com",
		"optzbjabrctirahs",
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"routhusiddhartha@gmail.com",
		to,
		// []string{"routhusiddhartha@gmail.com"},
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}

func sendMailSimpleHTML(subject string, templatePath string, to []string) {

	// Get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string }{Name: "Siddhu"})

	if err != nil {
		fmt.Println(err)
		return
	}

	auth := smtp.PlainAuth(
		"",
		"routhusiddhartha@gmail.com",
		"optzbjabrctirahs",
		"smtp.gmail.com",
	)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	msg := "Subject: " + subject + "\n" + headers + "\n\n" + body.String()

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"routhusiddhartha@gmail.com",
		to,
		// []string{"routhusiddhartha@gmail.com"},
		[]byte(msg),
	)
	if err != nil {
		fmt.Println(err)
	}
}

func sendGomail(templatePath string) {
	//Get html
	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	t.Execute(&body, struct{ Name string }{Name: "Siddhu"})

	if err != nil {
		fmt.Println(err)
		return
	}

	// Send with gomail
	m := gomail.NewMessage()
	m.SetHeader("From", "routhusiddhartha@gmail.com")
	m.SetHeader("To", "routhusiddhartha@gmail.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", body.String())
	m.Attach("./siddhu.jpg")

	d := gomail.NewDialer("smtp.gmail.com", 587, "routhusiddhartha@gmail.com", "optzbjabrctirahs")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

/*
func sendSendgrid() {
	from := mail.NewEmail("Siddhu", "hello@siddhu.com")
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Siddhu", "routhusiddhartha@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

*/

func main() {
	// sendMailSimple(
	// 	"Another subject",
	// 	"Another body",
	// 	[]string{"routhusiddhartha@gmail.com"})

	// sendMailSimpleHTML(
	// 	"Another subject",
	// 	//"<h1>I'm a heading</h1><p>I'm a paragraph</p>",
	// 	"./test.html",
	// 	[]string{"routhusiddhartha@gmail.com"})

	sendGomail("./test.html")
	//sendSendgrid()

}
