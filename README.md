# ChatBox
RPC based application where the client can directly send msg to server and vice-versa.



We are here implementing a simple application based on Golang in which Server and the Client can connect with each other and can exchange information with each other, this exchange of information is a two way.

1. First of all, we need to create a Golang server: - 
2. To create the Golang server we will use Server.go file. 
3. Here we define main() function in which we print some information and in the next line we use net.listen() which will start our go server. 
4. We also defined the write and read operation so that it can read the data from the client. 
5. We also defined the Go routine which will listen to the client response and run concurrently. 
6. The full code for Server.go file which will run our Go server is given below: -

![image](https://user-images.githubusercontent.com/72275403/110161966-144d7b80-7e14-11eb-83a8-0696b4a9c7e2.png)

![image](https://user-images.githubusercontent.com/72275403/110162008-229b9780-7e14-11eb-924e-c9b3a336ed40.png)


We need to create a Golang Client: -
1.To create a Golang Client we use Client.go file.
2. Here we define the main() function in which we define the net.conn() which will setup the connection with our Go Server for the sake of the communication.
3. We also defined the read and write operations here so that it can read the data from the server side.
4. The full code of Client.go file which will interact with our Go server for further communication ids given below: -

![image](https://user-images.githubusercontent.com/72275403/110162111-4959ce00-7e14-11eb-9701-672d644be342.png)

![image](https://user-images.githubusercontent.com/72275403/110162131-524a9f80-7e14-11eb-9f0a-e1c5812f4052.png)


EXECUTION: -

![image](https://user-images.githubusercontent.com/72275403/110162232-773f1280-7e14-11eb-8d34-6a0201514ca8.png)


![image](https://user-images.githubusercontent.com/72275403/110162263-81611100-7e14-11eb-892b-5b9e3e3f541a.png)


![image](https://user-images.githubusercontent.com/72275403/110162278-87ef8880-7e14-11eb-866f-d06fbe68467c.png)


![image](https://user-images.githubusercontent.com/72275403/110162327-976ed180-7e14-11eb-9e43-2816b93fc6ba.png)


##################################################### THANK YOU #####################################################################################

