package dbhelper

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/senzing-garage/go-messaging/messenger"
)

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
