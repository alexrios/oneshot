package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (s *Server) handleDownload(w http.ResponseWriter, r *http.Request) {
	if s.authenticating {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", "Basic")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Whichever field is missing is not checked
		if s.Username != "" && s.Username != username {
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Set("WWW-Authenticate", "Basic")
			return
		}
		if s.Password != "" && s.Password != password {
			w.Header().Set("WWW-Authenticate", "Basic")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	s.mutex.Lock()
	if s.done {
		s.mutex.Unlock()
		w.WriteHeader(http.StatusGone)
		return
	}
	s.done = true
	s.mutex.Unlock()

	if s.InfoLog != nil {
		s.InfoLog.Printf("client connected: %s\n", r.RemoteAddr)
	}

	err := s.file.Open()
	defer s.file.Close()
	if err != nil {
		if s.ErrorLog != nil {
			s.ErrorLog.Println(err.Error())
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	before := time.Now()
	_, err = io.Copy(w, s.file)
	duration := time.Since(before)
	if s.ErrorLog != nil && err != nil {
		go s.Stop(context.Background())
		s.ErrorLog.Println(err.Error())
		return
	}

	if s.Download {
		w.Header().Set("Content-Disposition",
			fmt.Sprintf("attachment;filename=%s", s.file.Name),
		)
	}

	w.Header().Set("Content-Type", s.file.MimeType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", s.file.Size()))

	if s.InfoLog != nil && err == nil {
		s.InfoLog.Printf(
			"file was downloaded:\n\tname: %s\n\ttime: %s\n\trate: %d Bytes / %s\n\tclient address: %s\n",
			s.file.Name,
			before.String(),
			s.file.Size(),
			duration.String(),
			r.RemoteAddr,
		)
	}

	// Stop() method needs to run on seperate goroutine.
	// Otherwise, we deadlock since http.Server.Shutdown()
	// wont return until this function returns.
	go s.Stop(context.Background())
}
