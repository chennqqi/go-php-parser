<?php
$p = new Phar(__FILE__);
var_dump($p["test.txt"]->getMetadata());
$p["test.txt"]->setMetadata("hi2");
var_dump($p["test.txt"]->getMetadata());
echo "ok\n";
__HALT_COMPILER(); ?>
?                     test.txt   ��H   ���E�  	   s:2:"hi";<?php __HALT_COMPILER();�����ο_�gL�?``g�F   GBMB