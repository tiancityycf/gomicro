package main

import (
	"fmt"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	consumerLibrary "github.com/aliyun/aliyun-log-go-sdk/consumer"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/go-kit/kit/log/level"
	"os"
	"os/signal"
	"sync"
	"time"
)

const Endpoint = "cn-beijing.log.aliyuncs.com"
const AccessKeyID = "LTAI4FsrrVzp9zck7ee89Dzm"
const AccessKeySecret = "7obh1j6hgRRDZbeVblYzy5MpPJBKlG"
const Project = "maka-test"
const LogStore = "test"

func main() {
	//recvLog()
	sendLog()
}

func recvLog() {
	option := consumerLibrary.LogHubConfig{
		Endpoint:          Endpoint,
		AccessKeyID:       AccessKeyID,
		AccessKeySecret:   AccessKeySecret,
		Project:           Project,
		Logstore:          LogStore,
		ConsumerGroupName: "consumerG1",
		ConsumerName:      "consumer1",
		IsJsonType:        true,
		//是否按序消费，默认为false。
		InOrder: true,
		//日志文件输出路径，不设置的话默认输出到stdout。
		LogFileName: "./consumer.log",
		//单个日志存储数量，默认为10M。
		LogMaxSize: 102400,
		//default 200(Millisecond), don't configure it too small (<100Millisecond)
		DataFetchIntervalInMs: 100,
		//从服务端一次拉取日志组数量，日志组可参考内容日志组，默认值是1000，其取值范围是1-1000。
		MaxFetchLogGroupCount: 1000,
		// This options is used for initialization, will be ignored once consumer group is created and each shard has been started to be consumed.
		// Could be "begin", "end", "specific time format in time stamp", it's log receiving time.
		CursorPosition: consumerLibrary.BEGIN_CURSOR,
	}

	consumerWorker := consumerLibrary.InitConsumerWorker(option, process)
	ch := make(chan os.Signal)
	signal.Notify(ch)
	consumerWorker.Start()
	if _, ok := <-ch; ok {
		level.Info(consumerWorker.Logger).Log("msg", "get stop signal, start to stop consumer worker", "consumer worker name", option.ConsumerName)
		consumerWorker.StopAndWait()
	}
}

// Fill in your consumption logic here, and be careful not to change the parameters of the function and the return value,
// otherwise you will report errors.
func process(shardId int, logGroupList *sls.LogGroupList) string {
	//消费失败处理办法
	for {
		err := func(shardId int, logGroupList *sls.LogGroupList) error {
			fmt.Println(shardId, logGroupList)
			for _, v := range logGroupList.LogGroups {
				fmt.Println(shardId, v.GetLogs())
			}
			// 在这个函数当中去写具体的消费逻辑，如果消费失败，返回自定义的error
			return nil
		}(shardId, logGroupList)
		if err != nil {
			// 当捕获到消费失败的error以后只要重复继续执行消费逻辑即可，不要跳出process方法。
			continue
		} else {
			// 消费成功的话，跳出循环，process方法执行完毕，会继续拉取数据进行下次消费。
			break
		}
	}
	return ""
}
func sendLog() {
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = Endpoint
	producerConfig.AccessKeyID = AccessKeyID
	producerConfig.AccessKeySecret = AccessKeySecret

	topic := "topictest"

	producerInstance := producer.InitProducer(producerConfig)
	ch := make(chan os.Signal)
	signal.Notify(ch)
	producerInstance.Start()
	var m sync.WaitGroup
	for i := 0; i < 2; i++ {
		m.Add(1)
		go func() {
			defer m.Done()
			for i := 0; i < 2; i++ {
				// GenerateLog  is producer's function for generating SLS format logs
				// GenerateLog has low performance, and native Log interface is the best choice for high performance.
				log := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"content": "test111", "content222": fmt.Sprintf("%v", i)})
				err := producerInstance.SendLog(Project, LogStore, topic, "127.0.0.2", log)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
	m.Wait()
	fmt.Println("Send completion")
	if _, ok := <-ch; ok {
		fmt.Println("Get the shutdown signal and start to shut down")
		producerInstance.Close(60)
	}
}
