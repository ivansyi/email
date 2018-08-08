package main

import (
    "github.com/ivansyi/email"
    "net/smtp"
    "strings"
    "log"
    "flag"
    "fmt"
    "os"
	)

var from = "*************此处写发件人地址********************"
var p1 = ""
var p2 = ""

func main() {
    to := flag.String("to", "", "email address you want to send mail to")
    subject := flag.String("subject", "", "email subject")
    msg := flag.String("msg", "", "email content")
    file := flag.String("file", "", "the (single)file to be attached")
    mserver := flag.String("mserver", "smtp.exmail.qq.com", "smtp server(say: smtp.exmail.qq.com)")
    from := flag.String("from", "yishunli@thunder.com.cn", "the mail account used to send this email")
    pwd := flag.String("pwd", "", "the password for 'from' account")
    flag.Parse()
    //fmt.Println("To: ", *to, "\nSubject: ", *subject, "\nMessage: ", *msg,"\nFile: ", *file, "\n")
    if len(*to) == 0 {
        fmt.Println("Invalid mail address")
        os.Exit(1)
    }
    var passwd string
    if len(*subject) == 0 {
        fmt.Println("We do need a non-empty \"subject\"")
        os.Exit(1)
    }

    if len(*pwd) == 0{
        //fmt.Println("Using empty password!")
        passwd = "*************此处写发件人邮箱密码(默认，可不写），不填写时需要在调用此程序时通过参数指出********************"
    }else{
        passwd = *pwd
    }

    if len(*from) == 0 || len(*mserver) == 0 || len(passwd) == 0{
        fmt.Println("Invalid mail server authentication info(mserver/from/passwd)")
        os.Exit(1)
    }
    
    /*if len(*file) == 0 {
        fmt.Println("Fogot to specify your attachment file path?")
        return
    }*/

    m := email.NewMessage(*subject, *msg)
    m.From = *from
    //m.To = []string{*to}
    m.To = strings.Split(*to, ",")

    if len(*file) > 0 {
        err := m.Attach(*file)
        if err != nil {
            log.Println(err)
        }
    }

    server := *mserver + ":25"
    err := email.Send(server, smtp.PlainAuth("", *from, passwd, *mserver), m)
    if err != nil {
        log.Println(err)
        os.Exit(2)
    } else{
        fmt.Printf("Successed send the email:\n  To: \t\t%s\n  Subject: \t%s\n  Message: \t%s\n  File: \t%s\n", *to, *subject, *msg, *file)
        os.Exit(0)
    }
}
