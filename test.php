<?php

$fp = fsockopen("tcp://localhost", 4000, $errno, $errstr, 30);

$channel = $argv[1];
if (!$fp) {
    echo "$errstr ($errno)<br />\n";
} else {
    sleep(2);
    echo "Listen server to port 4000 with channel $channel, waiting for message";
    fwrite($fp, "POST " . PHP_EOL);
    /*while (!feof($fp)) {
        echo fgets($fp, 128);
    }*/
    fclose($fp);
}