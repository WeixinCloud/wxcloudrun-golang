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

const html = `<!doctype html>
<html lang="en">
  
  <head>
    <meta charset="utf-8" />
    <link rel="icon" href="https://cloudbase-run-todolist-92bb28a0d-1258016615.tcloudbaseapp.com/favicon.ico" />
    <meta name="viewport" content="width=device-width,initial-scale=1" />
    <meta name="theme-color" content="#000000" />
    <meta name="description" content="Web site created using create-react-app" />
    <link rel="apple-touch-icon"
	 href="https://cloudbase-run-todolist-92bb28a0d-1258016615.tcloudbaseapp.com/logo192.png" />
    <link rel="manifest" href="https://cloudbase-run-todolist-92bb28a0d-1258016615.tcloudbaseapp.com/manifest.json" />
    <title>Todo List</title>
    <link href="https://cloudbase-run-todolist-92bb28a0d-1258016615.tcloudbaseapp.com/static/css/2.20aa2d7b.chunk.css"
	 rel="stylesheet">
    <link
	 href="https://cloudbase-run-todolist-92bb28a0d-1258016615.tcloudbaseapp.com/static/css/main.d8680f04.chunk.css"
	 rel="stylesheet">
   </head>
  
  <body>
    <noscript>You need to enable JavaScript to run this app.</noscript>
    <div id="root"></div>
    <script>!
      function(e) {
        function t(t) {
          for (var n, l, a = t[0], p = t[1], i = t[2], f = 0, s = []; f < a.length; f++) l = a[f],
          Object.prototype.hasOwnProperty.call(o, l) && o[l] && s.push(o[l][0]),
          o[l] = 0;
          for (n in p) Object.prototype.hasOwnProperty.call(p, n) && (e[n] = p[n]);
          for (c && c(t); s.length;) s.shift()();
          return u.push.apply(u, i || []),
          r()
        }
        function r() {
          for (var e, t = 0; t < u.length; t++) {
            for (var r = u[t], n = !0, a = 1; a < r.length; a++) {
              var p = r[a];
              0 !== o[p] && (n = !1)
            }
            n && (u.splice(t--, 1), e = l(l.s = r[0]))
          }
          return e
        }
        var n = {},
        o = {
          1 : 0
        },
        u = [];
        function l(t) {
          if (n[t]) return n[t].exports;
          var r = n[t] = {
            i: t,
            l: !1,
            exports: {}
          };
          return e[t].call(r.exports, r, r.exports, l),
          r.l = !0,
          r.exports
        }
        l.m = e,
        l.c = n,
        l.d = function(e, t, r) {
          l.o(e, t) || Object.defineProperty(e, t, {
            enumerable: !0,
            get: r
          })
        },
        l.r = function(e) {
          "undefined" != typeof Symbol && Symbol.toStringTag && Object.defineProperty(e, Symbol.toStringTag, {
            value: "Module"
          }),
          Object.defineProperty(e, "__esModule", {
            value: !0
          })
        },
        l.t = function(e, t) {
          if (1 & t && (e = l(e)), 8 & t) return e;
          if (4 & t && "object" == typeof e && e && e.__esModule) return e;
          var r = Object.create(null);
          if (l.r(r), Object.defineProperty(r, "default", {
            enumerable: !0,
            value: e
          }), 2 & t && "string" != typeof e) for (var n in e) l.d(r, n,
          function(t) {
            return e[t]
          }.bind(null, n));
          return r
        },
        l.n = function(e) {
          var t = e && e.__esModule ?
          function() {
            return e.
          default
          }:
          function() {
            return e
          };
          return l.d(t, "a", t),
          t
        },
        l.o = function(e, t) {
          return Object.prototype.hasOwnProperty.call(e, t)
        },
        l.p = "https://cloudbase-run-todolist-92bb28a0d-1258016615.tcloudbaseapp.com/";
        var a = this.webpackJsonptodo = this.webpackJsonptodo || [],
        p = a.push.bind(a);
        a.push = t,
        a = a.slice();
        for (var i = 0; i < a.length; i++) t(a[i]);
        var c = p;
        r()
      } ([])</script>
    <script src="https://cloudbase-run-todolist-92bb28a0d-1258016615.tcloudbaseapp.com/static/js/2.18b41bed.chunk.js">
	</script>
    <script
	 src="https://cloudbase-run-todolist-92bb28a0d-1258016615.tcloudbaseapp.com/static/js/main.bde3e603.chunk.js">
	</script>
  </body>

</html>`

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
		fmt.Fprint(w, html)
		return
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
		todo, err := AddToDoItem(r)
		if err != nil {
			res.Code = -1
			res.ErrorMsg = err.Error()
		} else {
			res.Data = todo
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
func AddToDoItem(r *http.Request) (*model.ToDoItemModel, error) {
	log.Print("received a AddToDoItem request.")
	if r.Method != http.MethodPost {
		return nil, fmt.Errorf("AddToDoItem only support POST method")
	}

	toDoItem, err := parseAndCheckBody(r)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	toDoItem.CreateTime = now
	toDoItem.UpdateTime = now

	err = dao.Imp.AddToDoItem(toDoItem)
	if err != nil {
		return nil, err
	}

	return toDoItem, nil
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
