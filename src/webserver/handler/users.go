package handler

import (
	"fmt"
	"html/template"
	"net/http"

	us "github.com/asepnur/iskandar/src/module/users"
	t "github.com/asepnur/iskandar/src/webserver/template"
	"github.com/julienschmidt/httprouter"
)

// User struct :: save value
type User struct {
	UserID     int    `json:"id"`
	UserEmail  string `json:"email"`
	FullName   string `json:"name"`
	MSISDN     int    `json:"msisdn"`
	CreateTime string `json:"create_time"`
}

var (
	emptyTime  = "0001-01-01 00:00:00"
	layoutTime = "2006-01-02 15:04:05"
)

// ViewHTML ::
func ViewHTML(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tpl, err := template.ParseFiles("src/webserver/template/users.html")
	el, err := us.GetVisitor()
	us.IncreaseVisitor(fmt.Sprintf("%d", el))
	data := map[string]interface{}{
		"Visitor": el,
	}
	if err != nil {
		fmt.Println(err)
	}
	tpl.Execute(w, data)
	return
}

// SelectUserHandler ..
func SelectUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var users []us.User
	var err error

	q := r.FormValue("name")
	resp := []User{}

	if q != "" {
		users, err = us.GetMultipleByFilter(q)
	} else {
		users, err = us.GetMultipleUser()
	}
	if err != nil {
		t.RenderJSONResponse(w, new(t.Response).
			SetCode(http.StatusInternalServerError))
		return
	}
	for _, el := range users {
		ct := el.CreateTime.Format(layoutTime)
		if ct == emptyTime {
			ct = "-"
		}
		resp = append(resp, User{
			UserID:     el.UserID,
			UserEmail:  el.UserEmail,
			FullName:   el.FullName,
			MSISDN:     el.MSISDN,
			CreateTime: ct,
		})
	}
	if err != nil {
		fmt.Println(err)
	}
	t.RenderJSONResponse(w, new(t.Response).
		SetCode(http.StatusOK).
		SetData(resp))
	return

}
