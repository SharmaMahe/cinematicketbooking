package mail

import (
    "log" 
    "net/smtp"
    "github.com/astaxie/beego"
)

// Send Mail
func SendMail(body string, to string, subject string) {
    from := beego.AppConfig.String("smtpusername")
    pass := beego.AppConfig.String("smtppassword")
    smtpHost := beego.AppConfig.String("smtphost")
    smtpPort := beego.AppConfig.String("smptport")

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: "+ subject + "\n\n" +
        body

    err := smtp.SendMail(smtpHost + ":" + smtpPort,
        smtp.PlainAuth("", from, pass, smtpHost),
        from, []string{to}, []byte(msg))

    if err != nil {
        log.Printf("could not connect with smtp: %s", err)
        return
    }
}

// Booking confirmation
func SendBookingMail(name string, email string, seats []string) {
    var seat string

    for _, s := range seats {
        seat = seat + "\n" + s
    }

    subject := "Ticket Confirmation"

    body := "Hi " + name + "," + "\n\n" +
        "You have successfully booked your tickets" + "\n\n" +
        "Booked seats:\n" + seat + "\n\n" +
        "Thanks & Regards"

    SendMail(body, email, subject)
}

// Cancel Booking Confirmation
func SendBookingCancelMail(name string, email string, seats []string) {
    var seat string

    msg := "seats"

    switch len(seats) {
        case 1:
            msg = "seat"
    }

    for _, s := range seats {
        seat = seat + "\n" + s
    }

    subject := "Booking Cancellation Confirmation"

    body := "Hi " + name + "," + "\n\n" +
        "You have successfully cancelled your "+ msg + "\n\n" +
        "Cancelled "+msg+":\n" + seat + "\n\n" +
        "Thanks & Regards"

    SendMail(body, email, subject)
}

