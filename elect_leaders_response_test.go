//go:build !functional

package sarama

import "testing"

var (
	electLeadersResponseOneTopicV1 = []byte{
		0, 0, 3, 232, // ThrottleTimeMs 1000
		0, 0, // errorCode
		0, 0, 0, 1, // number of topics
		0, 5, 116, 111, 112, 105, 99, // topic name "topic"
		0, 0, 0, 1, // number of partitions
		0, 0, 0, 0, // partition 0
		0, 0, // partition errorCode
		255, 255, // nil errorMessage
	}
	electLeadersResponseOneTopicV2 = []byte{
		0, 0, 3, 232, // ThrottleTimeMs 1000
		0, 0, // errorCode
		2,                         // 2-1=1 topic
		6, 116, 111, 112, 105, 99, // topic name "topic"
		2,          // 2-1=1 partition
		0, 0, 0, 0, // partition 0
		0, 0, // partition errorCode
		0, // nil errorMessage
		0, // empty tagged fields for partition result
		0, // empty tagged fields for topic result
		0, // empty tagged fields for response
	}
)

func TestElectLeadersResponse(t *testing.T) {
	var response = &ElectLeadersResponse{
		Version:        int16(1),
		ThrottleTimeMs: int32(1000),
		ReplicaElectionResults: map[string]map[int32]*PartitionResult{
			"topic": {
				0: {},
			},
		},
	}

	testResponse(t, "one topic V1", response, electLeadersResponseOneTopicV1)

	response.Version = 2
	testResponse(t, "one topic V2", response, electLeadersResponseOneTopicV2)
}
