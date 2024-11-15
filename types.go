package statmetrics

import (
	"encoding/json"
	"fmt"
)

type Stats struct {
	Name             string            `json:"name" metricHelp:"Handle instance name"`
	ClientID         string            `json:"client_id" metricHelp:"Client identifier"`
	Type             string            `json:"type" metricHelp:"Instance type (producer or consumer)"`
	Ts               int64             `json:"ts" metricHelp:"librdkafka's internal monotonic clock (microseconds)"`
	Time             int64             `json:"time" metricHelp:"Wall clock time in seconds since the epoch"`
	Age              int64             `json:"age" metricHelp:"Time since this client instance was created (microseconds)"`
	ReplyQ           int64             `json:"replyq" metricHelp:"Number of ops (callbacks, events, etc) waiting in queue for application to serve with rd_kafka_poll()"`
	MsgCnt           int64             `json:"msg_cnt" metricHelp:"Current number of messages in producer queues"`
	MsgSize          int64             `json:"msg_size" metricHelp:"Current total size of messages in producer queues"`
	MsgMax           int64             `json:"msg_max" metricHelp:"Threshold: maximum number of messages allowed on the producer queues"`
	MsgSizeMax       int64             `json:"msg_size_max" metricHelp:"Threshold: maximum total size of messages allowed on the producer queues"`
	Tx               int64             `json:"tx" metricHelp:"Total number of requests sent to Kafka brokers"`
	TxBytes          int64             `json:"tx_bytes" metricHelp:"Total number of message bytes (including framing, such as per-Message framing and MessageSet/batch framing) transmitted to Kafka brokers"`
	Rx               int64             `json:"rx" metricHelp:"Total number of responses received from Kafka brokers"`
	RxBytes          int64             `json:"rx_bytes" metricHelp:"Total number of bytes received from Kafka brokers"`
	TxMsgs           int64             `json:"txmsgs" metricHelp:"Total number of messages transmitted to Kafka brokers"`
	TxMsgsBytes      int64             `json:"txmsgs_bytes" metricHelp:"Total number of message bytes (including framing, such as per-Message framing and MessageSet/batch framing) transmitted to Kafka brokers"`
	RxMsgs           int64             `json:"rxmsgs" metricHelp:"Total number of messages consumed, not including ignored messages (due to offset, etc), from Kafka brokers"`
	RxMsgsBytes      int64             `json:"rxmsgs_bytes" metricHelp:"Total number of message bytes (including framing) received from Kafka brokers"`
	SimpleCnt        int64             `json:"simple_cnt" metricHelp:"Internal tracking of legacy vs new consumer API state"`
	MetadatacacheCnt int64             `json:"metadata_cache_cnt" metricHelp:"Number of topics in the metadata cache"`
	Brokers          map[string]Broker `json:"brokers" metricHelp:"Brokers handled by this client instance"`
	Topics           map[string]Topic  `json:"topics" metricHelp:"Topics handled by this client instance"`
	Cgrps            map[string]Cgrp   `json:"cgrps" metricHelp:"Consumer groups handled by this client instance"`
	Eos              Eos               `json:"eos" metricHelp:"EOS / Idempotent producer state and metrics"`
}

