package result

import (
	"fmt"
	"open-feature/database"
)

type Flag struct {
	Key   string `json:"key"`
	Rules []Rule `json:"rules"`
}

func (f *Flag) Map(v any) error {
	switch x := v.(type) {
	case database.Flag:
		f.Key = x.Key
		f.Rules = make([]Rule, len(x.Rules))
		for i, r := range x.Rules {
			rule := &Rule{}
			if err := rule.Map(r); err != nil {
				return err
			}

			f.Rules[i] = *rule
		}

		return nil
	case *database.Flag:
		// Sometimes we deal with pointers.
		return f.Map(*x)
	default:
		return fmt.Errorf("cannot map flag from type: %t", v)
	}
}
