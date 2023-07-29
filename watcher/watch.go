package watcher

import (
	"context"
	"fmt"
	"github.com/aboxofsox/npmw/npm"
	"github.com/aboxofsox/wfile"
	"log"
)

func Start(script, root string) {
	wfile.Listen(root, context.TODO(), func(e wfile.Event) {
		if e.Code == wfile.CHANGE {
			fmt.Println("Running:", script)
			err := npm.Run(script, root)
			if err != nil {
				log.Println(err.Error())
			}
		}
		if e.Code == wfile.ERROR {
			fmt.Println("error")
		}
	})
}
