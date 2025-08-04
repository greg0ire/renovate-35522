package stuff

import (
	"context"

	testcontainers_wiremock "github.com/wiremock/wiremock-testcontainers-go"
)

func main() {
	testcontainers_wiremock.RunContainerAndStopOnCleanup(context.Background(), nil)
}
