// package lognsq provides redirection of messages to a nsqd TCP service.
// It is used, for example, for aggregating logs from several server instances.
package lognsq

import (
	"errors"
	"os"

	"github.com/nsqio/go-nsq"
)

// Config provides data necessary to start NSQ session.
// NSQ is a distributed messaging service.
type Config struct {
	// PrintLogs enables printing of logs to STDERR.
	PrintLogs bool

	// Topic sets a namespace for messages.
	Topic string

	// NsqdURL is the URL to http service of an nsqd daemon
	NsqdURL string
}

// LogNSQ creates a "producer" to the nsqd service. The producer is able to
// publish messages to nsqd. It cannot consume the messages.
// If the same config is used for many different apps, all of them are able
// to contribute their logs to the same namespace using nsqd.
type LogNSQ struct {
	Config
	*nsq.Producer
}

// New Creates a new LogNSQ instance. If creation of "producer" failed, it
// returns an error.
func New(cfg Config) (l *LogNSQ, err error) {
	var prod *nsq.Producer
	if cfg.Topic == "" {
		err = errors.New("config for LogNSQ cannot have an empty Topic field")
		return nil, err
	}

	nsqCfg := nsq.NewConfig()
	prod, err = nsq.NewProducer(cfg.NsqdURL, nsqCfg)
	if err != nil {
		return nil, err
	}

	l = &LogNSQ{
		Config:   cfg,
		Producer: prod,
	}
	return l, err
}

// Write takes a slice of bytes and publishes it to STDERR as well as to
// nsqd service. It uses Topic given in the config.
func (l *LogNSQ) Write(bs []byte) (n int, err error) {
	if l.PrintLogs {
		n, err = os.Stderr.Write(bs)
	}
	if err == nil {
		err = l.Publish(l.Topic, bs)
	}
	return n, err
}
