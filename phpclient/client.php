<?php

require dirname(__FILE__) . '/vendor/autoload.php';

$hostname = 'localhost:50051';

$client       = new \GrpcPhpClient\Proto\GreeterClient(
    $hostname, [
        'credentials' => Grpc\ChannelCredentials::createInsecure(),
]
);
$indexRequest = new \GrpcPhpClient\Proto\HelloRequest();
$indexRequest->setData("测试");
/**
 * @var \Services\HelloReply $response
 */

[$response, $status] = $client->SayHello($indexRequest)->wait();


var_dump($response->serializeToJsonString());
