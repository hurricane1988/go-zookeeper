package main

// https://github.com/samuel/go-zookeeper
// 参考链接(https://www.cnblogs.com/zhichaoma/p/12640064.html)
//
import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"time"
)

// 全局变量
var path = "/test"

func main() {
	// 创建zk连接地址
	hosts := []string{"127.0.0.1:2181"}
	// 连接zk
	conn, _, err := zk.Connect(hosts, time.Second*5)
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}

	// 增加操作
	//zkAdd(conn, []byte("kubeSphere"), path)
	// 查询操作
	//zkGet(conn, []byte("test value"), path)
	// 修改操作
	//zkModify(conn, []byte("hell zookeeper"), path)
	// 删除操作
	zkDelete(conn, path)
}

// zk的增加操作
func zkAdd(conn *zk.Conn, data []byte, path string) {
	// flags有4种取值：
	// 0:永久，除非手动删除
	// zk.FlagEphemeral = 1:短暂，session断开则该节点也被删除
	// zk.FlagSequence  = 2:会自动在节点后面添加序号
	// 3:Ephemeral和Sequence，即，短暂且自动添加序号
	var flags int32 = 0
	// 获取访问控制权限
	acls := zk.WorldACL(zk.PermAll)
	s, err := conn.Create(path, data, flags, acls)
	if err != nil {
		log.Printf("创建失败,错误信息%s\n", err)
	}
	log.Printf("创建 %s成功\n", s)
}

// 查询操作
func zkGet(conn *zk.Conn, data []byte, path string) {
	data, _, err := conn.Get(path)
	if err != nil {
		fmt.Printf("查询%s失败, 错误信息: %s\n", path, err)
		return
	}
	fmt.Printf("%s的值为%s\n", path, string(data))
}

// 修改操作方法
func zkModify(conn *zk.Conn, newData []byte, path string) {
	_, sate, _ := conn.Get(path)
	_, err := conn.Set(path, newData, sate.Version)
	if err != nil {
		log.Printf("数据修改失败,错误信息: %s\n", err)
		return
	}
	log.Printf("数据%s查询成功\n", path)
}

// 删除操作
func zkDelete(conn *zk.Conn, path string) {
	_, sate, _ := conn.Get(path)
	err := conn.Delete(path, sate.Version)
	if err != nil {
		log.Printf("删除数据%s失败,错误信息: %s\n", path, err)
		return
	}
	log.Printf("删除数据%s成功!\n", path)
}
