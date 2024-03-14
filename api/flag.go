package api

import (
	"errors"
	"net/http"
	"open-feature/api/input"
	"open-feature/api/model"
	"open-feature/api/result"
	"open-feature/database"
)

func (s *Server) listFlags(w http.ResponseWriter, _ *http.Request) {
	flags, err := s.Database.All()
	if err != nil {
		// This has no user input effectively.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result.Write[*result.Flag](w, flags)
}

func (s *Server) getFlag(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue(PathValueKey)
	flag, err := s.Database.Get(key)
	if err != nil {
		if errors.Is(err, database.ErrFlagNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result.Write[*result.Flag](w, flag)
}

func (s *Server) putFlag(w http.ResponseWriter, r *http.Request) {
	inputFlag, err := input.GetValidatedInput[input.Flag](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.Put(databaseFlagFromInputFlag(inputFlag))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) deleteFlag(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue(PathValueKey)

	err := s.Database.Delete(key)
	if err != nil {
		if errors.Is(err, database.ErrFlagNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func databaseFlagFromInputFlag(inputFlag *input.Flag) database.Flag {
	rules := make([]database.Rule, len(inputFlag.Rules))
	for i, rule := range inputFlag.Rules {
		switch rule.Type {
		case model.RuleTypeStatic:
			rules[i] = database.StaticRule(rule.Data.(bool))
		}
	}
	return database.Flag{
		Key:   inputFlag.Key,
		Rules: rules,
	}
}

const (
	PathValueKey = "key"
)
