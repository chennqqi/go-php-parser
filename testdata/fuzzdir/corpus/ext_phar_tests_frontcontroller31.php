<?php
try {
Phar::webPhar("test.phar", "/index.php", null, array(), array("fail", "here"));
} catch (Exception $e) {
die($e->getMessage() . "\n");
}
echo "oops did not run\n";
var_dump($_ENV, $_SERVER);
__HALT_COMPILER(); ?>
7                  	   index.php   �hH   JVԋ�      <?php
echo "hi";
�P���f*�|��ݕP-7�V�J   GBMB