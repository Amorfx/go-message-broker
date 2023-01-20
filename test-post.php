<?php

$port = $argv[1];
$channel = $argv[2];

$fp = fsockopen("tcp://localhost", $port, $errno, $errstr, 30);

if (!$fp) {
    echo "$errstr ($errno)<br />\n";
} else {
    $value = 'mymessage is for you';
    $message = ['chan' => $channel, 'value' => $value];
    $message = json_encode($message);
    echo "Push message $message in port $port and in channel $channel";
    fwrite($fp, "SEND:$message " . PHP_EOL);

    fclose($fp);
}