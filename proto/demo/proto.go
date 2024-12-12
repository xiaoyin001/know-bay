package main

import (
	"github.com/emicklei/proto"
)

var (
	// lua的存放路径
	uLuaPath string = ""
	// .proto 中 PackageName
	uProtoPackageName string = ""
	// .proto 中所有的Service
	uProtoServices []protoService = make([]protoService, 0)
)

// 初始化数据
func initCommonParam() {
	uLuaPath = ""
	uProtoPackageName = ""
	uProtoServices = make([]protoService, 0)
}

// 以Option为粒度，处理每个Option定义
func handleOption(o *proto.Option) {
	if o.Name != "lua_package" {
		return
	}

	if !o.Constant.IsString {
		return
	}

	if o.Constant.Source == "" {
		return
	}

	uLuaPath = o.Constant.Source + "/"
	ensureDir(uLuaPath)
}

// 以Package为粒度，处理每个Package定义
func handlePackage(p *proto.Package) {
	uProtoPackageName = p.Name
}

// 以Service为粒度，处理每个Service定义
func handleService(s *proto.Service) {
	mRPCs := make([]protoRPC, 0)

	for _, serviceEach := range s.Elements {
		mTempRPC, ok := serviceEach.(*proto.RPC)
		if !ok {
			continue
		}

		mRPCComment := make([]string, 0)
		if mTempRPC.Comment != nil {
			mRPCComment = mTempRPC.Comment.Lines
		}

		mRPCInlineComment := make([]string, 0)
		if mTempRPC.InlineComment != nil {
			mRPCInlineComment = mTempRPC.InlineComment.Lines
		}

		mRPC := protoRPC{
			Name:          mTempRPC.Name,
			Comment:       mRPCComment,
			InlineComment: mRPCInlineComment,
			RequestType:   mTempRPC.RequestType,
			ResponseType:  mTempRPC.ReturnsType,
		}

		mRPCs = append(mRPCs, mRPC)
	}

	mSvrComment := make([]string, 0)
	if s.Comment != nil {
		mSvrComment = s.Comment.Lines
	}

	mService := protoService{
		Name:    s.Name,
		Comment: mSvrComment,
		RPCs:    mRPCs,
	}

	uProtoServices = append(uProtoServices, mService)
}
