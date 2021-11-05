package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg"`
	Data     interface{} `json:"data,omitempty"`
}

// RouteHandler Restful路由
func RouteHandler(w http.ResponseWriter, r *http.Request) {
	regex1, _ := regexp.Compile("^(/user)(/)?$")
	regex2, _ := regexp.Compile("^(/user)/(\\d)+(/)?$")
	path := r.URL.Path
	res := &JsonResult{}

	if regex1.MatchString(path) {
		res = AddOrUpdateUser(w, r)
	} else if regex2.MatchString(path) {
		res = QueryOrDelete(w, r)
	} else {
		res.Code = -1
		res.ErrorMsg = "url not match ant handler"
	}

	msg, err := json.Marshal(res)
	if err != nil {
		fmt.Fprint(w, "sever internel error")
		return
	}
	w.Write(msg)
}

// AddOrUpdateUser 新增或者更新用户
func AddOrUpdateUser(w http.ResponseWriter, r *http.Request) *JsonResult {
	res := &JsonResult{}
	if r.Method == http.MethodPost {
		user, err := AddUser(w, r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = user
		}
	} else if r.Method == http.MethodPut {
		err := UpdateUser(w, r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		}
	} else {
		res.Code = -1
		res.ErrorMsg = "Request method error"
	}
	return res
}

// QueryOrDelete 查询或者删除用户
func QueryOrDelete(w http.ResponseWriter, r *http.Request) *JsonResult {
	res := &JsonResult{}
	if r.Method == http.MethodGet {
		user, err := QueryUser(w, r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = user
		}
	} else if r.Method == http.MethodDelete {
		err := DeleteUser(w, r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		}
	} else {
		res.Code = -1
		res.ErrorMsg = "Request method error"
	}
	return res
}

// AddUser 新增用户
func AddUser(w http.ResponseWriter, r *http.Request) (*model.UserModel, error) {
	log.Print("received a AddUser request.")
	if r.Method != http.MethodPost {
		return nil, fmt.Errorf("AddUser only support POST method")
	}

	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{}, 0)
	if err := decoder.Decode(&body); err != nil {
		return nil, err
	}
	defer r.Body.Close()

	user := new(model.UserModel)
	var ok bool
	var data interface{}
	if data, ok = body["name"]; ok {
		if user.Name, ok = data.(string); !ok {
			return nil, fmt.Errorf("name need string type but %+v", reflect.TypeOf(data))
		}
	}
	if data, ok = body["email"]; ok {
		if user.Email, ok = data.(string); !ok {
			return nil, fmt.Errorf("email need string type but %+v", reflect.TypeOf(data))
		}
	}
	if data, ok = body["phone"]; ok {
		if user.Phone, ok = data.(string); !ok {
			return nil, fmt.Errorf("phone need string type but %+v", reflect.TypeOf(data))
		}
	}
	if data, ok = body["description"]; ok {
		if user.Description, ok = data.(string); !ok {
			return nil, fmt.Errorf("description need string type but %+v", reflect.TypeOf(data))
		}
	}
	if data, ok = body["age"]; ok {
		var age float64
		if age, ok = data.(float64); !ok {
			return nil, fmt.Errorf("age need int type but %+v", reflect.TypeOf(data))
		}
		user.Age = int32(age)
	}

	now := time.Now()
	user.CreateTime = now
	user.UpdateTime = now

	user, err := dao.Imp.AddUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser 删除用户
func DeleteUser(w http.ResponseWriter, r *http.Request) error {
	log.Print("received a DeleteUser request.")
	if r.Method != http.MethodDelete {
		return fmt.Errorf("DeleteUser only support DELETE method")
	}

	id := strings.TrimPrefix(r.URL.Path, "/user/")
	var err error
	var intId int
	intId, err = strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = dao.Imp.DeleteUserById(int32(intId))
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser 更新用户
func UpdateUser(w http.ResponseWriter, r *http.Request) error {
	log.Print("received a UpdateUser request.")
	if r.Method != http.MethodPut {
		return fmt.Errorf("UpdateUser only support PUT method")
	}

	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{}, 0)
	if err := decoder.Decode(&body); err != nil {
		return err
	}
	defer r.Body.Close()

	user := new(model.UserModel)
	user.Id = -1
	var ok bool
	var data interface{}
	if data, ok = body["name"]; ok {
		if user.Name, ok = data.(string); !ok {
			return fmt.Errorf("name need string type but %+v", reflect.TypeOf(data))
		}
	}
	if data, ok = body["email"]; ok {
		if user.Email, ok = data.(string); !ok {
			return fmt.Errorf("email need string type but %+v", reflect.TypeOf(data))
		}
	}
	if data, ok = body["phone"]; ok {
		if user.Phone, ok = data.(string); !ok {
			return fmt.Errorf("phone need string type but %+v", reflect.TypeOf(data))
		}
	}
	if data, ok = body["description"]; ok {
		if user.Description, ok = data.(string); !ok {
			return fmt.Errorf("description need string type but %+v", reflect.TypeOf(data))
		}
	}
	if data, ok = body["age"]; ok {
		var age float64
		if age, ok = data.(float64); !ok {
			return fmt.Errorf("age need int type but %+v", reflect.TypeOf(data))
		}
		user.Age = int32(age)
	}
	if data, ok = body["id"]; ok {
		var id float64
		if id, ok = data.(float64); !ok {
			return fmt.Errorf("id need int type but %+v", reflect.TypeOf(data))
		}
		user.Id = int32(id)
	}

	if user.Id < 0 {
		return fmt.Errorf("id[%d] not exist", user.Id)
	}

	now := time.Now()
	user.UpdateTime = now

	err := dao.Imp.UpdateUserById(user.Id, user)
	if err != nil {
		return err
	}

	return nil
}

// QueryUser 查询用户
func QueryUser(w http.ResponseWriter, r *http.Request) (*model.UserModel, error) {
	log.Print("received a QueryUser request.")
	if r.Method != http.MethodGet {
		return nil, fmt.Errorf("QueryUser only support GET method")
	}

	id := strings.TrimPrefix(r.URL.Path, "/user/")
	var err error
	var intId int
	intId, err = strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	user, err := dao.Imp.QueryUserById(int32(intId))
	if err != nil {
		return nil, err
	}

	return user, nil
}
