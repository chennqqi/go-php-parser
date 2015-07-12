<?php
ini_set('html_errors', false);
var_dump(filter_input(INPUT_GET, "a", FILTER_SANITIZE_STRIPPED));
var_dump(filter_input(INPUT_GET, "b", FILTER_SANITIZE_URL));
var_dump(filter_input(INPUT_GET, "a", FILTER_SANITIZE_SPECIAL_CHARS, array(1,2,3,4,5)));
var_dump(filter_input(INPUT_GET, "b", FILTER_VALIDATE_FLOAT, new stdClass));
var_dump(filter_input(INPUT_POST, "c", FILTER_SANITIZE_STRIPPED, array(5,6,7,8)));
var_dump(filter_input(INPUT_POST, "d", FILTER_VALIDATE_FLOAT));
var_dump(filter_input(INPUT_POST, "c", FILTER_SANITIZE_SPECIAL_CHARS));
var_dump(filter_input(INPUT_POST, "d", FILTER_VALIDATE_INT));

var_dump(filter_var(new stdClass, "d"));

var_dump(filter_input(INPUT_POST, "c", "", ""));
var_dump(filter_var("", "", "", "", ""));
var_dump(filter_var(0, 0, 0, 0, 0));

echo "Done\n";
?>
