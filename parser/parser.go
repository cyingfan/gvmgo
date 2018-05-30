package parser

import (
	"regexp"
	"strings"

	"github.com/cyingfan/gvmgo/structs"
)

func ParseCandidateList(listText string) []structs.Sdk {
	sdksText := regexp.MustCompile(`[\-]{10,}?`).Split(listText, -1)
	sdksText = sdksText[1 : len(sdksText)-1]
	re := regexp.MustCompile(`(?P<name>[^\(]+?)\((?P<version>[^\)]+?)\)\s+?(?P<url>https?://[^\n]+?)\n(?s)(?P<description>.+)\$ sdk install (?P<shortname>\w+)`)
	var sdks []structs.Sdk
	for i := 0; i < len(sdksText); i++ {
		fields := re.FindStringSubmatch(strings.TrimSpace(sdksText[i]))
		if len(fields) == 0 {
			continue
		}
		sdks = append(sdks, structs.Sdk{
			Name:        strings.TrimSpace(fields[1]),
			Version:     strings.TrimSpace(fields[2]),
			Url:         strings.TrimSpace(fields[3]),
			Description: strings.TrimSpace(fields[4]),
			ShortName:   strings.TrimSpace(fields[5]),
		})
	}
	return sdks
}
