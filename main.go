package token

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	// NOTE: Don't change this, the auth settings on the providers
	// are coded to this path for this example.
	Address string = ":8080"
)

func main() {
        m := martini.Classic()
        token := token.NewToken();
        // Hardcore for test
		m.Use(
			Auth.BasicFunc(func(username, password string) bool {
			    return username == "test" && password == "test"
			  }
			)
        webservice.RegisterWebService(token, m)
        webservice.RegisterWebService(user, m)

       	m.Get("/",func() string {
       		return "Hello, Andrew"
       		})



        m.Run()
}


