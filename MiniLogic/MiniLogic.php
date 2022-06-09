<?php
// header('Content-Type: application/json');

if (empty($_GET['question'])) {
	echo "<strong>Please provide parameters. E.g.:</strong><br>
		Question 1: question=1&expectedDate=dd mm yyyy&returnedDate=dd mm yyyy <br>
		Question 2: question=2&student=1&candies=1&firstStudent=1 <br>
		Question 3: question=3&arrayLength=1&array=[1,2,3] <br>";
	return;
}

switch ($_GET['question']) {
	case '1':
		if (empty($_GET['expectedDate']) || empty($_GET['returnedDate'])) {
			echo "Please use this example format as parameters: <strong>question=1&expectedDate=dd mm yyyy&returnedDate=dd mm yyyy</strong>";
			return;
		}
		
		question_1($_GET['expectedDate'], $_GET['returnedDate']);
		break;

	case '2':
		if (empty($_GET['student']) || empty($_GET['candies']) || empty($_GET['firstStudent'])) {
			echo "Please use this example format as parameters: <strong>question=2&student=1&candies=1&firstStudent=1</strong>";
			return;
		}

		question_2($_GET['student'], $_GET['candies'], $_GET['firstStudent']);
		break;

	case '3':
		if (empty($_GET['arrayLength']) || empty($_GET['array'])) {
			echo "Please use this example format as parameters: <strong>question=3&arrayLength=1&array=[1,2,3]</strong>";
			return;
		}

		$temp = str_replace("[", "", $_GET['array']);
		question_3($_GET['arrayLength'], explode(",", str_replace("]", "", $temp)));
		break;
	
	default:
		echo "Question did not found! Available question only for: 1, 2, & 3.";
		return;
		break;
}

function question_1($expectedDate, $returnedDate)
{
	#region Variables
	$expectedDate = explode(" ", $expectedDate);
	$returnedDate = explode(" ", $returnedDate);

	$dateReturn = date_create(date("Y-m-d", strtotime($returnedDate[0] ."-". $returnedDate[1] ."-". $returnedDate[2])));
	$dateExpect = date_create(date("Y-m-d", strtotime($expectedDate[0] ."-". $expectedDate[1] ."-". $expectedDate[2])));
	$differentMonths = date_diff($dateReturn, $dateExpect); 

	$yearsFine = 12000;
	#endregion

	#region Validations
	if (count($expectedDate) < 3 || count($returnedDate) < 3) {
		echo("Please use this format (dd mm yyyy).");
		return;
	}

	if (strtotime($expectedDate[0] ."-". $expectedDate[1] ."-". $expectedDate[2]) === false || strtotime($returnedDate[0] ."-". $returnedDate[1] ."-". $returnedDate[2]) === false) {
		echo("One of the given values is not a date therefore we cannot process it. Make sure use this format (dd mm yyyy).");
		return;
	}

	if (strtotime($returnedDate[0] ."-". $returnedDate[1] ."-". $returnedDate[2]) < strtotime($expectedDate[0] ."-". $expectedDate[1] ."-". $expectedDate[2])) {
		echo("Success returning.");
		return;
	}
	#endregion
	
	#region Outputs
	if ($differentMonths->days > 365) {
		echo $yearsFine;
		return;
	}elseif ($differentMonths->m > 0) {
		echo ($differentMonths->m * 500);
		return;
	}elseif ($differentMonths->m < 1) {
		echo ($differentMonths->d * 15);
		return;
	}
	#endregion
}

function question_2(int $student, int $candies, int $firstStudent)
{
	#region Variables
	$totalGivenCandy = 0;
	$firstStudent = $firstStudent;
	$lastStudent = 0;
	#endregion

	#region Validations
	if ($firstStudent > $student) {
		echo "firstStudent value should be lower than the total of students.";
		return;
	}
	#endregion

	#region Outputs
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
	return;
	#endregion

}

function question_3(int $arrayLength, array $array)
{
	#region Variables
	$arrayTotal = count($array);
	$middleArrayIndex = is_float($arrayTotal / 2) ? round($arrayTotal / 2)  : ($arrayTotal / 2) + 1;
	$leftValue = 0;
	$rightValue = 0;
	$count = 0;
	#endregion

	#region Validations
	if ($arrayLength != $arrayTotal) {
		echo "Length of array is not equal with the total of array.";
		return;
	}
	#endregion

	#region Outputs
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
		return;
	}else {
		echo("NO");
		return;
	}
	#endregion
}