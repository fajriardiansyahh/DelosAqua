<?php

// question_1("1 1 2023",  12 ." ". 1 ." ". 2023);
// question_2(4,2,2);
// question_3(5, [3,7,9,10,0]);

function question_1($expectedDate, $returnedDate)
{
	$expectedDate = explode(" ", $expectedDate);
	$returnedDate = explode(" ", $returnedDate);
	$dateReturn = new Datetime(date("Y-m-d", strtotime($returnedDate[0] ."-". $returnedDate[1] ."-". $returnedDate[2])));
	$dateExpect = new Datetime(date("Y-m-d", strtotime($expectedDate[0] ."-". $expectedDate[1] ."-". $expectedDate[2])));
	$differentMonths = date_diff($dateReturn, $dateExpect)->format('%m') < 1 ? + 1 : date_diff($dateReturn, $dateExpect)->format('%m'); 
	
	if (count($expectedDate) < 3 || count($returnedDate) < 3) {
		return var_dump("Please use this format (dd mm yyyy).");
	}

	if (strtotime($expectedDate[0] ."-". $expectedDate[1] ."-". $expectedDate[2]) === false || strtotime($returnedDate[0] ."-". $returnedDate[1] ."-". $returnedDate[2]) === false) {
		return var_dump("One of the given values is not a date therefore we cannot process it. Make sure use this format (dd mm yyyy).");
	}

	if (strtotime($returnedDate[0] ."-". $returnedDate[1] ."-". $returnedDate[2]) < strtotime($expectedDate[0] ."-". $expectedDate[1] ."-". $expectedDate[2])) {
		return var_dump("Success returning.");
	}

	if ($returnedDate[1] == $expectedDate[1] && $returnedDate[2] == $expectedDate[2] ) {
		$daysFine = ($returnedDate[0] - $expectedDate[0]) * 15;
		echo $daysFine;
	}elseif ($differentMonths > 12) {
		$yearsFine = 12000;
		echo $yearsFine;
	}else {
		$monthsFine = $differentMonths * 500;
		echo $monthsFine;
	}
}

function question_2(int $student, int $candies, int $firstStudent)
{
	$totalGivenCandy = 0;
	$firstStudent = $firstStudent;
	$lastStudent = 0;

	if ($firstStudent > $student) {
		return var_dump("First Student value exceeded total of students");
	}

	for ($i=0; $i < $candies; $i++) { 
		for ($j = $firstStudent; $j <= $student; $j++) { 
			
			if ($j == $student) {
				$students[$j][] = ($i + 1);
				$totalGivenCandy++;
				$firstStudent = 1;
			} else {
				$students[$j][] = ($i + 1);
				$totalGivenCandy++;
				$i++;
			}

			if ($i == $candies) {
				$lastStudent = $j;
				break;
			}
		}
	}
	echo($lastStudent);
	var_dump($students);
}

function question_3(int $arrayLength, array $array)
{
	$arrayTotal = count($array);
	$middleArrayIndex = is_float($arrayTotal / 2) ? round($arrayTotal / 2)  : ($arrayTotal / 2) + 1;
	$leftValue = 0;
	$rightValue = 0;
	$count = 0;

	if ($arrayLength != $arrayTotal) {
		return var_dump("Length of array is not equal with the total of array");
	}

	for ($i=0; $i < $arrayTotal ; $i++) { 
		if ($count == ($middleArrayIndex - 1)) {
			$rightValue = $rightValue + $array[$i];
		}else {
			$leftValue = $leftValue + $array[$i];
			$count++;
		}
	}
	
	$rightValue = $rightValue - $array[($middleArrayIndex - 1)];
	
	if ($rightValue == $leftValue) {
		echo("YES");
	}else {
		echo("NO");
	}
}