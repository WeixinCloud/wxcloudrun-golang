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

// JsonResult 返回结构
type JsonResult struct {
	Code     int         `json:"code"`
	ErrorMsg string      `json:"errorMsg"`
	Data     interface{} `json:"data,omitempty"`
}

// RouteHandler Restful路由
func RouteHandler(w http.ResponseWriter, r *http.Request) {
	rootUri := "/api/todos"
	regex1, _ := regexp.Compile(fmt.Sprintf("^(%s)(/)?$", rootUri))
	regex2, _ := regexp.Compile(fmt.Sprintf("^(%s)/(\\d)+(/)?$", rootUri))
	path := r.URL.Path
	res := &JsonResult{}

	if regex1.MatchString(path) {
		res = MatchUriNoID(r)
	} else if regex2.MatchString(path) {
		res = MatchUriWithID(r)
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

// MatchUriNoID
// uri: /api/todos
// 方法: Get、Post、Put
func MatchUriNoID(r *http.Request) *JsonResult {
	res := &JsonResult{}
	if r.Method == http.MethodGet {
		toDoList, err := GetToDoList(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = toDoList
		}
	} else if r.Method == http.MethodPost {
		err := AddToDoItem(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		}
	} else if r.Method == http.MethodPut {
		err := UpdateToDoItem(r)
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

// MatchUriNoID
// uri: /api/todos/:id
// 方法: Get、Delete
func MatchUriWithID(r *http.Request) *JsonResult {
	res := &JsonResult{}
	if r.Method == http.MethodGet {
		toDoItem, err := QueryToDoItem(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = toDoItem
		}
	} else if r.Method == http.MethodDelete {
		err := DeleteToDoItem(r)
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

// GetToDoList 获取所有的todo项
func GetToDoList(r *http.Request) ([]*model.ToDoItemModel, error) {
	log.Print("received a GetToDoList request.")
	if r.Method != http.MethodGet {
		return nil, fmt.Errorf("GetToDoList only support GET method")
	}

	toDoList, err := dao.Imp.GetToDoList()
	if err != nil {
		return nil, err
	}

	return toDoList, nil
}

// AddToDoItem 新增todo项
func AddToDoItem(r *http.Request) error {
	log.Print("received a AddToDoItem request.")
	if r.Method != http.MethodPost {
		return fmt.Errorf("AddToDoItem only support POST method")
	}

	toDoItem, err := parseAndCheckBody(r)
	if err != nil {
		return err
	}

	now := time.Now()
	toDoItem.CreateTime = now
	toDoItem.UpdateTime = now

	err = dao.Imp.AddToDoItem(toDoItem)
	if err != nil {
		return err
	}

	return nil
}

// DeleteToDoItem 删除todo项
func DeleteToDoItem(r *http.Request) error {
	log.Print("received a DeleteToDoItem request.")
	if r.Method != http.MethodDelete {
		return fmt.Errorf("DeleteToDoItem only support DELETE method")
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/todos/")
	var err error
	var intId int
	intId, err = strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = dao.Imp.DeleteToDoItemById(int32(intId))
	if err != nil {
		return err
	}

	return nil
}

// UpdateToDoItem 更新todo项
func UpdateToDoItem(r *http.Request) error {
	log.Print("received a UpdateToDoItem request.")
	if r.Method != http.MethodPut {
		return fmt.Errorf("UpdateToDoItem only support PUT method")
	}

	toDoItem, err := parseAndCheckBody(r)
	if err != nil {
		return err
	}

	if toDoItem.Id <= 0 {
		return fmt.Errorf("id[%d] not exist", toDoItem.Id)
	}

	now := time.Now()
	toDoItem.UpdateTime = now

	err = dao.Imp.UpdateToDoItemById(toDoItem.Id, toDoItem)
	if err != nil {
		return err
	}

	return nil
}

// QueryToDoItem 查询单个todo项
func QueryToDoItem(r *http.Request) (*model.ToDoItemModel, error) {
	log.Print("received a QueryToDoItem request.")
	if r.Method != http.MethodGet {
		return nil, fmt.Errorf("QueryToDoItem only support GET method")
	}

	id := strings.TrimPrefix(r.URL.Path, "/api/todos/")
	var err error
	var intId int
	intId, err = strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	toDoItem, err := dao.Imp.QueryToDoItemById(int32(intId))
	if err != nil {
		return nil, err
	}

	return toDoItem, nil
}

func parseAndCheckBody(r *http.Request) (*model.ToDoItemModel, error) {
	decoder := json.NewDecoder(r.Body)
	body := make(map[string]interface{}, 0)
	if err := decoder.Decode(&body); err != nil {
		return nil, err
	}
	defer r.Body.Close()

	toDoItem := new(model.ToDoItemModel)
	var ok bool
	var data interface{}
	if data, ok = body["title"]; ok {
		if toDoItem.Title, ok = data.(string); !ok {
			return nil, fmt.Errorf("name need string type but %+v", reflect.TypeOf(data))
		}
	}
	if data, ok = body["status"]; ok {
		if toDoItem.Status, ok = data.(string); !ok {
			return nil, fmt.Errorf("email need string type but %+v", reflect.TypeOf(data))
		}
	}
	if data, ok = body["id"]; ok {
		var id float64
		if id, ok = data.(float64); !ok {
			return nil, fmt.Errorf("id need int type but %+v", reflect.TypeOf(data))
		}
		toDoItem.Id = int32(id)
	}
	return toDoItem, nil
}
