package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"nursing_work/model"
	"strconv"
)

var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Msg struct {
	Type int
	Data []byte
}

type WsConnection struct {
	inChan    chan *Msg
	outChan   chan *Msg
	Socket    *websocket.Conn
	ID        string
	UserId    int
	isClosed  chan int
	closeChan chan byte
}

var Chs map[string]chan string

var Conns map[string]*WsConnection

func Chat(c *gin.Context) {
	wsSocket, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "Upgrade failed",
		})
		return
	}

	tmp, _ := c.Get("openID")
	Mid := model.UserID(fmt.Sprint(tmp))
	userId, err := strconv.Atoi(Mid)
	if err != nil {
		fmt.Println(err)
	}
	Oid := c.Query("id_object")
	fmt.Println(Mid, Oid)
	// m : me ; o : object
	om := fmt.Sprintf("%vTo%v", Oid, Mid)
	mo := fmt.Sprintf("%vTo%v", Mid, Oid)
	Conn := &WsConnection{
		inChan:    make(chan *Msg, 1000),
		outChan:   make(chan *Msg, 1000),
		Socket:    wsSocket,
		closeChan: make(chan byte),
		ID:        mo,
		UserId:    userId,
	}
	Conns[mo] = Conn
	if Chs[mo] == nil {
		Chs[mo] = make(chan string, 10)
	}
	Chs[mo] <- "needing"
	if Chs[om] == nil {
		Conn.Socket.WriteMessage(1, []byte("对方尚未未上线."))
		Chs[om] = make(chan string, 10)
		fmt.Println()
	}
	for {
		select {
		case <-Chs[om]:
			Conn.Socket.WriteMessage(1, []byte("对方在线ing."))
			goto HH
		}
	}
HH:
	oConn := Conns[om]
	go oConn.proLoop(Conn)
	go oConn.Read(Conn)
	go oConn.Write(Conn)
	//go oConn.W()
	return

}

func (oConn *WsConnection) proLoop(mConn *WsConnection) {
	//go func() {
	//不断的read
	for {
		select {
		case msg := <-oConn.inChan:
			mConn.outChan <- msg
		}
	}
	//}()
	//不断的write
	//for {
	//	select {
	//	case msg1 := <-oConn.writeChan:
	//		Conn.outChan <- msg1
	//	}
	//}
}

func (oConn *WsConnection) Read(mConn *WsConnection) {
	for {
		msgType, data, err := oConn.Socket.ReadMessage()
		if err != nil {
			Chs[oConn.ID] = nil
			Chs[mConn.ID] = nil
			Conns[oConn.ID] = nil
			Conns[mConn.ID] = nil
			oConn.Socket.Close()
			mConn.Socket.Close()
			oConn.isClosed <- 1
			mConn.isClosed <- 1
			return
		}
		msg := &Msg{
			Data: data,
			Type: msgType,
		}
		//存到数据库

		err = model.InsertOneMsg(&model.Message{
			Mid:  oConn.UserId,
			Oid:  mConn.UserId,
			Data: string(data),
			Type: msgType,
		})

		if err != nil {
			fmt.Println(err)
		}
		oConn.inChan <- msg
	}
}

func (oConn *WsConnection) Write(mConn *WsConnection) {
	for {
		msg := <-mConn.outChan
		err := mConn.Socket.WriteMessage(msg.Type, msg.Data)
		if err != nil {
			mConn.Socket.Close()
			return
		}
	}
}
