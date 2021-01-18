package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"

	"github.com/gorilla/mux"
)

type DBStub struct {}

func (ds *DBStub) CreateDB(data string) string {
	return fmt.Sprintf("%s %s", "test inserted", data)
}

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
			h := NewHandler(tt.fields.db)

			r := httptest.NewRequest(tt.args.httpMethod, tt.args.path, nil)
			router.HandleFunc("/", h.CreateHandler).Methods(http.MethodGet)
			router.ServeHTTP(tt.args.w, r)
			body, _ := ioutil.ReadAll(tt.args.w.Body)
			if string(tt.want) != string(body) {
				t.Errorf("CreateHandler() Body = %v, want = %v", string(body), string(tt.want))
			}
		})
	}

}
