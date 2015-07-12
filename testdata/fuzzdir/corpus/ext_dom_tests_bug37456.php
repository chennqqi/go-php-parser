<?php

$doc = new DOMDocument();
$doc->resolveExternals = true;
$doc->load(dirname(__FILE__)."/dom.xml");

$root = $doc->getElementsByTagName('foo')->item(0);
$root->setAttribute('bar', '&gt;');
$attr = $root->setAttribute('bar', 'newval');
print $attr->nodeValue;


?>
