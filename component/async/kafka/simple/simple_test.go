package simple

import (
	"errors"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/beatlabs/patron/component/async/kafka"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	brokers := []string{"192.168.1.1"}
	type args struct {
		name    string
		brokers []string
		topic   string
		options []kafka.OptionFunc
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "fails with missing name",
			args:    args{name: "", brokers: brokers, topic: "topic1"},
			wantErr: true,
		},
		{
			name:    "fails with missing brokers",
			args:    args{name: "test", brokers: []string{}, topic: "topic1"},
			wantErr: true,
		},
		{
			name:    "fails with one empty broker",
			args:    args{name: "test", brokers: []string{""}, topic: "topic1"},
			wantErr: true,
		},
		{
			name:    "fails with two brokers - one of the is empty",
			args:    args{name: "test", brokers: []string{" ", "broker2"}, topic: "topic1"},
			wantErr: true,
		},
		{
			name:    "fails with missing topics",
			args:    args{name: "test", brokers: brokers, topic: ""},
			wantErr: true,
		},
		{
			name:    "success",
			args:    args{name: "test", brokers: brokers, topic: "topic1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.name, tt.args.topic, tt.args.brokers, tt.args.options...)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

func TestFactory_Create(t *testing.T) {
	type fields struct {
		oo []kafka.OptionFunc
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "success", wantErr: false},
		{name: "failed with invalid option", fields: fields{oo: []kafka.OptionFunc{kafka.Buffer(-100)}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Factory{
				name:    "test",
				topic:   "topic",
				brokers: []string{"192.168.1.1"},
				oo:      tt.fields.oo,
			}
			got, err := f.Create()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

func TestWithDurationOffset(t *testing.T) {
	f := func(_ *sarama.ConsumerMessage) (time.Time, error) {
		return time.Time{}, nil
	}

	type args struct {
		since         time.Duration
		timeExtractor TimeExtractor
	}
	testCases := map[string]struct {
		args        args
		expectedErr error
	}{
		"success": {
			args: args{
				since:         time.Second,
				timeExtractor: f,
			},
		},
		"error - negative since duration": {
			args: args{
				since:         -time.Second,
				timeExtractor: f,
			},
			expectedErr: errors.New("duration must be positive"),
		},
		"error - nil time extractor": {
			args: args{
				since: time.Second,
			},
			expectedErr: errors.New("empty time extractor function"),
		},
	}
	for name, tt := range testCases {
		t.Run(name, func(t *testing.T) {
			c := kafka.ConsumerConfig{}
			err := WithDurationOffset(tt.args.since, tt.args.timeExtractor)(&c)
			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.True(t, c.DurationBasedConsumer)
				assert.Equal(t, time.Second, c.DurationOffset)
				assert.Equal(t,
					runtime.FuncForPC(reflect.ValueOf(tt.args.timeExtractor).Pointer()).Name(),
					runtime.FuncForPC(reflect.ValueOf(c.TimeExtractor).Pointer()).Name())
			}
		})
	}
}
