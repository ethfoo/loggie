/*
Copyright 2021 Loggie Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package file

import (
	"fmt"
	"time"

	"loggie.io/loggie/pkg/core/api"
	"loggie.io/loggie/pkg/core/log"
	"loggie.io/loggie/pkg/core/result"
	"loggie.io/loggie/pkg/pipeline"
	"loggie.io/loggie/pkg/sink/codec"
)

const Type = "file"

func init() {
	pipeline.Register(api.SINK, Type, makeSink)
}

func makeSink(info pipeline.Info) api.Component {
	return NewSink()
}

type Sink struct {
	config *Config
	writer *MultiFileWriter
	cod    codec.Codec

	filenameMatcher [][]string
}

func NewSink() *Sink {
	return &Sink{
		config: &Config{},
	}
}

func (s *Sink) Config() interface{} {
	return s.config
}

func (s *Sink) SetCodec(c codec.Codec) {
	s.cod = c
}

func (s *Sink) Category() api.Category {
	return api.SINK
}

func (s *Sink) Type() api.Type {
	return Type
}

func (s *Sink) String() string {
	return fmt.Sprintf("%s/%s", api.SINK, Type)
}

func (s *Sink) Init(context api.Context) {
	s.filenameMatcher = codec.InitMatcher(s.config.Filename)
}

func (s *Sink) Start() {
	c := s.config
	w, err := NewMultiFileWriter(&Options{
		InitCapacity: 256,
		MaxSize:      c.MaxSize,
		MaxAge:       c.MaxAge,
		MaxBackups:   c.MaxBackups,
		LocalTime:    c.LocalTime,
		Compress:     c.Compress,
		IdleTimeout:  5 * time.Minute,
	})
	if err != nil {
		log.Panic("start multi file writer failed, error: %v", err)
	}

	s.writer = w

	log.Info("file-sink start,filename: %s", s.config.Filename)
}

func (s *Sink) Stop() {
	if s.writer != nil {
		_ = s.writer.Close()
	}
}

func (s *Sink) Consume(batch api.Batch) api.Result {
	events := batch.Events()
	l := len(events)
	if l == 0 {
		return nil
	}
	msgs := make([]Message, 0, l)
	for _, e := range events {
		filename, err := s.selectFilename(e)
		if err != nil {
			log.Error("select filename error: %+v", err)
			return result.Fail(err)
		}
		msgs = append(msgs, Message{
			Filename: filename,
			Data:     e.Body(),
		})
	}
	err := s.writer.Write(msgs...)
	if err != nil {
		log.Error("write to kafka error: %v", err)
		return result.Fail(err)
	}
	return result.Success()
}

func (s *Sink) selectFilename(e api.Event) (string, error) {
	if len(s.filenameMatcher) == 0 {
		return s.config.Filename, nil
	}
	msg, err := s.cod.Encode(e)
	if err != nil {
		log.Warn("encode event error: %+v", err)
		return "", err
	}
	return codec.PatternSelect(msg, s.config.Filename, s.filenameMatcher)
}
