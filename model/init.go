package model

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/gorefa/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	err error
	DB        *gorm.DB
	Mongo        *mongo.Database
	Clientset *kubernetes.Clientset

)

func Init() {
	DB = InitDB()
	Mongo, err = InitMongo()
	if err != nil {
		panic(err)
	}

	//Clientset = InitK8S()
}

func InitDB() *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
		true,
		"Local")

	db, err := gorm.Open("mysql", config)
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", viper.GetString("db.username"))
	}

	db.DB().SetMaxOpenConns(2000)
	db.DB().SetMaxIdleConns(0)

	return db
}

func InitMongo() (*mongo.Database, error) {
	host := viper.GetStringSlice("mongodb.host")
	username := viper.GetString("mongodb.username")
	password := viper.GetString("mongodb.password")
	database := viper.GetString("mongodb.database")
	hosts := strings.Join(host, ",")
	uri := fmt.Sprintf(`mongodb://%s:%s@%s`,
		username,
		password,
		hosts,
	)
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	db := client.Database(database)

	return db, nil

}

func InitK8S() *kubernetes.Clientset {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	Clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return Clientset
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}


func NewInitK8S(clustername string) (*kubernetes.Clientset ,error){
	filter := bson.D{{"name", clustername}}
	cluster := Cluster{}

	if err := Mongo.Collection(CollectionCluster).FindOne(context.TODO(), filter).Decode(&cluster); err != nil {
		log.Errorf(err, "get cluster config from DB failed.")
		return nil,err
	}
	config, err := base64.StdEncoding.DecodeString(cluster.KubeConfig)
	if err != nil {
		log.Errorf(err, "base64 decode failed.cluster: %s ", clustername)
		return nil, err
	}

	cliConfig, err := clientcmd.NewClientConfigFromBytes(config)
	if err != nil {
		log.Errorf(err, "kube config error.")
		return nil, err
	}

	conf, err := cliConfig.ClientConfig()
	if err != nil {
		log.Errorf(err, " kube clinet config error.")
		return nil, err
	}

	Clientset, err = kubernetes.NewForConfig(conf)
	if err != nil {
		log.Errorf(err, "Create k8s clientset failed. ")
		return nil, err
	}

	return Clientset,nil
}
