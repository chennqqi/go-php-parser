<?php
$p = new Phar(__FILE__);
var_dump(isset($p["copied"]));
$p->copy("test.txt","copied");
var_dump(isset($p["copied"]));
echo "ok\n";
__HALT_COMPILER(); ?>
6                     test.txt   K��H   ���E�      <?php __HALT_COMPILER();*��I_B�.�֩�.F�"�_z   GBMB