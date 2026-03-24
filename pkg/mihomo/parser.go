package mihomo

import (
	"regexp"

	"github.com/tidwall/gjson"
)

type Connection struct {
	SourceIP string
	DestIP   string
	Host     string
	Protocol string
	Upload   int
	Download int
}

type Log struct {
	Protocol    string
	Source      string
	Destination string
	Chain       string
}

func ParseConnections(jsonStr string) []Connection {
	info := []Connection{}
	resultConnections := gjson.Get(jsonStr, "connections")
	resultConnections.ForEach(func(key, value gjson.Result) bool {
		metadata := gjson.Get(value.Raw, "metadata")
		sourceIP := gjson.Get(metadata.Raw, "sourceIP").Str
		destIP := gjson.Get(metadata.Raw, "destIP").Str
		host := gjson.Get(metadata.Raw, "host").Str
		protocol := gjson.Get(metadata.Raw, "type").Str
		upload := int(gjson.Get(value.Raw, "upload").Num)
		download := int(gjson.Get(value.Raw, "download").Num)
		info = append(info, Connection{
			SourceIP: sourceIP,
			DestIP:   destIP,
			Host:     host,
			Protocol: protocol,
			Upload:   upload,
			Download: download,
		})
		return true
	})
	return info
}

func ParseMemory(jsonStr string) (int, int) {
	inuse := int(gjson.Get(jsonStr, "inuse").Num)
	oslimit := int(gjson.Get(jsonStr, "oslimit").Num)
	return inuse, oslimit
}

func ParseTraffic(jsonStr string) (int, int) {
	up := int(gjson.Get(jsonStr, "up").Num)
	down := int(gjson.Get(jsonStr, "down").Num)
	return up, down
}

func ParseLog(jsonStr string) (string, Log) {
	typeStr := gjson.Get(jsonStr, "type").Str
	payload := gjson.Get(jsonStr, "payload").Str

	// Regex pattern: [TCP] 198.18.0.1:38386 --> www.google.com:443 match DomainKeyword(google) using 🔰 选择节点[🇭🇰 香港Y01]
	re := regexp.MustCompile(`\[([A-Z]+)\]\s+(\S+)\s*-->\s*(\S+).*using\s+(.+)$`)

	matches := re.FindStringSubmatch(payload)
	log := Log{}

	if len(matches) == 5 {
		log.Protocol = matches[1]
		log.Source = matches[2]
		log.Destination = matches[3]
		log.Chain = matches[4]
	}

	return typeStr, log
}
