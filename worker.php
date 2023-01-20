<?php

$port = $argv[1];
$channel = $argv[2];

$fp = stream_socket_client("tcp://localhost:" . $port, $errno, $errstr, 30);

if (!$fp) {
    echo "$errstr ($errno)<br />\n";
} else {
    echo "Listen server to port $port with channel $channel, waiting for message" . PHP_EOL;
    fwrite($fp, "CONNECT:$channel " . PHP_EOL);
    while (!feof($fp)) {
        $message = fgets($fp);
        echo "Receive message : " . $message . PHP_EOL;
    }
    fclose($fp);
}