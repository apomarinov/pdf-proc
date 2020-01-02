<?php

$sourceDir = "words_5_pairs";
$ocrOutput = "ocr.json";

chdir($sourceDir);

$files = scandir(".");
array_shift($files);
array_shift($files);
array_shift($files);

$output = [];
for ($i=1; $i < count($files); $i++) {
	if (strpos($files[$i], "EN") === false) {
		continue;
	}
	exec("tesseract -l eng ".$files[$i]." output");
	if (!is_readable("output.txt")) {
		echo $files[$i]." OCR Failed.\n";
		continue;
	}
	$output[$files[$i]] = rtrim(rtrim(file_get_contents("output.txt"), "\n"), "\f");
	file_put_contents($ocrOutput, json_encode($output, JSON_UNESCAPED_UNICODE | JSON_PRETTY_PRINT));
}