type Broker struct {
	Name           string            `json:"name" metricHelp:"Broker hostname, port and broker id"`
	NodeName       string            `json:"nodename" metricHelp:"Broker hostname"`
	Source         string            `json:"source" metricHelp:"Broker source (learned, configured, internal, logical)"`
	State          string            `json:"state" metricHelp:"Broker state (INIT, DOWN, CONNECT, AUTH, APIVERSION_QUERY, AUTH_HANDSHAKE, UP, UPDATE)"`
	NodeID         int64             `json:"nodeid" metricHelp:"Broker id (-1 for bootstraps)"`
	StateAge       int64             `json:"stateage" metricHelp:"Time since last broker state change (microseconds)"`
	OutbufCnt      int64             `json:"outbuf_cnt" metricHelp:"Number of requests awaiting transmission to broker"`
	OutbufMsgCnt   int64             `json:"outbuf_msg_cnt" metricHelp:"Number of messages awaiting transmission to broker"`
	WaitRespCnt    int64             `json:"waitresp_cnt" metricHelp:"Number of requests in-flight to broker awaiting response"`
	WaitRespMsgCnt int64             `json:"waitresp_msg_cnt" metricHelp:"Number of messages in-flight to broker awaiting response"`
	Tx             int64             `json:"tx" metricHelp:"Total number of requests sent"`
	TxBytes        int64             `json:"txbytes" metricHelp:"Total number of bytes sent"`
	TxErrs         int64             `json:"txerrs" metricHelp:"Total number of transmission errors"`
	TxRetries      int64             `json:"txretries" metricHelp:"Total number of request retries"`
	TxIdle         int64             `json:"txidle" metricHelp:"Microseconds since last socket send (or -1 if no sends yet for current connection)"`
	ReqTimeouts    int64             `json:"req_timeouts" metricHelp:"Total number of requests timed out"`
	Rx             int64             `json:"rx" metricHelp:"Total number of responses received"`
	RxBytes        int64             `json:"rxbytes" metricHelp:"Total number of bytes received"`
	RxErrs         int64             `json:"rxerrs" metricHelp:"Total number of receive errors"`
	RxCorriderrs   int64             `json:"rxcorriderrs" metricHelp:"Total number of unmatched correlation ids in response (typically for timed out requests)"`
	RxPartial      int64             `json:"rxpartial" metricHelp:"Total number of partial MessageSets received. The broker may return partial responses if the full MessageSet could not fit in the remaining Fetch response size"`
	RxIdle         int64             `json:"rxidle" metricHelp:"Microseconds since last socket receive (or -1 if no receives yet for current connection)"`
	ZBufGrow       int64             `json:"zbuf_grow" metricHelp:"Total number of decompression buffer size increases"`
	BufGrow        int64             `json:"buf_grow" metricHelp:"Total number of buffer size increases (deprecated, unused)"`
	Wakeups        int64             `json:"wakeups" metricHelp:"Broker thread poll loop wakeups"`
	Connects       int64             `json:"connects" metricHelp:"Number of connection attempts, including successful and failed, and name resolution failures"`
	Disconnects    int64             `json:"disconnects" metricHelp:"Number of disconnects (triggered by broker, network, load-balancer, etc.)"`
	IntLatency     WinStats          `json:"int_latency" metricHelp:"Internal producer queue latency in microseconds"`
	OutbufLatency  WinStats          `json:"outbuf_latency" metricHelp:"Internal request queue latency in microseconds. This is the time between a request is enqueued on the transmit (outbuf) queue and the time the request is written to the TCP socket. Additional buffering and latency may be incurred by the TCP stack and network"`
	Rtt            WinStats          `json:"rtt" metricHelp:"Broker latency / round-trip time in microseconds"`
	Throttle       WinStats          `json:"throttle" metricHelp:"Broker throttling time in milliseconds"`
	TopPars        map[string]TopPar `json:"toppars" metricHelp:"Partitions handled by this broker handle"`
}

type WinStats struct {
	Min        int64 `json:"min"`
	Max        int64 `json:"max"`
	Avg        int64 `json:"avg"`
	Sum        int64 `json:"sum"`
	Cnt        int64 `json:"cnt"`
	StdDev     int64 `json:"stddev"`
	HdrSize    int64 `json:"hdrsize"`
	P50        int64 `json:"p50"`
	P75        int64 `json:"p75"`
	P90        int64 `json:"p90"`
	P95        int64 `json:"p95"`
	P99        int64 `json:"p99"`
	P9999      int64 `json:"p99_99"`
	OutOfRange int64 `json:"outofrange"`
}

type TopPar struct {
	Topic     string `json:"topic"`
	Partition int64  `json:"partition"`
}

type Topic struct {
	Topic       string               `json:"topic"`
	Age         int64                `json:"age"`
	MetadataAge int64                `json:"metadata_age"`
	BatchSize   WinStats             `json:"batch_size"`
	BatchCnt    WinStats             `json:"batch_cnt"`
	Partitions  map[string]Partition `json:"partitions"`
}

