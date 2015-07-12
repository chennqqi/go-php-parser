<?php
include "connect.inc";

$link = ldap_connect_and_bind($host, $port, $user, $passwd, $protocol_version);
insert_dummy_data($link, $base);

$dn = "$base";
$filter = "(objectclass=person)";
var_dump(
	$result = ldap_search($link, $dn, $filter, array('sn'), 1, 1, 3),
	ldap_get_entries($link, $result)
);
?>
===DONE===
