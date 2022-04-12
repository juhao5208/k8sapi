package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

/**
 * @author  巨昊
 * @date  2021/11/8 15:01
 * @version 1.15.3
 */

//docker api : 1.41

/*func NewClient(server, port, version string) (*client.Client, error) {
	host := "tcp://192.168.99.100" + server + ":" + port
	return client.NewClient(host, version, nil,
		map[string]string{"Content-Type": "application/x-tar"})
}

func PushImage(cli *client.Client, registryUser, registryPassword, image string) error {
	authConfig := types.AuthConfig{
		Username: registryUser,
		Password: registryPassword,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		return err
	}
	log.Infof("Push docker image registry:%v %v", registryUser, registryPassword)

	authStr := base64.URLEncoding.EncodeToString(encodedJSON)
	out, err := cli.ImagePush(context.TODO(), image, types.ImagePushOptions{RegistryAuth: authStr})
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(out)
	if err != nil {
		return err
	}
	log.Infof("Push docker image output:%v", string(body))

	if strings.Contains(string(body), "error") {
		return fmt.Errorf("Push image to docker error")
	}
	return nil
}*/

func main() {
	//helper, err := connhelper.GetConnectionHelper("ssh://minikube@192.168.99.100:2376")
	//if err != nil {
	//	return
	//}
	//httpClient := &http.Client{
	//	Transport: &http.Transport{
	//		DialContext: helper.Dialer,
	//	},
	//}
	//cl, err := client.NewClientWithOpts(
	//	client.WithHTTPClient(httpClient),
	//	client.WithHost(helper.Host),
	//	client.WithDialContext(helper.Dialer),
	//)
	//if err != nil {
	//	fmt.Println("Unable to create docker client")
	//	panic(err)
	//}
	//
	//fmt.Println(cl.ImageList(context.Background(), types.ImageListOptions{}))

	cli, err := client.NewClient("tcp://192.168.99.100:2376", "v1.41", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	for _, image := range images {
		fmt.Println(image.Created)
	}
}
