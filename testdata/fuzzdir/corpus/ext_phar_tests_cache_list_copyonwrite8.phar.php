<?php
$p = new Phar(__FILE__);
var_dump($p->getAlias());
$p2 = new Phar(__FILE__);
$p->setAlias("hi");
echo $p2->getAlias(),"\n";
echo "ok\n";
__HALT_COMPILER(); ?>
6                     test.txt   t��H   zzo��      hi
�����Ji5���4QCڱ�   GBMB