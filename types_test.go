package statmetrics

import (
	"testing"
)

func TestParseStats(t *testing.T) {
	validJSON := `{
		"name": "test_instance",
		"client_id": "test_client",
		"type": "producer",
		"ts": 123456789,
		"time": 1609459200,
		"age": 1000000,
		"replyq": 10,
		"msg_cnt": 100,
		"msg_size": 2048,
		"msg_max": 1000,
		"msg_size_max": 1048576,
		"tx": 50,
		"tx_bytes": 102400,
		"rx": 45,
		"rx_bytes": 51200,
		"txmsgs": 40,
		"txmsgs_bytes": 40960,
		"rxmsgs": 35,
		"rxmsgs_bytes": 35840,
		"simple_cnt": 1,
		"metadata_cache_cnt": 5,
		"brokers": {},
		"topics": {},
		"cgrps": {},
		"eos": {}
	}`

	invalidJSON := `{
		"name": "test_instance",
		"client_id": "test_client",
		"type": "producer",
		"ts": "invalid_timestamp"
	}`

	t.Run("valid JSON", func(t *testing.T) {
		stats, err := ParseStats([]byte(validJSON))
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if stats.Name != "test_instance" {
			t.Errorf("Expected Name to be 'test_instance', got %s", stats.Name)
		}
		if stats.ClientID != "test_client" {
			t.Errorf("Expected ClientID to be 'test_client', got %s", stats.ClientID)
		}
		if stats.Type != "producer" {
			t.Errorf("Expected Type to be 'producer', got %s", stats.Type)
		}
		if stats.Ts != 123456789 {
			t.Errorf("Expected Ts to be 123456789, got %d", stats.Ts)
		}
	})

	t.Run("invalid JSON", func(t *testing.T) {
		_, err := ParseStats([]byte(invalidJSON))
		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}
