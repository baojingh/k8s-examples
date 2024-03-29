package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var service string

func init() {
	flag.StringVar(&service, "service", "", "filter which service to see")
}

func main() {
	flag.Parse()
	var b strings.Builder
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		// convert json to map for process
		m := make(map[string]interface{})
		err := json.Unmarshal([]byte(s), &m)
		if err != nil {
			if service == "" {
				fmt.Println(s)
			}
			continue
		}
		if service != "" && m["service"] != service {
			continue
		}

		/**
		{"level":"info","ts":"2024-03-13T11:16:40.876+0800",
		"caller":"sales-api/main.go:38","msg":"startup",
		"service":"SALES-API","GOMAXPROCS":8}
		**/
		// traceID := "00000000-0000-0000-0000-000000000000"
		b.Reset()
		b.WriteString(fmt.Sprintf("%s %s %s %s %s ",
			m["service"],
			m["ts"],
			m["level"],
			m["caller"],
			m["msg"],
		))

		for k, v := range m {
			switch k {
			case "service", "ts", "level", "caller", "msg":
				continue
			}
			b.WriteString(fmt.Sprintf("%s[%v]: ", k, v))
		}

		out := b.String()
		/**
		SALES-API: 2024-03-13T15:07:21.330+0800:info:
		00000000-0000-0000-0000-000000000000: sales-api/main.go:38: startup
		**/
		fmt.Println(out[:len(out)-2])

		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}

}
