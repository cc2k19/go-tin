package post

import (
	"github.com/cc2k19/go-tin/storage"
	"net/http"
)

type controller struct {
	repository *storage.Repository
}

func (c *controller) add(wr http.ResponseWriter, r *http.Request) {

}

func (c *controller) get(wr http.ResponseWriter, r *http.Request) {

}
