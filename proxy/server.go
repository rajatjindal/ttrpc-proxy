package proxy

import (
	context "context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	ttrpc "github.com/containerd/ttrpc"
	"github.com/radovskyb/watcher"
)

func Start(dir string) error {
	w := watcher.New()

	w.FilterOps(watcher.Create)
	r := regexp.MustCompile("^[a-z0-9]+.sock$")
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				err := StartOne(filepath.Join(dir, event.Name()))
				if err != nil {
					fmt.Println(err)
				}
			case err := <-w.Error:
				fmt.Println(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch this folder for changes.
	if err := w.Add(dir); err != nil {
		fmt.Println(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 50); err != nil {
		fmt.Println(err)
	}

	return nil
}

func StartOne(socket string) error {
	fmt.Printf("starting proxy for %s\n", socket)
	origSocket := fmt.Sprintf("%s.orig", socket)
	err := moveFile(socket, origSocket)
	if err != nil {
		return err
	}

	s, err := New(origSocket)
	if err != nil {
		return err
	}

	ss, err := ttrpc.NewServer(
		ttrpc.WithUnaryServerInterceptor(serverIntercept),
	)
	if err != nil {
		panic(err)
	}
	defer ss.Close()

	RegisterProxyService(ss, s)

	l, err := net.Listen("unix", socket)
	if err != nil {
		return err
	}

	defer func() {
		l.Close()
		os.Remove(socket)
		moveFile(origSocket, socket)
	}()

	return ss.Serve(context.Background(), l)
}

func serverIntercept(ctx context.Context, um ttrpc.Unmarshaler, i *ttrpc.UnaryServerInfo, m ttrpc.Method) (interface{}, error) {
	log.Println("server interceptor")
	dumpMetadata(ctx)
	return m(ctx, um)
}

func dumpMetadata(ctx context.Context) {
	md, ok := ttrpc.GetMetadata(ctx)
	if !ok {
		return
	}

	if err := json.NewEncoder(os.Stdout).Encode(md); err != nil {
		panic(err)
	}
}

func moveFile(sourcePath, destPath string) error {
	cmd := exec.Command("mv", sourcePath, destPath)
	return cmd.Run()
}

func moveFile2(sourcePath, destPath string) error {
	inputFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't open source file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("Couldn't open dest file: %v", err)
	}
	defer outputFile.Close()

	_, err = io.Copy(outputFile, inputFile)
	if err != nil {
		return fmt.Errorf("Couldn't copy to dest from source: %v", err)
	}
	inputFile.Close()

	err = os.Remove(sourcePath)
	if err != nil {
		return fmt.Errorf("Couldn't remove source file: %v", err)
	}
	return nil
}
