package names

import (
	"errors"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

const (
	ErrActionReading           = "reading"
	ErrActionDeleting          = "deleting"
	ErrActionUpdating          = "updating"
	ErrActionCreating          = "creating"
	ErrActionSetting           = "setting"
	ErrActionCheckingExistence = "checking existence"
	ErrActionCheckingDestroyed = "checking destroyed"
)

// ProblemStandardMessage is a standardized message for reporting errors and warnings
func ProblemStandardMessage(service, action, resource, id string, gotError error) string {
	hf, err := FullHumanFriendly(service)

	if err != nil {
		return fmt.Sprintf("finding human-friendly name for service (%s) while creating error (%s, %s, %s, %s): %s", service, action, resource, id, gotError, err)
	}

	if gotError == nil {
		return fmt.Sprintf("%s %s %s (%s)", action, hf, resource, id)
	}

	return fmt.Sprintf("%s %s %s (%s): %s", action, hf, resource, id, gotError)
}

// Error returns an errors.Error with a standardized error message
func Error(service, action, resource, id string, gotError error) error {
	return errors.New(ProblemStandardMessage(service, action, resource, id, gotError))
}

// DiagError returns a 1-length diag.Diagnostics with a diag.Error-level diag.Diagnostic
// with a standardized error message
func DiagError(service, action, resource, id string, gotError error) diag.Diagnostics {
	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Error,
			Summary:  ProblemStandardMessage(service, action, resource, id, gotError),
		},
	}
}

// AddWarning returns diag.Diagnostics with an additional diag.Diagnostic containing
// a warning using a standardized problem message
func AddWarning(diags diag.Diagnostics, service, action, resource, id string, gotError error) diag.Diagnostics {
	return append(diags,
		diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  ProblemStandardMessage(service, action, resource, id, gotError),
		},
	)
}

// WarnLog logs to the default logger a standardized problem message
func WarnLog(service, action, resource, id string, gotError error) {
	log.Printf("[WARN] %s", ProblemStandardMessage(service, action, resource, id, gotError))
}

func LogNotFoundRemoveState(service, action, resource, id string) {
	WarnLog(service, action, resource, id, errors.New("not found, removing from state"))
}
