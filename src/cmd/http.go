package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "fit tracker http server",
	Long:  `fit tracker http server`,
	Run: func(cmd *cobra.Command, args []string) {

		h := server.NewFakeFixerIoHandler(fixer)
		router := mux.NewRouter()
		h.Routes(router)

		srv := &http.Server{
			Addr:         ":2000",
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		finishChan := make(chan bool)
		go func(h *http.Server, finish chan bool) {
			e := h.ListenAndServe()
			if e != nil && e != http.ErrServerClosed {
				fmt.Errorf("Could not Listen and server, error %v", e)
				finish <- true
			}
		}(srv, finishChan)
		<-finishChan

	},
}

func init() {
	rootCmd.AddCommand(httpCmd)
}
