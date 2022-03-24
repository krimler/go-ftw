package check

import (
	"github.com/rs/zerolog/log"
	"strings"
)

// AssertResponseContains checks that the http response contains the needle
func (c *FTWCheck) AssertResponseContains(response string) bool {
	if c.expected.ResponseContains != "" {
		log.Trace().Msgf("ftw/check: is %s contained in response %s", c.expected.ResponseContains, response)
		return strings.Contains(response, c.expected.ResponseContains)
	}
	return false
}

// AssertResponseHeaderContains checks that the http response contains the needle
func (c *FTWCheck) AssertResponseHeadersContains(response string) bool {
	if c.expected.ResponseHeadersContains != "" {
		log.Trace().Msgf("ftw/check: is %s contained in response headers %s", c.expected.ResponseHeadersContains, response)
		return strings.Contains(response, c.expected.ResponseHeadersContains)
	}
	return false
}

// AssertResponseContains checks that the http response contains the needle
func (c *FTWCheck) AssertResponseOrHeadersContains(response, headers string) bool {
	if c.expected.ResponseOrHeadersContains != "" {
		log.Trace().Msgf("ftw/check: is %s contained in response  or response headers %s", c.expected.ResponseOrHeadersContains, headers+response)
		return strings.Contains(response, c.expected.ResponseOrHeadersContains)
	}
	return false
}
