package splunk

import "github.com/roger-russel/splunk-exporter/internal/cmd/dto"

//GetData from splunk
func GetData(flags *dto.Flags) {
	login(flags.User, flags.Password)
}
