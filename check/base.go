package check

import (
	"time"

	"github.com/fzipi/go-ftw/config"
	"github.com/fzipi/go-ftw/test"
	"github.com/fzipi/go-ftw/waflog"
	"github.com/rs/zerolog/log"
)

// FTWCheck is the base struct for checking test results
type FTWCheck struct {
	log       *waflog.FTWLogLines
	expected  *test.Output
	overrides *config.FTWTestOverride
}

// NewCheck creates a new FTWCheck, allowing to inject the configuration
func NewCheck(c *config.FTWConfiguration) *FTWCheck {
	check := &FTWCheck{
		log: &waflog.FTWLogLines{
			FileName:     c.LogFile,
			TimeRegex:    c.LogType.TimeRegex,
			TimeFormat:   c.LogType.TimeFormat,
			Since:        time.Now(),
			Until:        time.Now(),
			TimeTruncate: c.LogType.TimeTruncate,
			LogTruncate:  c.LogTruncate,
		},
		expected:  &test.Output{},
		overrides: &c.TestOverride,
	}

	log.Trace().Msgf("check/base: time will be truncated to %s", check.log.TimeTruncate)
	log.Trace().Msgf("check/base: logfile will be truncated? %t", check.log.LogTruncate)

	return check
}

// SetRoundTripTime sets the time the roundtrip took so we can check logs with it
func (c *FTWCheck) SetRoundTripTime(since time.Time, until time.Time) {
	c.log.Since = since
	c.log.Until = until
}

// SetExpectTestOutput sets the combined expected output from this test
func (c *FTWCheck) SetExpectTestOutput(t *test.Output) {
	c.expected = t
}

// SetExpectStatus sets to expect the HTTP status from the test to be in the integer range passed
func (c *FTWCheck) SetExpectStatus(s []int) {
	c.expected.Status = s
}

// SetExpectResponse sets the response we expect in the text from the server
func (c *FTWCheck) SetExpectResponse(response string) {
	c.expected.ResponseContains = response
}

// SetExpectResponseHeaders sets the response we expect in the text from the server
func (c *FTWCheck) SetExpectResponseHeaders(response string) {
	c.expected.ResponseHeadersContains = response
}

// SetExpectResponseOrHeaders sets the response we expect in the text from the server
func (c *FTWCheck) SetExpectResponseOrHeaders(response string) {
	c.expected.ResponseOrHeadersContains = response
}

// SetExpectError sets the boolean if we are expecting an error from the server
func (c *FTWCheck) SetExpectError(expect bool) {
	c.expected.ExpectError = expect
}

// SetLogContains sets the string to look for in logs
func (c *FTWCheck) SetLogContains(contains string) {
	c.expected.LogContains = contains
}

// SetNoLogContains sets the string to look that should not present in logs
func (c *FTWCheck) SetNoLogContains(contains string) {
	c.expected.NoLogContains = contains
}

// ForcedIgnore check if this id need to be ignored from results
func (c *FTWCheck) ForcedIgnore(id string) bool {
	_, ok := c.overrides.Ignore[id]
	return ok
}

// ForcedPass check if this id need to be ignored from results
func (c *FTWCheck) ForcedPass(id string) bool {
	_, ok := c.overrides.ForcePass[id]
	return ok
}

// ForcedFail check if this id need to be ignored from results
func (c *FTWCheck) ForcedFail(id string) bool {
	_, ok := c.overrides.ForceFail[id]
	return ok
}
