<?php	
// rename images of pages created from pdf:
// images from page 1 + 93
// images from page 2 + 94
// ...
// images from page 92 + 184

// simpler solution would be first create left side and rename to odd numbers
// then create right side and rename to even numbers...

$folder = "words_3_cut/"

$files = scandir($folder);
array_shift($files);
array_shift($files);
array_shift($files);

$new = [];
$added = [];
$a = 1;
$b = 1;
for ($i=1; $i < 185; $i++) { 
	if ($i < 93) {
		if ($i % 2 == 0) {
			$new["$i"] = $a + 92;
			$a++;
		} else {
			$new["$i"] = $b;
			$b++;
		}
	} else {
		if ($i % 2 == 0) {
			$new["$i"] = 91 + ceil(($i + 1) / 2);
		} else {
			$new["$i"] = ceil(($i + 1) / 2);
		}
		$a++;
	}
}


$i = 1;
foreach ($new as $key => $value) {
	exec("cp ./$value.jpg ./a/$key.jpg");
    $i++;
};
