package kmtr

import "time"

const DEFAULT_TIMEOUT_MS = 800
const DEFAULT_PACKET_SIZE = 56
const DEFAULT_COUNT = 10
const DEFAULT_INTERVAL_MS = 10
const DEFAULT_TTL = 128

type PingReturn struct {
	destAddr string
	success  bool
	dropRate float64
	allTime  time.Duration
	bestTime time.Duration
	avgTime  time.Duration
	wrstTime time.Duration
}

type PingResult struct {
	success  bool
	succSum  int
	allTime  time.Duration
	bestTime time.Duration
	avgTime  time.Duration
	wrstTime time.Duration
}

type PingOptions struct {
	count      int
	timeoutMs  int
	intervalMs int
	packetSize int
}

func (options *PingOptions) Count() int {
	if options.count == 0 {
		options.count = DEFAULT_COUNT
	}
	return options.count
}

func (options *PingOptions) SetCount(count int) {
	options.count = count
}

func (options *PingOptions) TimeoutMs() int {
	if options.timeoutMs == 0 {
		options.timeoutMs = DEFAULT_TIMEOUT_MS
	}
	return options.timeoutMs
}

func (options *PingOptions) SetTimeoutMs(timeoutMs int) {
	options.timeoutMs = timeoutMs
}

func (options *PingOptions) IntervalMs() int {
	if options.intervalMs == 0 {
		options.intervalMs = DEFAULT_INTERVAL_MS
	}
	return options.intervalMs
}

func (options *PingOptions) SetIntervalMs(intervalMs int) {
	options.intervalMs = intervalMs
}

func (options *PingOptions) PacketSize() int {
	if options.packetSize == 0 {
		options.packetSize = DEFAULT_PACKET_SIZE
	}
	return options.packetSize
}

func (options *PingOptions) SetPacketSize(packetSize int) {
	options.packetSize = packetSize
}
