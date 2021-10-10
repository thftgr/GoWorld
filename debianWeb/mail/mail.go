package mail

import (
	"crypto/tls"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/mail"
	"net/smtp"
	"os"
	"strconv"
	"thftgr.com/GoWorld/debianWeb/mariadb"
	"thftgr.com/GoWorld/debianWeb/middlewareFunc"
	"thftgr.com/GoWorld/debianWeb/src"
	"time"
)

type Data struct {
	Userid        string
	VerifyKey     string
	VerifyTimeout string
	DisableKey    string
	Disable       bool
	Ip            string
	DeviceType    string
}

func readFormat(name string) (f string, err error) {
	file, err := os.ReadFile("./mail/"+name + ".mail")
	if err != nil {
		return
	}
	f = string(file)
	return

}

func genCode(r *rand.Rand, l int) (code string) {

	for i := 0; i < l; i++ {
		code += strconv.Itoa(r.Intn(10))
	}
	return

}

func sendMail(sendTo, messages string) (err error) {

	from := mail.Address{"", src.Config.Mail.From}
	to := mail.Address{"", sendTo}
	subj := "debian osu! Account authentication"
	body := messages

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := src.Config.Mail.Host

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", src.Config.Mail.Id, src.Config.Mail.Passwd, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
	return
}
func emailVerify(email, username string) (err error) {
	format, err := readFormat("register")
	if err != nil {
		return err
	}
	return sendMail(email, fmt.Sprintf(format, ""))

}

func SendCertificationMail(c echo.Context) (err error) {


	token := c.Get("JWT").(map[string]interface{})
	user := token["user"].(map[string]interface{})
	var (
		email    = user["email"].(string)
		userId   = user["userID"]
		username = user["username"].(string)
		vc       string
		dc       string
	)
	if email == "" {
		return middlewareFunc.RequestIDToJson(http.StatusBadRequest, c, "email address null")
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	vc = genCode(r, 6)
	dc = genCode(r, 6)

	if err = mariadb.InsertVerifyCode(userId, vc, dc); err != nil {
		middlewareFunc.RequestIDToJson(http.StatusInternalServerError, c, "")
		return err
	}

	format, err := readFormat("certification")
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusInternalServerError, c, "")
		return err
	}
	err = sendMail(email, fmt.Sprintf(format,
		username, c.RealIP(), vc, userId, vc, userId, dc,
		username, c.RealIP(), vc, userId, vc, userId, dc,
	))
	if err != nil {
		return middlewareFunc.RequestIDToJson(http.StatusInternalServerError, c, "")
	}
	err = mariadb.SetAccountStatus(userId, "email verify pending")
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusInternalServerError, c, "")
		return err
	}
	//
	return

}
