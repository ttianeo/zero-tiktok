// Code generated by goctl. DO NOT EDIT.
// Source: interaction.proto

package interactionservice

import (
	"context"

	"zero-tiktok/service/interaction/pb/zero-tiktok/service/interaction"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Comment              = interaction.Comment
	CommentListRequest   = interaction.CommentListRequest
	CommentListResponse  = interaction.CommentListResponse
	CommentRequest       = interaction.CommentRequest
	CommentResponse      = interaction.CommentResponse
	FollowListRequest    = interaction.FollowListRequest
	FollowListResponse   = interaction.FollowListResponse
	FollowerListRequest  = interaction.FollowerListRequest
	FollowerListResponse = interaction.FollowerListResponse
	FriendListRequest    = interaction.FriendListRequest
	FriendListResponse   = interaction.FriendListResponse
	HasFollowedRequest   = interaction.HasFollowedRequest
	HasFollowedResponse  = interaction.HasFollowedResponse
	RelationRequest      = interaction.RelationRequest
	RelationResponse     = interaction.RelationResponse

	InteractionService interface {
		Comment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
		CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error)
		Relation(ctx context.Context, in *RelationRequest, opts ...grpc.CallOption) (*RelationResponse, error)
		FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error)
		FollowerList(ctx context.Context, in *FollowerListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error)
		FriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error)
		HasFollowed(ctx context.Context, in *HasFollowedRequest, opts ...grpc.CallOption) (*HasFollowedResponse, error)
	}

	defaultInteractionService struct {
		cli zrpc.Client
	}
)

func NewInteractionService(cli zrpc.Client) InteractionService {
	return &defaultInteractionService{
		cli: cli,
	}
}

func (m *defaultInteractionService) Comment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	client := interaction.NewInteractionServiceClient(m.cli.Conn())
	return client.Comment(ctx, in, opts...)
}

func (m *defaultInteractionService) CommentList(ctx context.Context, in *CommentListRequest, opts ...grpc.CallOption) (*CommentListResponse, error) {
	client := interaction.NewInteractionServiceClient(m.cli.Conn())
	return client.CommentList(ctx, in, opts...)
}

func (m *defaultInteractionService) Relation(ctx context.Context, in *RelationRequest, opts ...grpc.CallOption) (*RelationResponse, error) {
	client := interaction.NewInteractionServiceClient(m.cli.Conn())
	return client.Relation(ctx, in, opts...)
}

func (m *defaultInteractionService) FollowList(ctx context.Context, in *FollowListRequest, opts ...grpc.CallOption) (*FollowListResponse, error) {
	client := interaction.NewInteractionServiceClient(m.cli.Conn())
	return client.FollowList(ctx, in, opts...)
}

func (m *defaultInteractionService) FollowerList(ctx context.Context, in *FollowerListRequest, opts ...grpc.CallOption) (*FollowerListResponse, error) {
	client := interaction.NewInteractionServiceClient(m.cli.Conn())
	return client.FollowerList(ctx, in, opts...)
}

func (m *defaultInteractionService) FriendList(ctx context.Context, in *FriendListRequest, opts ...grpc.CallOption) (*FriendListResponse, error) {
	client := interaction.NewInteractionServiceClient(m.cli.Conn())
	return client.FriendList(ctx, in, opts...)
}

func (m *defaultInteractionService) HasFollowed(ctx context.Context, in *HasFollowedRequest, opts ...grpc.CallOption) (*HasFollowedResponse, error) {
	client := interaction.NewInteractionServiceClient(m.cli.Conn())
	return client.HasFollowed(ctx, in, opts...)
}