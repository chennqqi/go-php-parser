<?php
mkdir("chroot_001_x");
var_dump(is_dir("chroot_001_x"));
var_dump(chroot("chroot_001_x"));
var_dump(is_dir("chroot_001_x"));
var_dump(realpath("."));
?>
