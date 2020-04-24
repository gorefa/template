package model

import (
	"context"
	"time"
	"gogin/pkg/errno"
	"github.com/gorefa/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Cluster struct {
	Enable     bool   `json:"enable" `
	Name       string `json:"name" binding:"required"`
	KubeConfig string `json:"kubeconfig" binding:"required"`
}

const (
	CollectionCluster string = "cluster"
)

func ClusterCreate(cluster Cluster) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"name": cluster.Name}
	if err := Mongo.Collection(CollectionCluster).FindOne(ctx, filter).Decode(&Cluster{}); err != mongo.ErrNoDocuments {
		log.Info("Cluster is exist!")
		return errno.ErrClusterIsExists
	}
	cluster.KubeConfig = cluster.KubeConfig
	//cluster.KubeConfig = base64.StdEncoding.EncodeToString([]byte(cluster.KubeConfig))
	_, err := Mongo.Collection(CollectionCluster).InsertOne(ctx, cluster)
	if err != nil {
		log.Errorf(err, "Create Cluster failed. Name:%s", cluster.Name)
	}
	log.Info("create cluster success")
	return nil
}