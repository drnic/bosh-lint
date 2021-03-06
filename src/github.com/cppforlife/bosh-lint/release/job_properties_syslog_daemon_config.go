package release

import (
	"strings"

	boshjob "github.com/cloudfoundry/bosh-cli/release/job"

	check "github.com/cppforlife/bosh-lint/check"
)

type JobPropertiesSyslogDaemonConfig struct {
	context check.Context
	job     *boshjob.Job
}

func NewJobPropertiesSyslogDaemonConfig(context check.Context, job *boshjob.Job) JobPropertiesSyslogDaemonConfig {
	return JobPropertiesSyslogDaemonConfig{context, job}
}

func (c JobPropertiesSyslogDaemonConfig) Description() check.Description {
	return check.Description{
		Context_:   c.context,
		Purpose_:   "if job unnecessarily asks for syslog configuration",
		Reasoning_: "It's recommended to use cloudfoundry/syslog-release for syslog configuration",
	}
}

func (c JobPropertiesSyslogDaemonConfig) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	for propName, _ := range c.job.Properties {
		if strings.Contains(propName, "syslog_daemon_config") {
			sugs = append(sugs, check.Simple{
				Context_:    c.context,
				Problem_:    "Asks for syslog configuration",
				Resolution_: "Remove 'syslog_daemon_config' set of properties",
			})

			break
		}
	}

	return sugs, nil
}
