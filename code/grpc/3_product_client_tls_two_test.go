package grpc1

//证书生成 步骤 https://blog.csdn.net/qq_35306993/article/details/126907049

import (
	"a/code/grpc/service"
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	"testing"
	"time"
)

//tls 双向认证客户端 + token 认证
func Test66661(t *testing.T) {
	user := &Authentication{
		User:     "admin",
		Password: "admin",
	}

	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, _ := tls.LoadX509KeyPair("tls/client/client.pem", "tls/client/client.key")
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("tls/ca.crt")
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "*.example.com",
		RootCAs:    certPool,
	})

	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(user))

	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}
	defer conn.Close()

	prodClient := service.NewProdServiceClient(conn)

	request := &service.ProductRequest{
		ProdId: 123,
	}
	stockResponse, err := prodClient.GetProductStock(context.Background(), request)
	if err != nil {
		log.Fatal("查询库存出错", err)
	}
	fmt.Println("查询成功", stockResponse.ProdStock)
}

//客户端流
func Test666611(t *testing.T) {
	user := &Authentication{
		User:     "admin",
		Password: "admin",
	}

	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, _ := tls.LoadX509KeyPair("tls/client/client.pem", "tls/client/client.key")
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("tls/ca.crt")
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "*.example.com",
		RootCAs:    certPool,
	})

	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(user))

	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}
	defer conn.Close()

	prodClient := service.NewProdServiceClient(conn)
	//....--------------------------------
	stream, err := prodClient.UpdateStockClientStream(context.Background())
	if err != nil {
		log.Fatal("获取流出错", err)
	}
	rsp := make(chan struct{}, 1)
	go prodRequest(stream, rsp)
	select {
	case <-rsp:
		recv, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatal(err)
		}
		stock := recv.ProdStock
		fmt.Println("客户端收到响应：", stock)
	}
}

//服务端流
func Test66661112(t *testing.T) {
	user := &Authentication{
		User:     "admin",
		Password: "admin",
	}

	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, _ := tls.LoadX509KeyPair("tls/client/client.pem", "tls/client/client.key")
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("tls/ca.crt")
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "*.example.com",
		RootCAs:    certPool,
	})

	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(user))

	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}
	defer conn.Close()

	prodClient := service.NewProdServiceClient(conn)
	//----------------------------------------
	request := &service.ProductRequest{
		ProdId: 123,
	}
	stream, err := prodClient.GetProductStockServerStream(context.Background(), request)
	if err != nil {
		log.Fatal("获取流出错", err)
	}

	for {
		recv, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端数据接收完成")
				err := stream.CloseSend()
				if err != nil {
					log.Fatal(err)
				}
				break
			}
			log.Fatal(err)
		}
		fmt.Println("客户端收到的流", recv.ProdStock)
	}

}

//双向流
func Test666611112132(t *testing.T) {
	user := &Authentication{
		User:     "admin",
		Password: "admin",
	}

	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, _ := tls.LoadX509KeyPair("tls/client/client.pem", "tls/client/client.key")
	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("tls/ca.crt")
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)
	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		// 要求必须校验客户端的证书。可以根据实际情况选用以下参数
		ServerName: "*.example.com",
		RootCAs:    certPool,
	})

	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(user))

	if err != nil {
		log.Fatal("服务端出错，连接不上", err)
	}
	defer conn.Close()

	prodClient := service.NewProdServiceClient(conn)
	//-------------------------------
	stream, err := prodClient.SayHelloStream(context.Background())

	for {
		request := &service.ProductRequest{
			ProdId: 123,
		}
		err = stream.Send(request)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
		recv, err := stream.Recv()
		if err != nil {
			log.Fatal(err)
		}
		//websocket
		fmt.Println("客户端收到的流信息", recv.ProdStock)
	}
}
func prodRequest(stream service.ProdService_UpdateStockClientStreamClient, rsp chan struct{}) {
	count := 0
	for {
		request := &service.ProductRequest{
			ProdId: 123,
		}
		err := stream.Send(request)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
		count++
		if count > 10 {
			rsp <- struct{}{}
			break
		}
	}
}

//客户端需要实现 `PerRPCCredentials` 接口。
type Authentication struct {
	User     string
	Password string
}

//`GetRequestMetadata` 方法返回认证需要的必要信息
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

//`RequireTransportSecurity` 方法表示是否启用安全链接，在生产环境中，一般都是启用的，但为了测试方便，暂时这里不启用了。
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
