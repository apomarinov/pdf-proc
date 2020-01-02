<?php

$sourceDir = "words_5_pairs/";
$ocrOutput = "ocr.txt";

$files = scandir($folder);
array_shift($files);
array_shift($files);
array_shift($files);

for ($i=1; $i < count($files); $i++) {

}


$lineData = json_decode(file_get_contents($pointsFile), true);

$wordCounter = 1;
foreach ($lineData as $name => $points) {
	$prev = 0;
	$points[] = 9999;
	for ($i=0; $i < count($points); $i++) {
	    $w2 = $width / 2;
		exec("convert $sourceDir/$name.jpg -crop $w2x".($points[$i] - $prev)."+0+$prev $targetDir/{$wordCounter}_EN.jpg");
		exec("convert $sourceDir/$name.jpg -crop $widthx".($points[$i] - $prev)."+$w2+$prev $targetDir/{$wordCounter}_TH.jpg");
		$prev = $points[$i];
		$wordCounter++;
	}
}