type Partition struct {
	Partition            int64  `json:"partition"`
	Broker               int64  `json:"broker"`
	Leader               int64  `json:"leader"`
	Desired              bool   `json:"desired"`
	Unknown              bool   `json:"unknown"`
	MsgQCount            int64  `json:"msgq_count"`
	MsgQBytes            int64  `json:"msgq_bytes"`
	XmitMsgQCnt          int64  `json:"xmit_msgq_cnt"`
	XmitMsgQBytes        int64  `json:"xmit_msgq_bytes"`
	FetchQCnt            int64  `json:"fetchq_cnt"`
	FetchQSize           int64  `json:"fetchq_size"`
	FetchState           string `json:"fetch_state"`
	QueryOffset          int64  `json:"query_offset"`
	NextOffset           int64  `json:"next_offset"`
	AppOffset            int64  `json:"app_offset"`
	StoredOffset         int64  `json:"stored_offset"`
	StoredLeaderEpoch    int64  `json:"stored_leader_epoch"`
	CommitedOffset       int64  `json:"commited_offset"`
	CommittedLeaderEpoch int64  `json:"committed_leader_epoch"`
	EofOffset            int64  `json:"eof_offset"`
	LoOffset             int64  `json:"lo_offset"`
	HiOffset             int64  `json:"hi_offset"`
	LsOffset             int64  `json:"ls_offset"`
	ConsumerLag          int64  `json:"consumer_lag"`
	ConsumerLagStored    int64  `json:"consumer_lag_stored"`
	LeaderEpoch          int64  `json:"leader_epoch"`
	TxMsgs               int64  `json:"txmsgs"`
	Txbytes              int64  `json:"txbytes"`
	RxMsgs               int64  `json:"rxmsgs"`
	RxBytes              int64  `json:"rxbytes"`
	Msgs                 int64  `json:"msgs"`
	RxVerDrops           int64  `json:"rx_ver_drops"`
	MsgsInflight         int64  `json:"msgs_inflight"`
	NextAckSeq           int64  `json:"next_ack_seq"`
	NextErrSeq           int64  `json:"next_err_seq"`
	AckedMsgId           int64  `json:"acked_msgid"`
}

type Cgrp struct {
	State           string `json:"state"`
	StateAgeMs      int64  `json:"stateage"`
	JoinState       string `json:"join_state"`
	RebalanceAgeMs  int64  `json:"rebalance_age"`
	RebalanceCnt    int64  `json:"rebalance_cnt"`
	RebalanceReason string `json:"rebalance_reason"`
	AssignmentSize  int64  `json:"assignment_size"`
}

type Eos struct {
	IdempState      string `json:"idemp_state"`
	IdempStateAgeMs int64  `json:"idemp_stateage"`
	TxnState        string `json:"txn_state"`
	TxnStateAgeMs   int64  `json:"txn_stateage"`
	TxnMayEnq       bool   `json:"txn_may_enq"`
	ProducerID      int64  `json:"producer_id"`
	ProducerEpoch   int64  `json:"producer_epoch"`
	EpochCnt        int64  `json:"epoch_cnt"`
}

