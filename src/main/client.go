package main

import (
    "net"
)

type Client struct {
    conn     *net.Conn
    player   Player
}

type Player struct {
    name        string
    x           int
    y           int
    inventory   []int
}
