package xmlcodec

import (
	"encoding/xml"

	"code.google.com/p/go.net/websocket"
)

func xmlMarshal(v interface{}) (msg []byte, payloadType byte, err error) {
	// buff := &bytes.Buffer{}
	msg, err = xml.Marshal(v)
	// msgRet := buff.Bytes()
	return msg, websocket.TextFrame, nil
}

func xmlUnmarshal(msg []byte, payloadType byte, v interface{}) (err error) {
	err = xml.Unmarshal(msg, v)
	return err
}

var XMLCodec = websocket.Codec{xmlMarshal, xmlUnmarshal}
