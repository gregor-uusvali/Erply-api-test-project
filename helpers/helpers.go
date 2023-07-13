package helpers

import (
	"Erply-api-test-project/api"
	"Erply-api-test-project/models"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var Session = models.Session{}

func ClientError(w http.ResponseWriter, status int) {
	fmt.Println(http.StatusText(status))
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	fmt.Println(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func CheckForSession() bool {
	if Session.SessionKey != "" {
		skUserInfo, _ := api.GetSessionKeyInfo(Session.User.ClientCode, Session.SessionKey)
		i64, err := strconv.ParseInt(skUserInfo.Records[0].ExpireUnixTime, 10, 64)
		if err != nil {
			panic(err)
		}
		if i64-time.Now().Unix() > 0 {
			return true
		}
	}
	return false
}
