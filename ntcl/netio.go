package ntcl

// netio contains functions relating to network stack

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const DELIM = '|'

func EncodeMsg(content string) string {
	return content + string(DELIM)
}

//TODO! factor and replace old
func NtwkWrite(ntchan Ntchan, content string) (int, error) {
	//READLINE uses \n
	NEWLINE := '\n'
	//respContent := fmt.Sprintf("%s%c%c", content, DELIM, NEWLINE)
	respContent := fmt.Sprintf("%s%c", content, NEWLINE)
	//log.Println("write > ", content, respContent)
	writer := bufio.NewWriter(ntchan.Conn)
	n, err := writer.WriteString(respContent)
	if err == nil {
		err = writer.Flush()
	}
	s := fmt.Sprintf("bytes written", n, " ", ntchan.SrcName, ntchan.DestName)
	vlog(s)
	return n, err
}

func NtwkRead(ntchan Ntchan, delim byte) (string, error) {
	//log.Println("NtwkRead ", ntchan.SrcName, ntchan.DestName)
	reader := bufio.NewReader(ntchan.Conn)
	var buffer bytes.Buffer
	for {
		//READLINE uses \n
		ba, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		buffer.Write(ba)
		if !isPrefix {
			break
		}
	}
	return buffer.String(), nil
}

func MsgRead(ntchan Ntchan) (string, error) {
	msg_string, err := NtwkRead(ntchan, DELIM)
	msg_string = strings.Trim(msg_string, string(DELIM))
	return msg_string, err
}
