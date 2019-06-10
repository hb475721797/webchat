package controller

import (
	"webchat/app/model"
	"webchat/util"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"webchat/app/format"
	"webchat/app/service"
	)

var contactService service.ContactService
func AddFriend(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	contact := &model.Contact{}
	err := util.Bind(r, contact)
	if err != nil {
		format.Fail(w, err.Error())
		return
	}

	err = contactService.AddFriend(contact.UserId, contact.AddId)
	if err != nil {
		format.Fail(w, err.Error())
		return
	}
	format.Success(w,nil, "添加好友成功")
	return
}

func LoadFriend(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	contact := &model.Contact{}
	err := util.Bind(r, contact)
	if err != nil {
		format.Fail(w, err.Error())
		return
	}
	users := contactService.LoadFriend(contact.UserId)
	format.SuccessList(w,users, len(users))
	return
}

func LoadGroup(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	contact := &model.Contact{}
	err := util.Bind(r, contact)
	if err != nil {
		format.Fail(w, err.Error())
		return
	}
	groups := contactService.LoadGroup(contact.UserId)
	format.SuccessList(w,groups, len(groups))
	return
}
