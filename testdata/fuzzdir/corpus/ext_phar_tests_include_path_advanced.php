<?php
set_include_path('.' . PATH_SEPARATOR . 'phar://' . dirname(__FILE__) . '/files/include_path2.phar' );
include 'phar://' . __FILE__ . '/hello/test.php';
set_include_path('.' . PATH_SEPARATOR . 'phar://' . dirname(__FILE__) . '/files/include_path2.phar/test');
include 'phar://' . __FILE__ . '/hello/test.php';
echo "ok\n";
__HALT_COMPILER(); ?>
<                     hello/test.php   r��H   ��U�      <?php
include "file1.php";���Ob�YN�1c��T�ʸ/   GBMB