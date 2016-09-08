package push

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

// Dispatch
func (i *ios) Dispatch(result []map[string]string, pushContent, pushType, pushAction string) {
	resultLength := len(result)
	ch := make([]chan int, resultLength)
	for k, v := range result {
		ch[k] = make(chan int)
		go i.Comply(ch[k], v["device"], pushContent, pushType, pushAction)
	}
	for _, c := range ch {
		<-c // 读取数据
	}
}

// Comply
func (i *ios) Comply(ch chan int, pushToken, pushContent, pushType, pushAction string) {
	err := i.task(pushToken, pushContent, pushType, pushAction)
	if err != nil {
		return
	}
	ch <- 1 // channel 写入数据
}

// pushDevice
func (i *ios) task(pushToken, pushContent, pushType, pushAction string) error {
	message, err := i.message(pushToken, pushContent, pushType, pushAction)
	if err != nil {
		return err
	}
	conn, err := i.server()
	if err != nil {
		return err
	}
	err = i.push(conn, message)
	return err
}

// pushData 推送数据
func (i *ios) message(pushToken, pushContent, pushType, pushAction string) ([]byte, error) {
	// 设备
	tokenByte, err := hex.DecodeString(pushToken)
	if err != nil {
		return nil, err
	}
	// 负载
	payload := map[string]interface{}{"aps": map[string]string{"alert": pushContent, "type": pushType, "action": pushAction}}
	payloadByte, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	// 缓冲
	buffer := new(bytes.Buffer)
	// 推送数据
	err = binary.Write(buffer, binary.BigEndian, struct {
		cmd       uint8  // 命令
		id        uint32 // 传输 id，optional
		timestamp uint32 // 过期时间，1小时
	}{uint8(1), uint32(1), uint32(time.Second + 60*60)})
	if err != nil {
		return nil, err
	}
	// 推送设备令牌
	err = binary.Write(buffer, binary.BigEndian, uint16(len(tokenByte)))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.BigEndian, tokenByte)
	if err != nil {
		return nil, err
	}
	// 推送 payload
	err = binary.Write(buffer, binary.BigEndian, uint16(len(payloadByte)))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buffer, binary.BigEndian, payloadByte)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// pushServer 推送服务器
func (i *ios) server() (*tls.Conn, error) {
	// 证书私钥
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		return nil, err
	}
	// 配置文件
	conf := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "gateway.push.apple.com",
	}
	// 连接
	conn, err := net.Dial("tcp", "gateway.push.apple.com:2195")
	if err != nil {
		return nil, err
	}
	// 加密
	connTLS := tls.Client(conn, conf)
	// 握手
	err = connTLS.Handshake()
	if err != nil {
		return nil, err
	}
	return connTLS, nil
}

// pushMessage 推送消息
func (i *ios) push(conn *tls.Conn, message []byte) (err error) {
	_, err = conn.Write(message)
	if err != nil {
		return err
	}
	_, err = i.result(conn)
	if err != nil {
		return err
	}
	return nil
}

// pushResult 推送结果
func (i *ios) result(conn *tls.Conn) ([]byte, error) {
	// 接收返回信息
	chResult := make(chan []byte)
	// 接收错误信息
	chError := make(chan error)
	go func(chResult chan []byte, chError chan error) {
		data := make([]byte, 6, 6)
		_, err := conn.Read(data)
		if err != nil {
			chError <- err
			return
		}
		chResult <- data
	}(chResult, chError)
	var response []byte
	var err error
	select {
	// 处理响应
	case response = <-chResult:
		fmt.Println("推送未知", string(response))
		// 处理错误
	case err = <-chError:
		fmt.Println("推送失败", err.Error())
	// 处理成功
	case <-time.Tick(5 * time.Second):
		fmt.Println("推送成功")
	}
	return response, err
}
