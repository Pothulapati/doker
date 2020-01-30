package docker

import (
	"testing"

	"github.com/deislabs/smi-sdk-go/pkg/apis/metrics"
	"github.com/stretchr/testify/suite"
)

func (s *EdgeTestSuite) TestGet() {
	for kind := range metrics.AvailableKinds {
		kind := kind

		s.Run(kind, func() {
			s.testCase(kind)
		})
	}
}

func TestEdge(t *testing.T) {
	suite.Run(t, new(EdgeTestSuite))
}
