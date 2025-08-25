module github.com/mg6/daily

go 1.24.5

require (
	github.com/adrg/xdg v0.5.3
	github.com/cyp0633/libcaldora v0.4.0
	github.com/emersion/go-ical v0.0.0-20250609112844-439c63cef608
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/teambition/rrule-go v1.8.2 // indirect
	golang.org/x/sys v0.26.0 // indirect
)

replace github.com/emersion/go-ical => github.com/mg6/go-ical v0.0.0-20250817170232-1bb4eca62a2b
