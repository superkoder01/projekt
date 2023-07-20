package rbac

import (
	"regexp"
	"strings"
)

func (r *Rbac) IsAllowed(role string, endpoint string, httpMethod string) bool {
	return matchRule(r.GetRules(role), endpoint, httpMethod, r.GetApiPrefix())
}

func (r *Rbac) GetRules(tokenRole string) []Rule {
	for _, role := range r.Roles {
		if role.Name == tokenRole {
			return role.Rules
		}
	}
	return nil
}

func (r *Rbac) GetApiPrefix() string {
	return r.ApiPrefix
}

func (r *Rbac) Omit(path string) bool {
	for _, endpoint := range r.OmitEndpoints {
		if strings.HasPrefix(path, r.GetApiPrefix()+endpoint) {
			return true
		}
	}
	return false
}

func matchRule(rules []Rule, endpoint string, httpMethod string, apiPrefix string) bool {
	for _, rule := range rules {
		reg := compileRegexp(apiPrefix + rule.Resource)
		if reg.MatchString(endpoint) {
			return verifyHttpMethod(httpMethod, rule.Verbs)
		}
	}
	return false
}

func verifyHttpMethod(given string, allowed []string) bool {
	for _, method := range allowed {
		if method == "*" || given == method {
			return true
		}
	}
	return false
}

func compileRegexp(input string) *regexp.Regexp {
	pattern := strings.ReplaceAll(input, `/`, `\/`)
	pattern = strings.ReplaceAll(pattern, `{id}`, `\d+`)
	pattern = strings.ReplaceAll(pattern, `*`, `.*`)
	pattern = pattern + "$"
	return regexp.MustCompile(pattern)
}
