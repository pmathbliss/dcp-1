<?php

//get list of numbers from command line and the sum
$numberOfArgs = count($argv);
if($numberOfArgs < 4)
{
    die("Not enough arguments to process. There must be at least 3 arguments.");
}

$args = $argv;
array_shift($args);
array_shift($args);

$k = (int)$argv[1];


$searchInts = [];

$currentInt = 0;
$findValue = 0;

$found = false;
for($i = 2; $i < $numberOfArgs; $i ++)
{
    $currentInt = (int)$argv[$i];
    if($currentInt >= $k)
    {
        //skip
        continue;
    }
    $findValue = $k - $currentInt;
    $searchInts[] = $findValue;
    if(in_array($currentInt,$searchInts) || in_array($findValue, $args))
    {
        $found = true;
        break;
    }
}

echo $found?"Found two integers {$currentInt} and {$findValue} add to {$k}":"No two integers add to {$k}";