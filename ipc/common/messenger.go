package common

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
)

// Messenger is
type Messenger struct {
	Name   string
	Nonce  uint32
	reader *bufio.Reader
	writer *bufio.Writer
}

// NewMessenger creates
func NewMessenger(name string, reader *bufio.Reader, writer *bufio.Writer) *Messenger {
	return &Messenger{
		Name:   name,
		reader: reader,
		writer: writer,
	}
}

// Send sends
func (messenger *Messenger) Send(message Message) error {
	messenger.Nonce++
	message.SetNonce(messenger.Nonce)

	dataBytes, err := messenger.marshal(message)
	if err != nil {
		return err
	}

	err = messenger.sendMessageLength(dataBytes)
	if err != nil {
		return err
	}

	_, err = messenger.writer.Write(dataBytes)
	if err != nil {
		return err
	}

	err = messenger.writer.Flush()
	fmt.Printf("[MSG %d] %s: SENT message of size %d\n", message.GetNonce(), messenger.Name, len(dataBytes))
	return err
}

func (messenger *Messenger) sendMessageLength(marshalizedMessage []byte) error {
	buffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(buffer, uint32(len(marshalizedMessage)))
	_, err := messenger.writer.Write(buffer)
	return err
}

// Receive receives
func (messenger *Messenger) Receive(message Message) error {
	fmt.Printf("%s: Receive message...\n", messenger.Name)

	// Wait for the start of a message
	messenger.blockingPeek(4)

	length, err := messenger.receiveMessageLength()
	if err != nil {
		return err
	}

	// Now read the body of [length]
	messenger.blockingPeek(length)
	buffer := make([]byte, length)
	_, err = io.ReadFull(messenger.reader, buffer)
	if err != nil {
		return err
	}

	err = messenger.unmarshal(buffer, message)
	if err != nil {
		return err
	}

	fmt.Printf("[MSG %d] %s: RECEIVED message of size %d\n", message.GetNonce(), messenger.Name, length)
	messageNonce := message.GetNonce()
	if messageNonce != messenger.Nonce+1 {
		return ErrInvalidMessageNonce
	}

	messenger.Nonce = messageNonce
	return nil
}

func (messenger *Messenger) receiveMessageLength() (int, error) {
	buffer := make([]byte, 4)
	_, err := io.ReadFull(messenger.reader, buffer)
	if err != nil {
		return 0, err
	}

	length := binary.LittleEndian.Uint32(buffer)
	return int(length), nil
}

func (messenger *Messenger) blockingPeek(n int) {
	for {
		_, err := messenger.reader.Peek(n)
		if err == nil {
			break
		}
	}
}

func (messenger *Messenger) marshal(data interface{}) ([]byte, error) {
	return marshalJSON(data)
}

func (messenger *Messenger) unmarshal(dataBytes []byte, data interface{}) error {
	return unmarshalJSON(dataBytes, data)
}

func marshalGob(data interface{}) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(data)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func unmarshalGob(dataBytes []byte, data interface{}) error {
	buffer := bytes.NewBuffer(dataBytes)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func marshalJSON(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func unmarshalJSON(dataBytes []byte, data interface{}) error {
	return json.Unmarshal(dataBytes, data)
}
