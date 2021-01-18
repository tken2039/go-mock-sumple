package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"

	"github.com/gorilla/mux"
)

// DBStub は CreateDBi(string) stringをメソッドに持つクラス
// すなわち、暗黙的に interface "DBModel" を実装している
type DBStub struct {}

// CreateDB は DBStub のメソッド
func (ds *DBStub) CreateDB(data string) string {
	return fmt.Sprintf("%s %s", "test inserted", data)
}

// NewDBStub は *DBStub を返却する
func NewDBStub() *DBStub {
	return &DBStub{}
}

func TestHandler_CreateHandler(t *testing.T) {
	router := mux.NewRouter()

	type fields struct {
		db *DBStub
	}
	type args struct {
		w		*httptest.ResponseRecorder
		httpMethod	string
		path		string
	}
	tests := []struct {
		name	string
		fields	fields
		args	args
		want	[]byte
	}{
		{
			name:	"nomal",
			fields:	fields {
				db: NewDBStub(),
			},
			args:	args {
				w:		httptest.NewRecorder(),
				httpMethod:	http.MethodGet,
				path:		"/",
			},
			want:	[]byte("test inserted Data"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// NewHandler 関数の引数には、 CreateDB(string) string を実装
			// している DBStub 型の変数を与えることができる
			h := NewHandler(tt.fields.db)

			// 仮リクエストを生成する便利関数
			r := httptest.NewRequest(tt.args.httpMethod, tt.args.path, nil)

			// 受け付ける router を定義
			router.HandleFunc("/", h.CreateHandler).Methods(http.MethodGet)

			// 仮リクエストを実行 
			router.ServeHTTP(tt.args.w, r)
			body, _ := ioutil.ReadAll(tt.args.w.Body)
			if string(tt.want) != string(body) {
				t.Errorf("CreateHandler() Body = %v, want = %v", string(body), string(tt.want))
			}
		})
	}

}
