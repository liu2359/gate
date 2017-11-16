package main

import (
	"github.com/golang/protobuf/proto"

	"util/logs"

	"core/net/dispatcher"
	"core/net/dispatcher/pb"
	"core/net/lan"

	"share"
)

var _ = logs.Debug

//
var (
	g_lan *lan.Lan
)

// 服务器间相关处理
func InitServers() {
	//
	share.InitLans(Cfg().LanCfg, Cfg().EtcdCfg, func(f *pb.PbFrame) {
		dispatcher.Dispatch(f, func(dstUrl string) {
			// 通知offline
			NoticeServerOffline(dstUrl, *f.SrcUrl)
		})
	})
}

func ToServer(c *Client, dstUrl string, d []byte) bool {
	// message
	f := &pb.PbFrame{
		SrcUrl:  proto.String(c.Url),
		DstUrls: []string{dstUrl},
		AccId:   proto.Int64(c.AccId),
		MsgRaw:  d,
		Offline: proto.Bool(false),
	}

	return share.SendFrame2Server(dstUrl, f)
}

//
func NoticeServerOffline(srcUrl, dstUrl string) bool {
	// message
	f := &pb.PbFrame{
		SrcUrl:  proto.String(srcUrl),
		DstUrls: []string{dstUrl},
		Offline: proto.Bool(true),
	}

	return share.SendFrame2Server(dstUrl, f)
}
