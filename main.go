package main

import (
	"flag"
	"log"
	"log/slog"
	"os"
	"slices"
	"strings"
	"text/template"

	"github.com/cyp0633/libcaldora/davclient"
	"github.com/emersion/go-ical"
)

type CalObject = davclient.CalendarObject

func main() {
	tplName := flag.String("tpl", "daily", "Template to use")
	flag.Parse()

	logger := GetDefaultLogger()
	slog.SetDefault(logger)

	settings := GetSettings()
	calendars := GetCalendars(settings.CalDAV)

	funcMap := template.FuncMap{
		"LabelOrName": func(todo CalendarConfig) string {
			return todo.GetLabel()
		},
		"TasksByCal": func(name string) []CalObject {
			cal, ok := calendars[name]
			if !ok {
				log.Fatalf("Referenced calendar not found: %v", name)
			}

			opts := davclient.Options{
				Username: settings.CalDAV.Username,
				Password: settings.CalDAV.Password,

				CalendarURL: settings.GetAbsoluteURL(cal.URI),

				Logger: logger,
			}

			client, err := davclient.NewDAVClient(opts)
			if err != nil {
				log.Fatal(err)
			}

			events, err := client.GetAllEvents().ObjectType(ical.CompToDo).Do()
			if err != nil {
				log.Fatal(err)
			}

			return events
		},
		"OnlyCompleted": func(vevs []CalObject) []CalObject {
			return slices.DeleteFunc(vevs, func(v CalObject) bool {
				return v.Event.Props[ical.PropStatus][0].Value != ical.PropCompleted
			})
		},
		"OnlyTodo": func(vevs []CalObject) []CalObject {
			return slices.DeleteFunc(vevs, func(v CalObject) bool {
				return v.Event.Props[ical.PropStatus][0].Value == ical.PropCompleted
			})
		},
		"Top": func(max int, vevs []CalObject) []CalObject {
			if len(vevs) > max {
				return vevs[:max]
			}
			return vevs
		},
		"ByCtimeDesc": func(vevs []CalObject) []CalObject {
			slices.SortStableFunc(vevs, func(a, b CalObject) int {
				return -strings.Compare(GetCtime(a), GetCtime(b))
			})
			return vevs
		},
		"ByMtimeDesc": func(vevs []CalObject) []CalObject {
			slices.SortStableFunc(vevs, func(a, b CalObject) int {
				return -strings.Compare(GetMtime(a), GetMtime(b))
			})
			return vevs
		},
		"Summary": func(v CalObject) string {
			return v.Event.Props[ical.PropSummary][0].Value
		},
	}

	src, ok := settings.Templates[*tplName]
	if !ok {
		log.Fatalf("Template not found: %v", *tplName)
	}

	tpl, err := template.New("note").Funcs(funcMap).Parse(src)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(os.Stdout, settings.ToDos)
	if err != nil {
		log.Fatal(err)
	}
}

// Return last modified time for event.
func GetMtime(e CalObject) string {
	if v, ok := e.Event.Props[ical.PropLastModified]; ok {
		return v[0].Value
	}
	if v, ok := e.Event.Props[ical.PropDateTimeStamp]; ok {
		return v[0].Value
	}
	return ""
}

// Return created time for event.
func GetCtime(e CalObject) string {
	if v, ok := e.Event.Props[ical.PropCreated]; ok {
		return v[0].Value
	}
	if v, ok := e.Event.Props[ical.PropDateTimeStamp]; ok {
		return v[0].Value
	}
	return ""
}
