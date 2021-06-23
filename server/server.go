package server


package server

import (
	//"github.com/eapache/queue"
	"github.com/netyut1703/shopingProject/queue"
	"github.com/netyut1703/shopingProject//store"
	"github.com/netyut1703/shopingProject//aws/s3"
)

type Server struct {
	Store *store.Store
	Queue *queue.Queue
	S3Store *s3.S3Store
}

func New(store *store.Store, queue *queue.Queue, s3Store *s3.S3Store) *Server {
	srv := &Server{}
	srv.Store = store
	srv.Queue = queue
	srv.S3Store = s3Store
	return srv
}
