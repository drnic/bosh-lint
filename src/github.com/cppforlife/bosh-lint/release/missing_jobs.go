package release

import (
	boshrel "github.com/cloudfoundry/bosh-cli/release"

	check "github.com/cppforlife/bosh-lint/check"
)

type MissingJobs struct {
	context check.Context
	release boshrel.Release
}

func NewMissingJobs(context check.Context, release boshrel.Release) MissingJobs {
	return MissingJobs{context, release}
}

func (c MissingJobs) Description() check.Description {
	return check.Description{
		Context_: c.context,
		Purpose_: "if at least one job is present",
	}
}

func (c MissingJobs) Check() ([]check.Suggestion, error) {
	var sugs []check.Suggestion

	if len(c.release.Jobs()) == 0 && len(c.release.Packages()) > 0 {
		sugs = append(sugs, check.Simple{
			Context_:    c.context,
			Problem_:    "Zero jobs",
			Resolution_: "Add at least one job which references a package",
		})
	}

	return sugs, nil
}
