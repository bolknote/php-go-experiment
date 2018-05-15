#!/usr/bin/php -d extension=./sse2_strlen.so
<?php
$str = str_repeat("J", 1000);

function benchmark($str, $func)
{
    $start = microtime(true);
    for ($i = 0; $i<1e6; $i++) {
        $func($str);
    }

    printf("%13s:\t%s\n", $func, microtime(true) - $start);
}

mb_internal_encoding('utf-8');
ini_set('iconv.internal_encoding', 'utf-8');

benchmark($str, 'sse2_strlen');

benchmark($str, 'mb_strlen');

benchmark($str, 'iconv_strlen');

benchmark($str, 'strlen');
