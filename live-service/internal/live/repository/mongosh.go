package repository

import (
	"context"
	"fmt"
	"live-service/internal/live/pkg/mongosh"

	pb "github.com/Bekzodbekk/paris2024_livestream_protos/genproto/livepb"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type MongoshLiveRepository struct {
	Client mongosh.Mongo
}

func NewMongoshLiveRepository(client mongosh.Mongo) LiveRepository {
	return &MongoshLiveRepository{
		Client: client,
	}
}

func (db *MongoshLiveRepository) CreateLiveStream(req *pb.LiveStream) (*pb.ResponseMessage, error) {
	_, err := db.Client.Collection.InsertOne(context.Background(), req)
	if err != nil {
		return nil, fmt.Errorf("failed to create live stream: %v", err)
	}

	return &pb.ResponseMessage{
		Status:  "success",
		Message: "Live stream created successfully",
	}, nil
}

func (db *MongoshLiveRepository) GetLiveStream(req *pb.GetStreamRequest) (*pb.LiveStream, error) {
	// Find the document by ID
	var result pb.LiveStream
	err := db.Client.Collection.FindOne(context.Background(), bson.M{"event_id": req.Id}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("live stream not found")
		}
		return nil, fmt.Errorf("failed to get live stream: %v", err)
	}

	return &result, nil
}
