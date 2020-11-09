package dao

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//	"io/ioutil"
	"log"
	//	"net/http"
	"test/model"
	"time"
)

var client = connect("mongodb+srv://accountYan:Covid19webarrivingsoon!@test.rdjlc.mongodb.net/test?retryWrites=true&w=majority")

func connect(addr string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(addr)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to MongoDB :)")
	//collection := client.Database("AiProj").Collection("AiPaper")
	return client
}

func SearchByKeyword(option string, keyword string) []bson.M {
	collection := client.Database("AiProj").Collection("AiPaper")
	start := time.Now()
	result, err := collection.Find(context.TODO(), bson.M{option: bson.M{"$regex": keyword, "$options": "$i"}})
	times := time.Since(start)
	fmt.Println(times)
	var papers []bson.M
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(result)
	if err = result.All(context.TODO(), &papers); err != nil {
		log.Fatal(err)
	}
	_ = result.Close(context.TODO())
	//fmt.Println(papers)
	return papers
}

func PaperNumberGraph(keyword string) []bson.M {
	//mongodb://accountUser:Covid19webarrivingsoon!@localhost:27017/test?authSource=test&readPreference=primary&appname=MongoDB%20Compass&ssl=false

	collection := client.Database("AiProj").Collection("AiPaper")
	pipeline := String2Interface(CustomLineGraph(keyword))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pipe, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		panic(err)
	}
	var showWithInfo []bson.M
	if err = pipe.All(ctx, &showWithInfo); err != nil {
		panic(err)
	}
	fmt.Println(showWithInfo)
	return showWithInfo
}

func JournalNumberGraph(keyword string) []bson.M {

	collection := client.Database("AiProj").Collection("AiPaper")
	pipeline := String2Interface(CustomPipeline(keyword))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pipe, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		panic(err)
	}
	var showWithInfo []bson.M
	if err = pipe.All(ctx, &showWithInfo); err != nil {
		panic(err)
	}
	//fmt.Println(showWithInfo)
	return showWithInfo
}

func GeoMap() []bson.M {
	collection := client.Database("paperDetail").Collection("pdf")
	pipeline := String2Interface(model.Geo)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pipe, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		panic(err)
	}
	var showWithInfo []bson.M
	if err = pipe.All(ctx, &showWithInfo); err != nil {
		panic(err)
	}
	fmt.Println(showWithInfo)
	return showWithInfo
}

func RadarGraph(keyword string) []bson.M {
	collection := client.Database("AiProj").Collection("AiPaper")
	pipeline := String2Interface(CustomRadar(keyword))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pipe, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		panic(err)
	}
	var showWithInfo []bson.M
	if err = pipe.All(ctx, &showWithInfo); err != nil {
		panic(err)
	}
	fmt.Println(showWithInfo)
	return showWithInfo
}

func RelatedTopicPie(keyword string) []bson.M {
	collection := client.Database("AiProj").Collection("AiPaper")
	pipeline := String2Interface(CustomTopK(keyword))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pipe, err := collection.Aggregate(ctx, pipeline)
	fmt.Println("pipeline..")
	if err != nil {
		fmt.Println("crash...")
		panic(err)
	}
	var showWithInfo []bson.M
	if err = pipe.All(ctx, &showWithInfo); err != nil {
		panic(err)
	}
	//fmt.Println(showWithInfo)
	return showWithInfo
}
func geoMap() []bson.M {
	collection := client.Database("AiProj").Collection("AiPaper")

	pipeline := String2Interface(model.Pie)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pipe, err := collection.Aggregate(ctx, pipeline)
	fmt.Println("pipeline..")
	if err != nil {
		fmt.Println("crash...")
		panic(err)
	}
	var showWithInfo []bson.M
	if err = pipe.All(ctx, &showWithInfo); err != nil {
		panic(err)
	}
	//fmt.Println(showWithInfo)
	return showWithInfo
}

func WordCloud() []bson.M {
	collection := client.Database("AiProj").Collection("AiPaper")

	pipeline := String2Interface(model.WordCloud)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pipe, err := collection.Aggregate(ctx, pipeline)
	fmt.Println("pipeline..")
	if err != nil {
		fmt.Println("crash...")
		panic(err)
	}
	var showWithInfo []bson.M
	if err = pipe.All(ctx, &showWithInfo); err != nil {
		panic(err)
	}
	//fmt.Println(showWithInfo)
	return showWithInfo
}

func AuthorPaperGraph(keyword string) []bson.M {
	collection := client.Database("AiProj").Collection("AiPaper")
	pipeline := String2Interface(CustomAuthorPaper(keyword))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pipe, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		panic(err)
	}
	var showWithInfo []bson.M
	if err = pipe.All(ctx, &showWithInfo); err != nil {
		panic(err)
	}
	fmt.Println(showWithInfo)
	return showWithInfo
}

func String2Interface(pipe string) interface{} {
	pipelineJson := pipe
	var pipeline interface{}
	err := bson.UnmarshalExtJSON([]byte(pipelineJson), true, &pipeline)
	if err != nil {
		fmt.Println("error", err)
	}
	return pipeline
}

func main() {
	//searchByKeyword("title","pneumonia")
	//	fmt.Println(SearchByKeyword("title","2020"))
	////	PaperNumberGraph()
	//	fmt.Println(RadarGraph())
	//fmt.Println(RelatedTopicPie())
	//fmt.Println(SolrSearchByKeyword("2020"))
	//SolrSearchByKeyword()
}
