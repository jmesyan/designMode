package main

//所有连接的抽象
type iConnect interface {
	close()
}
