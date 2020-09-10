package types

// NewSimplePacketData returns a new packet data
func NewSimplePacketData(h Header, payload []byte) PacketData {
	return PacketData{Header: h, Payload: payload}
}

// NewHeader returns a new header
func NewHeader() Header {
	return Header{}
}

// Get gets the first value associated with the given key.
func (h Header) Get(key string) *HeaderField {
	f, _ := h.keyIndex(key)
	return f
}

// Set sets the header fields associated with key to the
// single element value.
func (h *Header) Set(key string, value []byte) {
	f := HeaderField{Key: key, Value: value}
	_, idx := h.keyIndex(key)
	if idx < 0 {
		h.Fields = append(h.Fields, f)
	} else {
		h.setByIndex(f, idx)
	}
}

// Keys returns all keys in the header
func (h Header) Keys() []string {
	var keys []string
	for _, f := range h.Fields {
		keys = append(keys, f.Key)
	}
	return keys
}

func (h *Header) setByIndex(f HeaderField, idx int) {
	h.Fields[idx] = f
}

func (h Header) keyIndex(key string) (*HeaderField, int) {
	for i, f := range h.Fields {
		if f.Key == key {
			return &f, i
		}
	}
	return nil, -1
}
