package sonic_fast

import (
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	sonic "github.com/bytedance/sonic"
)

type SonicSerializer struct{}

func (s SonicSerializer) TimePrecision() time.Duration {
	return time.Millisecond
}

func (s SonicSerializer) Marshal(o interface{}) ([]byte, error) {
	return sonic.ConfigFastest.Marshal(o)
}

func (m SonicSerializer) Unmarshal(d []byte, o interface{}) error {
	return sonic.ConfigFastest.Unmarshal(d, o)
}

func NewSonicSerializer() goserbench.Serializer {
	return SonicSerializer{}
}