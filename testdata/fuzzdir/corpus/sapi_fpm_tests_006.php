<?php

include "include.inc";

$logfile = dirname(__FILE__).'/php-fpm.log.tmp';
$port = 9000+PHP_INT_SIZE;

$cfg = <<<EOT
[global]
error_log = $logfile
[unconfined]
listen = [::]:$port
listen.allowed_clients = ::1
pm = dynamic
pm.max_children = 5
pm.start_servers = 2
pm.min_spare_servers = 1
pm.max_spare_servers = 3
EOT;

$fpm = run_fpm($cfg, $tail);
if (is_resource($fpm)) {
    fpm_display_log($tail, 2);
    try {
		run_request('127.0.0.1', $port);
		echo "IPv4 ok\n";
	} catch (Exception $e) {
		echo "IPv4 error\n";
	}
    try {
		run_request('[::1]', $port);
		echo "IPv6 ok\n";
	} catch (Exception $e) {
		echo "IPv6 error\n";
	}
    proc_terminate($fpm);
    stream_get_contents($tail);
    fclose($tail);
    proc_close($fpm);
}

?>
