# Daily

Generate daily notes from CalDAV tasks and format using templates.

## Features

- Connects to CalDAV server to retrieve TODO items
- Supports multiple calendar sources
- Customizable templates for note formatting
- Filter tasks by completion status
- Sort by creation or modification time
- Limit number of results displayed

## Installation

```sh
go install github.com/mg6/daily@latest
```

## Configuration

Copy the example settings file:

```sh
cp settings.example.yaml ~/.config/daily/settings.yaml
```

Edit `settings.yaml` as required:

- Update `caldav.url` with your CalDAV server link
- Set `caldav.username` and `caldav.password`
- Configure TODO calendars to read from in the `todos` section
- Customize templates as needed

## Usage

```sh
# Generate daily note using default template
daily

# Store with other daily notes
daily | tee ~/Journal/$(date +%Y)/$(date +%F).md

# Use a specific template
daily -tpl weekly
```

## Template Functions

The following functions are available in templates:

- `LabelOrName`: Get calendar label or name
- `TasksByCal`: Retrieve tasks from a specific calendar
- `OnlyCompleted`: Filter for completed tasks only
- `OnlyTodo`: Filter for pending tasks only
- `Top N`: Limit results to first N items
- `ByCtimeDesc`: Sort by creation time (newest first)
- `ByMtimeDesc`: Sort by modification time (newest first)
- `Summary`: Get task summary/title
- `Completed`: Get task completion time

## Example Output

```
* * *

## Tasks
What's the next step right now?

- [ ] Review pull request 42
- [ ] Update documentation
- [ ] Fix bug in authentication
- [x] Complete project setup
- [x] Deploy to staging

## Goals
What's important this month?

- [ ] Launch new feature
- [ ] Improve test coverage
- [x] Refactor legacy code
```
