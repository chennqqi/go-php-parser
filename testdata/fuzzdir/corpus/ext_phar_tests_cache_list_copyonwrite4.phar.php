<?php
var_dump(file_exists("phar://" . __FILE__ . "/test.txt"));
clearstatcache();
Phar::mount("test.txt", "phar://" . __FILE__ . "/tobemounted");
var_dump(file_exists("phar://" . __FILE__ . "/test.txt"), file_get_contents("phar://" . __FILE__ . "/test.txt"));
echo "ok\n";
__HALT_COMPILER(); ?>
9                     tobemounted   ���H   �*�ض      hi�����X���PF�.��3   GBMB