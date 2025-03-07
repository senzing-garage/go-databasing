package dbhelper

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/senzing-garage/go-messaging/messenger"
)

/*
Function ExtractSqliteDatabaseFilename parses an SQLite Database URL
and returns the fully qualified filename.

Input
  - databaseURL: An SQLite style database URL.

Output
  - A fully qualified path to the SQLite database file.
*/
func ExtractSqliteDatabaseFilename(databaseURL string) (string, error) {
	var result = ""

	if !strings.HasPrefix(databaseURL, "sqlite3:") {
		return result, fmt.Errorf("sqlite3 URL schema needed")
	}

	parsedURL, err := url.Parse(databaseURL)
	if err != nil {
		return result, err
	}

	if parsedURL.Scheme != "sqlite3" {
		return result, fmt.Errorf("sqlite3 URL schema needed")
	}

	return extractSqliteDatabaseFilenameForOsArch(parsedURL)
}

/*
Function GetMessenger is a factory to produce a [messenger.Messenger].

Input
  - componentID: A 4-digit number (0...9999) identifying the component creating the message.
  - idMessages: A map of error numbers to messaage templates.
  - callerSkip: Number of stack frames to ascend in [runtime.Caller].
  - options: 0 or more [messenger.OptionsXxx].

Output
  - A configured [messenger.Messenger] implementation.

[messenger.OptionsXxx]: https://pkg.go.dev/github.com/senzing-garage/go-messaging/messenger#OptionCallerSkip
[runtime.Caller]: https://pkg.go.dev/runtime#Caller
[messenger.Messenger]: https://pkg.go.dev/github.com/senzing-garage/go-messaging/messenger#Messenger
*/
func GetMessenger(componentID int, idMessages map[int]string, callerSkip int, options ...interface{}) messenger.Messenger {
	optionMessageIDTemplate := fmt.Sprintf("%s%04d", MessageIDPrefix, componentID) + "%04d"
	messengerOptions := []interface{}{
		messenger.OptionCallerSkip{Value: callerSkip},
		messenger.OptionIDMessages{Value: idMessages},
		messenger.OptionMessageFields{Value: []string{"id", "reason"}},
		messenger.OptionMessageIDTemplate{Value: optionMessageIDTemplate},
	}
	messengerOptions = append(messengerOptions, options...)
	result, err := messenger.New(messengerOptions...)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Function ParseDatabaseURL is fixes database URLs prior to parsing them.

Input
  - databaseURL: The database URL to be parsed.

Output
  - [url.URL]

[url.URL]: https://pkg.go.dev/net/url#URL
*/
// func ParseDatabaseURL(databaseURL string) (*url.URL, error) {
// 	result, err := url.Parse(databaseURL)
// 	if err != nil {
// 		if strings.HasPrefix(databaseURL, "postgresql") {
// 			index := strings.LastIndex(databaseURL, ":")
// 			newDatabaseURL := databaseURL[:index] + "/" + databaseURL[index+1:]
// 			result, err = url.Parse(newDatabaseURL)
// 		}
// 	}
// 	return result, err
// }
