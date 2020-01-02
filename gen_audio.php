<?php 
/*
Gen audio for text from json:
[
	{
		"word": "มานี"
	}
]
*/


$wordsFile = "ocr.json";

$folder = "audio";

$makeAudio = 'say -v Samantha "{text}" -r 150 -o "'.$folder.'/{filename}.aiff"';
$convertToMp3 = "lame -m m {$folder}/{filename}.aiff {$folder}/{filename}.mp3";

$words1 = json_decode(file_get_contents($wordsFile), JSON_UNESCAPED_UNICODE);
$words = [];
foreach ($words1 as $key => $value) {
	$words[] = [
		"word" => $value,
		"name" => str_replace(".jpg", "", $key)
	];
}

$failed = [];
$failedConvert = [];
for ($i=0; $i < count($words); $i++) { 
	$filename = $words[$i]["name"];
	$cmd = str_replace("{text}", $words[$i]['word'], $makeAudio);
	$cmd = str_replace("{filename}", $filename, $cmd);
	shell_exec("$cmd");
	if (!file_exists("{$folder}/".$filename.".aiff")) {
		$failed[] = $filename;
		continue;
	}
	$cmd = str_replace("{filename}", $filename, $convertToMp3);
	shell_exec($cmd);
	if (!file_exists("{$folder}/".$filename.".mp3")) {
		$failedConvert[] = $filename;
		continue;
	}
	unlink("{$folder}/".$filename.".aiff");
}

echo json_encode($failed, JSON_UNESCAPED_UNICODE | JSON_PRETTY_PRINT)."\n";
echo json_encode($failedConvert, JSON_UNESCAPED_UNICODE | JSON_PRETTY_PRINT);