func (s *Stats) GetMetricHelp() map[string]string {
	return map[string]string{
		"name":               "Handle instance name",
		"client_id":          "Client identifier",
		"type":               "Instance type (producer or consumer)",
		"ts":                 "librdkafka's internal monotonic clock (microseconds)",
		"time":               "Wall clock time in seconds since the epoch",
		"age":                "Time since this client instance was created (microseconds)",
		"replyq":             "Number of ops (callbacks, events, etc) waiting in queue for application to serve with rd_kafka_poll()",
		"msg_cnt":            "Current number of messages in producer queues",
		"msg_size":           "Current total size of messages in producer queues",
		"msg_max":            "Threshold: maximum number of messages allowed on the producer queues",
		"msg_size_max":       "Threshold: maximum total size of messages allowed on the producer queues",
		"tx":                 "Total number of requests sent to Kafka brokers",
		"tx_bytes":           "Total number of message bytes (including framing, such as per-Message framing and MessageSet/batch framing) transmitted to Kafka brokers",
		"rx":                 "Total number of responses received from Kafka brokers",
		"rx_bytes":           "Total number of bytes received from Kafka brokers",
		"txmsgs":             "Total number of messages transmitted to Kafka brokers",
		"txmsgs_bytes":       "Total number of message bytes (including framing, such as per-Message framing and MessageSet/batch framing) transmitted to Kafka brokers",
		"rxmsgs":             "Total number of messages consumed, not including ignored messages (due to offset, etc), from Kafka brokers",
		"rxmsgs_bytes":       "Total number of message bytes (including framing) received from Kafka brokers",
		"simple_cnt":         "Internal tracking of legacy vs new consumer API state",
		"metadata_cache_cnt": "Number of topics in the metadata cache",
		"brokers":            "Brokers handled by this client instance",
		"topics":             "Topics handled by this client instance",
		"cgrps":              "Consumer groups handled by this client instance",
		"eos":                "EOS / Idempotent producer state and metrics",
	}
}

func (b *Broker) GetMetricHelp() map[string]string {
	return map[string]string{
		"name":             "Broker hostname, port and broker id",
		"nodename":         "Broker hostname",
		"source":           "Broker source (learned, configured, internal, logical)",
		"state":            "Broker state (INIT, DOWN, CONNECT, AUTH, APIVERSION_QUERY, AUTH_HANDSHAKE, UP, UPDATE)",
		"nodeid":           "Broker id (-1 for bootstraps)",
		"stateage":         "Time since last broker state change (microseconds)",
		"outbuf_cnt":       "Number of requests awaiting transmission to broker",
		"outbuf_msg_cnt":   "Number of messages awaiting transmission to broker",
		"waitresp_cnt":     "Number of requests in-flight to broker awaiting response",
		"waitresp_msg_cnt": "Number of messages in-flight to broker awaiting response",
		"tx":               "Total number of requests sent",
		"txbytes":          "Total number of bytes sent",
		"txerrs":           "Total number of transmission errors",
		"txretries":        "Total number of request retries",
		"txidle":           "Microseconds since last socket send (or -1 if no sends yet for current connection)",
		"req_timeouts":     "Total number of requests timed out",
		"rx":               "Total number of responses received",
		"rxbytes":          "Total number of bytes received",
		"rxerrs":           "Total number of receive errors",
		"rxcorriderrs":     "Total number of unmatched correlation ids in response (typically for timed out requests)",
		"rxpartial":        "Total number of partial MessageSets received. The broker may return partial responses if the full MessageSet could not fit in the remaining Fetch response size",
		"rxidle":           "Microseconds since last socket receive (or -1 if no receives yet for current connection)",
		"zbuf_grow":        "Total number of decompression buffer size increases",
		"buf_grow":         "Total number of buffer size increases (deprecated, unused)",
		"wakeups":          "Broker thread poll loop wakeups",
		"connects":         "Number of connection attempts, including successful and failed, and name resolution failures",
		"disconnects":      "Number of disconnects (triggered by broker, network, load-balancer, etc.)",
		"int_latency":      "Internal producer queue latency in microseconds",
		"outbuf_latency":   "Internal request queue latency in microseconds. This is the time between a request is enqueued on the transmit (outbuf) queue and the time the request is written to the TCP socket. Additional buffering and latency may be incurred by the TCP stack and network",
		"rtt":              "Broker latency / round-trip time in microseconds",
		"throttle":         "Broker throttling time in milliseconds",
		"toppars":          "Partitions handled by this broker handle",
	}
}

func ParseStats(jdata []byte) (*Stats, error) {
	var s Stats
	err := json.Unmarshal(jdata, &s)
	if err != nil {
		err = fmt.Errorf("Failed to parse stats: %s", err)
		return nil, err
	}
	return &s, nil
}
