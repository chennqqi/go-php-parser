<?php

$tidy = new tidy();
$str  = <<<EOF
<p>Isto � um texto em Portugu�s<br>
para testes.</p>
EOF;

$tidy->parseString($str, array('output-xhtml'=>1), 'latin1');
$tidy->cleanRepair();
$tidy->diagnose();
var_dump(tidy_warning_count($tidy) > 0);
var_dump(strlen($tidy->errorBuffer) > 50);

echo $tidy;
?>
