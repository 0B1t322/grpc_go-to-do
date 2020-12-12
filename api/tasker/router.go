package tasker

import (
	"context"
	"encoding/json"
	"errors"
	"go-to-do/api"
	"net/http"
	"time"

	pb "github.com/0B1t322/tasker/tasker"
	taskerclient "github.com/0B1t322/tasker/tasker_client"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var conf taskerclient.Config

var (
	ErrNotFoundConn = api.NewError(errors.New("Not found conn"))
)

func init() {
	conf = taskerclient.Config{
		Port: ":8080",
	}
}

//CreateTask ....
func CreateTask(w http.ResponseWriter, r *http.Request) {
	log.Info("Create Task")
	conn, ok := r.Context().Value("conn").(*grpc.ClientConn)
	if !ok || conn == nil {
		log.Error("Not found conn")

		if data, err := ErrNotFoundConn.Marshall(); err != nil {
			log.Error(err)
		} else {
			w.Write(data)
		}

		return
	}

	client := pb.NewTaskerClient(conn)
	req := &pb.TaskRequest{}

	json.NewDecoder(r.Body).Decode(req)
	req.CreatesTime = time.Now().Format(time.Stamp)

	resp, err := client.CreateTask(context.Background(), req)
	if err != nil {
		log.Error(err)
	}

	


	if data, err := json.Marshal(resp); err != nil {
		log.Error(err)
	} else {
		w.Write(data)
	}
}

// MarkTask ....
func MarkTask(w http.ResponseWriter, r *http.Request) {
	conn, ok := r.Context().Value("conn").(*grpc.ClientConn)
	if !ok || conn == nil {
		log.Error("Not found conn")

		if data, err := ErrNotFoundConn.Marshall(); err != nil {
			log.Error(err)
		} else {
			w.Write(data)
		}

		return
	}

	client := pb.NewTaskerClient(conn)
	req := &pb.MarkRequest{}

	json.NewDecoder(r.Body).Decode(req)

	resp, err := client.MarkTask(context.Background(), req)
	if err != nil {
		log.Error(err)
	}
	log.Info("Error: ",resp.Error)
	log.Info("err: ", err)
	

	if data, err := json.Marshal(resp); err != nil {
		log.Error(err)
	} else {
		w.Write(data)
	}
}

// ArchiveTask ....
func ArchiveTask(w http.ResponseWriter, r *http.Request) {
	conn, ok := r.Context().Value("conn").(*grpc.ClientConn)
	if !ok || conn == nil {
		log.Error("Not found conn")

		if data, err := ErrNotFoundConn.Marshall(); err != nil {
			log.Error(err)
		} else {
			w.Write(data)
		}

		return
	}

	client := pb.NewTaskerClient(conn)
	req := &pb.ArchiveRequest{}

	json.NewDecoder(r.Body).Decode(req)

	resp, err := client.ArchiveTask(context.Background(), req)
	if err != nil {
		log.Error(err)
	}
	log.Info("Error: ",resp.Error)
	log.Info("err: ", err)

	if data, err := json.Marshal(resp); err != nil {
		log.Error(err)
	} else {
		w.Write(data)
	}
}

// GetTask ...
func GetTask(w http.ResponseWriter, r *http.Request) {
	conn, ok := r.Context().Value("conn").(*grpc.ClientConn)
	if !ok || conn == nil {
		log.Error("Not found conn")

		if data, err := ErrNotFoundConn.Marshall(); err != nil {
			log.Error(err)
		} else {
			w.Write(data)
		}

		return
	}

	client := pb.NewTaskerClient(conn)
	req := &pb.GetTaskRequest{}

	json.NewDecoder(r.Body).Decode(req)

	resp, err := client.GetTask(context.Background(), req)
	if err != nil {
		log.Error(err)
	}
	log.Info("Error: ",resp.Error)
	log.Info("err: ", err)

	if data, err := json.Marshal(resp); err != nil {
		log.Error(err)
	} else {
		w.Write(data)
	}
}

// GetAllTasks ....
func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	conn, ok := r.Context().Value("conn").(*grpc.ClientConn)
	if !ok || conn == nil {
		log.Error("Not found conn")

		if data, err := ErrNotFoundConn.Marshall(); err != nil {
			log.Error(err)
		} else {
			w.Write(data)
		}

		return
	}

	client := pb.NewTaskerClient(conn)
	req := &pb.GetAllTaskRequest{}

	json.NewDecoder(r.Body).Decode(req)

	resp, err := client.GetAllTasks(context.Background(), req)
	if err != nil {
		log.Error(err)
	}

	log.Info("Error: ",resp.Error)
	log.Info("err: ", err)

	if data, err := json.Marshal(resp); err != nil {
		log.Error(err)
	} else {
		w.Write(data)
	}
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/tasker/CreateTask", CreateTask).Methods("POST")
	r.HandleFunc("/api/tasker/MarkTask", MarkTask).Methods("PATCH")
	r.HandleFunc("/api/tasker/ArchiveTask", ArchiveTask).Methods("PUT")
	r.HandleFunc("/api/tasker/GetTask", GetTask).Methods("GET")
	r.HandleFunc("/api/tasker/GetAllTasks", GetAllTasks).Methods("GET")

	r.Use(taskerclient.MiddlewareTasker(conf))

	return r
}