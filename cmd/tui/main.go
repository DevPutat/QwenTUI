package main

import (
	// "flag"
	// "fmt"

	"github.com/DevPutat/QwenTUI/internal/config"
	"github.com/DevPutat/QwenTUI/internal/tui"
	// "github.com/DevPutat/QwenTUI/internal/request"
)

func main() {
	// flagQuery := flag.String("query", "", "query for request")
	//flagKey := flag.String("key", "", "key for api")
	// flag.Parse()

	c := config.Config()
	tui.NewApp(c)

	// if len(*flagQuery) > 0 {
	// 	fmt.Println(request.Send(*flagQuery, c))
	// }
}
