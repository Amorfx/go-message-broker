<?php

$port = $argv[1];
$channel = $argv[2];

$fp = fsockopen("tcp://localhost", $port, $errno, $errstr, 30);

if (!$fp) {
    echo "$errstr ($errno)<br />\n";
} else {
    echo "Listen server to port $port with channel $channel, waiting for message" . PHP_EOL;
    fwrite($fp, "CONNECT:$channel " . PHP_EOL);
    while (!feof($fp)) {
        echo fgets($fp);
        //$decodeMessage = json_decode($message);
        //echo "Receive message : " . $decodeMessage . PHP_EOL;
    }
    fclose($fp);
}