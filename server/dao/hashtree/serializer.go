package hashtree

import (
	"github.com/topfreegames/pitaya/serialize"
	"github.com/topfreegames/pitaya/serialize/json"
	"github.com/topfreegames/pitaya/serialize/protobuf"
)

var (
	jsonSerializer  serialize.Serializer = json.NewSerializer()
	protoSerializer serialize.Serializer = protobuf.NewSerializer()
)
