package net

import (
	"bytes"
	"encoding/binary"
	"github.com/Sirupsen/logrus"
	"net"
)

const (
	MaxPacketSize = 1024 * 16
)

type GameClient struct {
	conn    net.Conn
	handler PacketHandler
	db      PacketDatabase
	log     *logrus.Entry
}

func NewGameClient(conn net.Conn, handler PacketHandler, db PacketDatabase) *GameClient {
	return &GameClient{
		conn:    conn,
		handler: handler,
		db:      db,
		log:     logrus.WithField("component", "client"),
	}
}

func (c *GameClient) Start() {
	go c.run()
}

func (c *GameClient) Disconnect() error {
	return c.conn.Close()
}

func (c *GameClient) run() {
	var packet uint16
	var size uint16

	raw := make([]byte, MaxPacketSize)
	buffer := bytes.NewBuffer(raw)
	offset := 0
	state := 0

	for true {
		read, err := c.conn.Read(raw[offset:])

		if err != nil {
			c.log.WithError(err).Error("error reading from socket")
			return
		}

		offset += read

		if state == 0 {
			if offset >= 2 {
				err := binary.Read(buffer, binary.LittleEndian, &packet)

				if err != nil {
					c.log.WithError(err).Error("error reading from socket")
					return
				}

				state++
			}
		}

		if state == 1 {
			size = c.db[packet]

			if size == 0xFFFF {
				if offset >= 4 {
					err := binary.Read(buffer, binary.LittleEndian, &size)

					if err != nil {
						c.log.WithError(err).Error("error reading from socket")
						return
					}
				}
			} else {
				state++
			}
		}

		if state == 2 {
			if offset >= int(size) {
				p := &Packet{
					Packet: packet,
					Size:   size,
					Data:   bytes.NewBuffer(raw[:size]),
				}

				c.handler(p)

				state = 0
				offset = 0
				buffer.Reset()
			}
		}
	}
}