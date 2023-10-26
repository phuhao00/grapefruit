# s3 访问控制

- 为了进一步加强您的 S3 桶安全性，如果您只是希望您 S3 桶内的对象被公开访问，而不需要允许公开上传，列出桶内对象的话，您可以参考以下桶策略以减少公开的权限：
```json
{
"Version": "2012-10-17",
"Statement": [
{
"Sid": "Statement1",
"Effect": "Allow",
"Principal": "*",
"Action": "s3:GetObject",
"Resource": "arn:aws-cn:s3:::tutu-backend/*"
}
]
}

```

```html
https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#S3.PutObject 
```

# [Fargate On-Demand vCPU resource count]

```html
亚马逊云科技 Trusted Advisor 来查看一些特定服务（DynamoDB、EC2、VPC）的当前限额和使用量，有关其他服务的当前限额：
https://console.amazonaws.cn/trustedadvisor/home#/category/service-limits

```

# 费用

```html
1. 对于您在账户下使用的亚马逊云科技服务，计费从资源启动运行时开始，到资源终止或停止时结束。当您不再使用亚马逊云科技服务时，请您及时关闭或删除资源，避免产生未预期费用。
2. 我们建议您定期监测您账户的使用情况，避免不必要的费用。您也可以使用CloudWatch监测服务费用，设置支出警报，更加便捷有效的管理您的费用支出。 以下链接是设置账单警报的详细步骤：
https://docs.amazonaws.cn/AmazonCloudWatch/latest/monitoring/monitor_estimated_charges_with_cloudwatch.html
3. 您账户的资源池容量增加时将增加发生恶意活动（非法访问、拒绝服务攻击、数据丢失）的风险。
4. 我们建议您使用免费的亚马逊云科技 Trusted Advisor对您账户的安全性进行检查，通过关闭缺口、启用各种安全功能以及检查权限，以提高应用程序的安全性。请参考如下文档了解并开始使用Trusted Advisor：
https://www.amazonaws.cn/support/trustedadvisor/


```

# EKS 创建

```html
1. 创建1.28版本的EKS集群backend-bff-dev。其中，Cluster role为新建的IAM角色eks-dev。创建步骤可参考文档[1]。Cluster endpoint accessInfo选择了Public。关于终端节点公有和私有访问的区别可参考文档[2]。

2. 创建EKS集群节点组backend-bff-group。其中Node Role为新建的IAM角色eks-node-group-dev。创建步骤可参考文档[3]。

3. 在客户端配置kubectl。安装步骤参考文档[4]。

4. 执行完aws eks update-kubeconfig --region cn-northwest-1 --name backend-bff-dev命令后，使用kubectl get nodes能看到集群的两个节点。


参考文档：
[1] https://docs.amazonaws.cn/eks/latest/userguide/service_IAM_role.html#create-service-role
[2] https://docs.amazonaws.cn/eks/latest/userguide/cluster-endpoint.html#modify-endpoint-access
[3] https://docs.amazonaws.cn/eks/latest/userguide/create-node-role.html#create-worker-node-role
[4] https://docs.amazonaws.cn/eks/latest/userguide/install-kubectl.html

```

# NLB

```html
“我们将 TCP 流的空闲超时值设置为 350 秒。您无法修改此值。客户端或目标可以使用 TCP keepalive 数据包重置空闲超时值。为维护 TLS 连接而发送的 Keepalive 数据包不能包含数据或负载。”
而对于ALB，您同样可以参考文档[2]中的配置步骤。

因此，如果您的应用存在较长的空闲连接时间，建议您可以考虑在客户端一侧通过一定的频率发送 keepalive 信号来重置NLB的350秒空闲超时。

而对于 ECS Service Connect，实际上是需要15秒内获取到HTTP响应，目前还未公开相关参数能够调整这个限制，因此发送Keep-Alive来重置空闲等待时长的机制不一定适用于ECS Service Connect。

【参考文档】
[1] https://docs.amazonaws.cn/elasticloadbalancing/latest/network/network-load-balancers.html#connection-idle-timeout
[2] https://docs.amazonaws.cn/elasticloadbalancing/latest/application/application-load-balancers.html#connection-idle-timeout



```

