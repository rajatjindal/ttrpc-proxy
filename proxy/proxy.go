package proxy

import (
	context "context"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/containerd/containerd/api/runtime/task/v2"
	ttrpc "github.com/containerd/ttrpc"
)

type Server struct {
	tc task.TaskService
}

func New(proxyTo string) (*Server, error) {
	conn, err := net.Dial("unix", proxyTo)
	if err != nil {
		return nil, err
	}

	client := ttrpc.NewClient(conn)

	return &Server{
		tc: task.NewTaskClient(client),
	}, nil
}

func RegisterProxyService(srv *ttrpc.Server, svc *Server) {
	srv.RegisterService("containerd.task.v2.Task", &ttrpc.ServiceDesc{
		Methods: map[string]ttrpc.Method{
			"State": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.StateRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("State")
				defer func() {
					fmt.Println("State done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}

				return svc.tc.State(ctx, &req)
			},
			"Create": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.CreateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Create")
				defer func() {
					fmt.Println("Create done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Create(ctx, &req)
			},
			"Start": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.StartRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Start")
				defer func() {
					fmt.Println("Start done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Start(ctx, &req)
			},
			"Delete": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.DeleteRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Delete")
				defer func() {
					fmt.Println("Delete done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Delete(ctx, &req)
			},
			"Pids": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.PidsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Pids")
				defer func() {
					fmt.Println("Pids done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Pids(ctx, &req)
			},
			"Pause": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.PauseRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Pause")
				defer func() {
					fmt.Println("Pause done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Pause(ctx, &req)
			},
			"Resume": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.ResumeRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Resume")
				defer func() {
					fmt.Println("Resume done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Resume(ctx, &req)
			},
			"Checkpoint": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.CheckpointTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Checkpoint")
				defer func() {
					fmt.Println("Checkpoint done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Checkpoint(ctx, &req)
			},
			"Kill": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.KillRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Kill")
				defer func() {
					fmt.Println("Kill done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Kill(ctx, &req)
			},
			"Exec": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.ExecProcessRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Exec")
				defer func() {
					fmt.Println("Exec done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Exec(ctx, &req)
			},
			"ResizePty": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.ResizePtyRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("ResizePty")
				defer func() {
					fmt.Println("ResizePty done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.ResizePty(ctx, &req)
			},
			"CloseIO": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.CloseIORequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("CloseIO")
				defer func() {
					fmt.Println("CloseIO done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.CloseIO(ctx, &req)
			},
			"Update": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.UpdateTaskRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Update")
				defer func() {
					fmt.Println("Update done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Update(ctx, &req)
			},
			"Wait": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.WaitRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Wait")
				defer func() {
					fmt.Println("Wait done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Wait(ctx, &req)
			},
			"Stats": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.StatsRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Stats")
				defer func() {
					fmt.Println("Stats done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Stats(ctx, &req)
			},
			"Connect": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.ConnectRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Connect")
				defer func() {
					fmt.Println("Connect done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Connect(ctx, &req)
			},
			"Shutdown": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
				var req task.ShutdownRequest
				if err := unmarshal(&req); err != nil {
					return nil, err
				}

				fmt.Println("Shutdown")
				defer func() {
					fmt.Println("Shutdown done")
				}()
				if err := json.NewEncoder(os.Stdout).Encode(&req); err != nil {
					panic(err)
				}
				return svc.tc.Shutdown(ctx, &req)
			},
		},
	})
}
