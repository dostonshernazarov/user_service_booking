package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"user_service_booking/config"
	pb "user_service_booking/genproto/user_proto"
	"user_service_booking/pkg/db"
	"user_service_booking/pkg/logger"
	"user_service_booking/queue/kafka/consumer"
	"user_service_booking/service"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "user-booking-service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	userService := service.NewUserService(connDB, log)

	consumer, err := consumer.NewKafkaConsumerInit([]string{"localhost:9092"}, "test-topic", "1")
	if err != nil {
		log.Fatal("NewKafkaConsumerInit: %v", logger.Error(err))
	}

	defer consumer.Close()

	go func() {
		consumer.ConsumeMessages(consumerHandler)
	}()

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}

func consumerHandler(message []byte) {
	fmt.Println(string(message))
}
