package encrypt

import (
	uuid "github.com/satori/go.uuid"
)

func UUID() string {
	return uuid.NewV4().String()
}

func UUIDV1() string {
	return uuid.NewV1().String()
}
