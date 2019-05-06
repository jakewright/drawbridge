package retry

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Knetic/govaluate"
)

const retryServerErrors = "statusCode == 0 || (statusCode >= 500 && statusCode < 600)"

// defaultRetryMethods are methods that are defined as safe and/or idempotent
var defaultRetryMethods = []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions, http.MethodTrace}

// Config holds options for the retry middleware
type Options struct {
	Attempts        int      `mapstructure:"attempts"`
	RetryMethods    []string `mapstructure:"retry_methods"`
	Predicate       string   `mapstructure:"predicate"`
	BackoffStrategy string   `mapstructure:"backoff_strategy"`

	expression      *govaluate.EvaluableExpression `mapstructure:"-"`
	backoffStrategy BackoffStrategy                `mapstructure:"-"`
}

// Validate returns an error if there are any problems with the config
func (o *Options) Validate() error {
	if o == nil {
		return errors.New("config is nil")
	}

	if o.Attempts < 1 {
		return errors.New("attempts must be > 0")
	}

	if len(o.RetryMethods) == 0 {
		o.RetryMethods = defaultRetryMethods
	}

	if o.Predicate == "" {
		o.Predicate = retryServerErrors
	}

	expression, err := govaluate.NewEvaluableExpression(o.Predicate)
	if err != nil {
		return err
	}
	o.expression = expression

	switch o.BackoffStrategy {
	case "exponential", "":
		o.backoffStrategy = &ExponentialBackoff{}
	default:
		return fmt.Errorf("unknown backoff strategy %s", o.BackoffStrategy)
	}

	return nil
}
