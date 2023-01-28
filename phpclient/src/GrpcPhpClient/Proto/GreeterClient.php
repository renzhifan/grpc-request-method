<?php
// GENERATED CODE -- DO NOT EDIT!

namespace GrpcPhpClient\Proto;

/**
 * 定义服务
 */
class GreeterClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * SayHello 一元RPC方法
     * @param \GrpcPhpClient\Proto\HelloRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function SayHello(\GrpcPhpClient\Proto\HelloRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/hellostream.Greeter/SayHello',
        $argument,
        ['\GrpcPhpClient\Proto\HelloResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * GetStream 服务端流
     * @param \GrpcPhpClient\Proto\HelloRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function GetStream(\GrpcPhpClient\Proto\HelloRequest $argument,
      $metadata = [], $options = []) {
        return $this->_serverStreamRequest('/hellostream.Greeter/GetStream',
        $argument,
        ['\GrpcPhpClient\Proto\HelloResponse', 'decode'],
        $metadata, $options);
    }

    /**
     * SetStream 客户端流
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function SetStream($metadata = [], $options = []) {
        return $this->_clientStreamRequest('/hellostream.Greeter/SetStream',
        ['\GrpcPhpClient\Proto\HelloResponse','decode'],
        $metadata, $options);
    }

    /**
     * AllStream 双向流
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function AllStream($metadata = [], $options = []) {
        return $this->_bidiRequest('/hellostream.Greeter/AllStream',
        ['\GrpcPhpClient\Proto\HelloResponse','decode'],
        $metadata, $options);
    }

}
