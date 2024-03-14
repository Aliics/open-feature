package input

import (
	"errors"
	"open-feature/api/model"
)

type Flag struct {
	Key   string `json:"key"`
	Rules []Rule `json:"rules"`
}

func (f Flag) Validate() error {
	if f.Key == "" {
		return ErrFlagKeyRequired
	}

	if len(f.Rules) == 0 {
		return ErrFlagRulesRequired
	}

	for _, rule := range f.Rules {
		switch rule.Type {
		case model.RuleTypeStatic:
			if _, ok := rule.Data.(bool); !ok {
				return ErrStaticRuleExpectsBoolData
			}
		default:
			return ErrUnknownRuleType
		}
	}

	return nil
}

var (
	ErrFlagKeyRequired           = errors.New("key field is required")
	ErrFlagRulesRequired         = errors.New("rules field is required")
	ErrUnknownRuleType           = errors.New("unknown rule type")
	ErrStaticRuleExpectsBoolData = errors.New("static rule must have bool data")
)
