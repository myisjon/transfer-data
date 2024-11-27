// Created at: 2024/11/27
// Project : transfer-data

package config

import "time"
import kafka "github.com/segmentio/kafka-go"

type ReaderConfig struct {
    // The list of broker addresses used to connect to the kafka cluster.
    Brokers []string `json:"brokers,omitempty" yaml:"brokers"`

    // GroupID holds the optional consumer group id.  If GroupID is specified, then
    // Partition should NOT be specified e.g. 0
    GroupID string `json:"group_id,omitempty" yaml:"groupID"`

    // GroupTopics allows specifying multiple topics, but can only be used in
    // combination with GroupID, as it is a consumer-group feature. As such, if
    // GroupID is set, then either Topic or GroupTopics must be defined.
    GroupTopics []string `json:"group_topics,omitempty" yaml:"groupTopics"`

    // The topic to read messages from.
    Topic string `json:"topic,omitempty" yaml:"topic"`

    // Partition to read messages from.  Either Partition or GroupID may
    // be assigned, but not both
    Partition int `json:"partition,omitempty" yaml:"partition"`

    // An dialer used to open connections to the kafka server. This field is
    // optional, if nil, the default dialer is used instead.
    Dialer *kafka.Dialer `json:"dialer,omitempty" yaml:"dialer"`

    // The capacity of the internal message queue, defaults to 100 if none is
    // set.
    QueueCapacity int `json:"queue_capacity,omitempty" yaml:"queueCapacity"`

    // MinBytes indicates to the broker the minimum batch size that the consumer
    // will accept. Setting a high minimum when consuming from a low-volume topic
    // may result in delayed delivery when the broker does not have enough data to
    // satisfy the defined minimum.
    //
    // Default: 1
    MinBytes int `json:"min_bytes,omitempty" yaml:"minBytes"`

    // MaxBytes indicates to the broker the maximum batch size that the consumer
    // will accept. The broker will truncate a message to satisfy this maximum, so
    // choose a value that is high enough for your largest message size.
    //
    // Default: 1MB
    MaxBytes int `json:"max_bytes,omitempty" yaml:"maxBytes"`

    // Maximum amount of time to wait for new data to come when fetching batches
    // of messages from kafka.
    //
    // Default: 10s
    MaxWait time.Duration `json:"max_wait,omitempty" yaml:"maxWait"`

    // ReadBatchTimeout amount of time to wait to fetch message from kafka messages batch.
    //
    // Default: 10s
    ReadBatchTimeout time.Duration `json:"read_batch_timeout,omitempty" yaml:"readBatchTimeout"`

    // ReadLagInterval sets the frequency at which the reader lag is updated.
    // Setting this field to a negative value disables lag reporting.
    ReadLagInterval time.Duration `json:"read_lag_interval,omitempty" yaml:"readLagInterval"`

    // GroupBalancers is the priority-ordered list of client-side consumer group
    // balancing strategies that will be offered to the coordinator.  The first
    // strategy that all group members support will be chosen by the leader.
    //
    // Default: [Range, RoundRobin]
    //
    // Only used when GroupID is set
    GroupBalancers []kafka.GroupBalancer `json:"group_balancers,omitempty" yaml:"groupBalancers"`

    // HeartbeatInterval sets the optional frequency at which the reader sends the consumer
    // group heartbeat update.
    //
    // Default: 3s
    //
    // Only used when GroupID is set
    HeartbeatInterval time.Duration `json:"heartbeat_interval,omitempty" yaml:"heartbeatInterval"`

    // CommitInterval indicates the interval at which offsets are committed to
    // the broker.  If 0, commits will be handled synchronously.
    //
    // Default: 0
    //
    // Only used when GroupID is set
    CommitInterval time.Duration `json:"commit_interval,omitempty" yaml:"commitInterval"`

    // PartitionWatchInterval indicates how often a reader checks for partition changes.
    // If a reader sees a partition change (such as a partition add) it will rebalance the group
    // picking up new partitions.
    //
    // Default: 5s
    //
    // Only used when GroupID is set and WatchPartitionChanges is set.
    PartitionWatchInterval time.Duration `json:"partition_watch_interval,omitempty" yaml:"partitionWatchInterval"`

    // WatchForPartitionChanges is used to inform kafka-go that a consumer group should be
    // polling the brokers and rebalancing if any partition changes happen to the topic.
    WatchPartitionChanges bool `json:"watch_partition_changes,omitempty" yaml:"watchPartitionChanges"`

    // SessionTimeout optionally sets the length of time that may pass without a heartbeat
    // before the coordinator considers the consumer dead and initiates a rebalance.
    //
    // Default: 30s
    //
    // Only used when GroupID is set
    SessionTimeout time.Duration `json:"session_timeout,omitempty" yaml:"sessionTimeout"`

    // RebalanceTimeout optionally sets the length of time the coordinator will wait
    // for members to join as part of a rebalance.  For kafka servers under higher
    // load, it may be useful to set this value higher.
    //
    // Default: 30s
    //
    // Only used when GroupID is set
    RebalanceTimeout time.Duration `json:"rebalance_timeout,omitempty" yaml:"rebalanceTimeout"`

    // JoinGroupBackoff optionally sets the length of time to wait between re-joining
    // the consumer group after an error.
    //
    // Default: 5s
    JoinGroupBackoff time.Duration `json:"join_group_backoff,omitempty" yaml:"joinGroupBackoff"`

    // RetentionTime optionally sets the length of time the consumer group will be saved
    // by the broker
    //
    // Default: 24h
    //
    // Only used when GroupID is set
    RetentionTime time.Duration `json:"retention_time,omitempty" yaml:"retentionTime"`

    // StartOffset determines from whence the consumer group should begin
    // consuming when it finds a partition without a committed offset.  If
    // non-zero, it must be set to one of FirstOffset or LastOffset.
    //
    // Default: FirstOffset
    //
    // Only used when GroupID is set
    StartOffset int64 `json:"start_offset,omitempty" yaml:"startOffset"`

    // BackoffDelayMin optionally sets the smallest amount of time the reader will wait before
    // polling for new messages
    //
    // Default: 100ms
    ReadBackoffMin time.Duration `json:"read_backoff_min,omitempty" yaml:"readBackoffMin"`

    // BackoffDelayMax optionally sets the maximum amount of time the reader will wait before
    // polling for new messages
    //
    // Default: 1s
    ReadBackoffMax time.Duration `json:"read_backoff_max,omitempty" yaml:"readBackoffMax"`

    // If not nil, specifies a logger used to report internal changes within the
    // reader.
    Logger kafka.Logger `json:"logger,omitempty" yaml:"logger"`

    // ErrorLogger is the logger used to report errors. If nil, the reader falls
    // back to using Logger instead.
    ErrorLogger kafka.Logger `json:"error_logger,omitempty" yaml:"errorLogger"`

    // IsolationLevel controls the visibility of transactional records.
    // ReadUncommitted makes all records visible. With ReadCommitted only
    // non-transactional and committed records are visible.
    IsolationLevel kafka.IsolationLevel `json:"isolation_level,omitempty" yaml:"isolationLevel"`

    // Limit of how many attempts to connect will be made before returning the error.
    //
    // The default is to try 3 times.
    MaxAttempts int `json:"max_attempts,omitempty" yaml:"maxAttempts"`

    // OffsetOutOfRangeError indicates that the reader should return an error in
    // the event of an OffsetOutOfRange error, rather than retrying indefinitely.
    // This flag is being added to retain backwards-compatibility, so it will be
    // removed in a future version of kafka-go.
    OffsetOutOfRangeError bool `json:"offset_out_of_range_error,omitempty" yaml:"offsetOutOfRangeError"`
}
