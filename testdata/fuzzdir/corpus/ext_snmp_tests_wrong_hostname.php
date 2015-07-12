<?php
require_once(dirname(__FILE__).'/snmp_include.inc');

//EXPECTF format is quickprint OFF
snmp_set_quick_print(false);
snmp_set_valueretrieval(SNMP_VALUE_PLAIN);

var_dump(snmpget('192.168..6.1', 'community', '.1.3.6.1.2.1.1.1.0', $timeout, $retries));

?>
