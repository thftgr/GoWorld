package Route

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"thftgr.com/GoWorld/debianWeb/mariadb"
	"thftgr.com/GoWorld/debianWeb/middlewareFunc"
)

func AccountDisable(c echo.Context)(err error){

	// 비활성 성공 html 페이지
	//code := c.QueryParam("code")
	return nil
}

func AccountVerify(c echo.Context)(err error){
	err = mariadb.CheckVerifyCode(c.Param("userid"), c.QueryParam("code"))
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusUnauthorized, c, "Account Verify fail. timeout or not matched code")
		return err
	}
	err = mariadb.SetAccountStatus(c.Param("userid"), "email verify success")
	if err != nil {
		middlewareFunc.RequestIDToJson(http.StatusInternalServerError, c, "")
		return err
	}
	return c.String(http.StatusOK,"Account Verify success")
}
//func MailCertification(c echo.Context) (err error) {
//	fmt.Println(c.QueryParam("code"))
//	if mail.CertificationCodeList[c.QueryParam("code")].Timeout > time.Now().Second(){
//		delete(mail.CertificationCodeList,c.QueryParam("code"))
//	}
//	return c.NoContent(200)
//}