# 限流

```html
1. 您可以使用usage plan去限制某些固定API Key的请求速率和数量【1】。 
2. 您可以在阶段处，限制您的某个阶段的请求速率【2】。
3. 您可以在cloud watch log中，搜索您的API GatewayID找到您的API执行日志。
4. 您可以参考如下文档创建alarm【3】。
5. 您可以参考如下文档了解API Gateway集成WAF【4】。

【1】https://docs.amazonaws.cn/apigateway/latest/developerguide/api-gateway-create-usage-plans-with-console.html 
【2】https://docs.amazonaws.cn/apigateway/latest/developerguide/api-gateway-request-throttling.html 
【3】https://docs.amazonaws.cn/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html 
【4】https://docs.amazonaws.cn/apigateway/latest/developerguide/apigateway-control-access-aws-waf.html 

```

# lambda 不支持 websocket

```html
由于调用Lambda只能使用HTTP的方式，也无法利用您建立好的Websocket连接

```

# NLB

```html
经过我自行测试,NLB是支持websocket协议的。
您反馈您ECS上部署websocket之后，通过微信小程序访问NLB去连接websocket的应用无响应，通过电脑浏览器访问NLB去连接websocket可以正确响应。
NLB是四层的网络负载均衡器，本身不会解封装七层的应用层报文。
建议您可以在ECS上进行抓包，tcpdump的命令是sudo tcpdump port 8888 -w test.pcap，然后您可以看下ECS收到的报文是否相同。
或者您也可以通过流量镜像，将NLB的三个ENI的流量都镜像到一个EC2上，然后再通过tcpdump进行抓包。
关于流量镜像的例子和配置方法，您可以查看以下文档：
https://docs.amazonaws.cn/en_us/vpc/latest/mirroring/tm-example-inbound-tcp.html 

关于api gateway，经过和您的确认，websocket流量的整条链路上是没有api gateway的。
以及关于ECS的服务，创建服务后，无法从 Amazon Web Services Management Console 更改负载均衡器配置。
相关文档：
https://docs.amazonaws.cn/AmazonECS/latest/userguide/create-service-review.html 


```

# mysql 最大连接数设置

```html
1）创建自定义参数组。
2）修改了max_connections参数至500
3）挂载新的参数组到RDS实例上

后续请您重启一次RDS实例，使参数组正常应用到RDS实例上。
您可以在modify页面修改机型、增加存储空间，这些操作不会影响已经存储的数据。

如下官方文档请您参考：
https://docs.aws.amazon.com/zh_cn/AmazonRDS/latest/UserGuide/USER_WorkingWithParamGroups.html 

```

# ECS Service固定公网IP地址。

```html
您可以参考使用NLB的方式暴露您的ECS Service服务，对此，您可以参考NLB的文档描述[1]：
"支持将静态 IP 地址用于负载均衡器。还可以针对为负载均衡器启用的每个子网分配一个弹性 IP 地址。"

因此，如电话中沟通，结合您的实际需求，您可以参考在ECS Service中使用NLB的方式暴露服务，对于关联API Gateway域名的问题，您可以修改API Gateway自定义域名的配置并在您的测试环境中进行验证。

关于GWLB的使用场景，大多数是使用在将流量转发到防火墙服务，做流量检测等，您可以参考以下的GWLB介绍文档[2]。

【参考文档】
[1] https://aws.amazon.com/cn/blogs/aws/introducing-aws-gateway-load-balancer-easy-deployment-scalability-and-high-availability-for-partner-appliances/ 
[2] https://docs.amazonaws.cn/elasticloadbalancing/latest/network/introduction.html#network-load-balancer-benefits 

```


 
