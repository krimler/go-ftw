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
		return strings.Contains(headers+response, c.expected.ResponseOrHeadersContains)
	}
	return false
}

// AssertResponseHeaderContains checks that the http response contains the needle
func (c *FTWCheck) AssertNoResponseHeadersContains(response string) bool {
	if c.expected.NoResponseHeadersContains != "" {
		log.Trace().Msgf("ftw/check: is NOT %s contained in response headers %s", c.expected.NoResponseHeadersContains, response)
		return !(strings.Contains(response, c.expected.NoResponseHeadersContains))
	}
	return false
}

// AssertResponseContains checks that the http response contains the needle
func (c *FTWCheck) AssertNoResponseOrHeadersContains(response, headers string) bool {
	if c.expected.NoResponseOrHeadersContains != "" {
		log.Trace().Msgf("ftw/check: is NOT %s contained in response  or response headers %s", c.expected.NoResponseOrHeadersContains, headers+response)
		return !(strings.Contains(headers+response, c.expected.NoResponseOrHeadersContains))
	}
	return false
}
