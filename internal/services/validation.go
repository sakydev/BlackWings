package services

import (
	"BlackWings/internal/types"
	"fmt"
)

type FlagRule struct {
	Required      bool
	MinimumLength int
	MaximumLength int
	OneOf         []string
}

func ValidateSearchFlags(flags types.SearchFlags) error {
	searchRules := getSearchRules()

	for flagName, rule := range searchRules {
		value := getValue(flags, flagName)
		err := validateFlag(flagName, value, rule)
		if err != nil {
			return err
		}
	}
	return nil
}

func getValue(flags types.SearchFlags, flagName string) interface{} {
	switch flagName {
	case "Query":
		return flags.Query
	case "Apps":
		return flags.Apps
	case "Include":
		return flags.Include
	case "Exclude":
		return flags.Exclude
	case "Before":
		return flags.Before
	case "After":
		return flags.After
	case "Sort":
		return flags.Sort
	case "Order":
		return flags.Order
	case "Limit":
		return flags.Limit
	case "Offset":
		return flags.Offset
	default:
		return nil
	}
}
func isEmpty(isString, isInteger, isArrayOfStrings, isArrayOfIntegers bool, value interface{}) bool {
	if value == nil {
		return true
	}
	if isString && value == "" {
		return true
	}
	if isInteger && value == 0 {
		return true
	}
	if isArrayOfStrings && len(value.([]string)) == 0 {
		return true
	}
	if isArrayOfIntegers && len(value.([]int)) == 0 {
		return true
	}
	return false
}
func validateFlag(flagName string, value interface{}, rule FlagRule) error {
	stringValue, isString := value.(string)
	arrayOfStringsValue, isArrayOfStrings := value.([]string)
	_, isInteger := value.(string)
	_, isArrayOfIntegers := value.([]int)
	emptyField := isEmpty(isString, isInteger, isArrayOfStrings, isArrayOfIntegers, value)
	if rule.Required && emptyField {
		return fmt.Errorf("%s is required", flagName)
	}
	if emptyField {
		return nil
	}
	// only applies to strings
	if isArrayOfStrings {
		if rule.MinimumLength > 0 {
			if len(stringValue) < rule.MinimumLength {
				fmt.Printf("Field Name: %s, Field Value: %s\n", flagName, stringValue)
				return fmt.Errorf("%s must be at least %d characters long", flagName, rule.MinimumLength)
			}
		}
		if rule.MaximumLength > 0 {
			if len(stringValue) > rule.MaximumLength {
				return fmt.Errorf("%s must be at most %d characters long", flagName, rule.MaximumLength)
			}
		}
	}
	// only applies to arrays of strings or integers
	if isArrayOfStrings || isArrayOfIntegers {
		if len(rule.OneOf) > 0 {
			inputValues := []string{}
			if isString {
				inputValues = append(inputValues, stringValue)
			} else if isArrayOfStrings {
				inputValues = append(inputValues, arrayOfStringsValue...)
			}
			var found bool
			for _, inputValue := range inputValues {
				found = false
				for _, allowedValue := range rule.OneOf {
					if inputValue == allowedValue {
						found = true
						break
					}
				}
			}
			if !found {
				return fmt.Errorf("%s must be one of: %s", flagName, formatAllowedValues(rule.OneOf))
			}
		}
	}
	return nil
}
func formatAllowedValues(allowedValues []string) string {
	return "\"" + fmt.Sprintf("%s", allowedValues) + "\""
}
func getSearchRules() map[string]FlagRule {
	return map[string]FlagRule{
		"Query": {
			Required:      true,
			MinimumLength: 3,
			MaximumLength: 100,
		},
		"Apps": {
			Required: false,
			OneOf:    []string{"gmail", "slack", "drive"},
		},
		"Include": {
			Required:      false,
			MinimumLength: 3,
			MaximumLength: 100,
		},
		"Exclude": {
			Required:      false,
			MinimumLength: 3,
			MaximumLength: 100,
		},
		"FileTypes": {
			Required: false,
			OneOf:    []string{"pdf", "doc", "xls"},
		},
		"Before": {
			Required:      false,
			MinimumLength: 10,
			MaximumLength: 10,
		},
		"After": {
			Required:      false,
			MinimumLength: 10,
			MaximumLength: 10,
		},
		"Sort": {
			Required: false,
			OneOf:    []string{"relevance", "date"},
		},
		"Order": {
			Required: false,
			OneOf:    []string{"asc", "desc"},
		},
		"Limit": {
			Required: false,
		},
		"Offset": {
			Required: false,
		},
	}
}
