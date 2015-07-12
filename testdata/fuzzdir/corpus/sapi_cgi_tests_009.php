<?php

include "include.inc";

$php = get_cgi_path();
reset_env_vars();

$f = tempnam(sys_get_temp_dir(), 'cgitest');

putenv("TRANSLATED_PATH=".$f."/x");
putenv("SCRIPT_FILENAME=".$f."/x");
file_put_contents($f, '<?php var_dump($_SERVER["TRANSLATED_PATH"]); ?>');

echo (`$php -n $f`);

echo "Done\n";

@unlink($f);
?>
