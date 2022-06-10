package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func QuestionsPage(w http.ResponseWriter, r *http.Request) {
	no := r.URL.Query().Get("no")
	err_msg := `Please provide parameters. E.g.:
Question 1: no=1&expectedDate=dd mm yyyy&returnedDate=dd mm yyyy 
Question 2: no=2&student=1&candies=1&firstStudent=1 
Question 3: no=3&arrayLength=1&array=[1,2,3] `

	if len(no) < 1 {
		w.Write([]byte(err_msg))
		return
	}

	switch no {
	case "1":
		returnedDate := r.URL.Query().Get("returnedDate")
		expectedDate := r.URL.Query().Get("expectedDate")
		err_msg := "Please use this example format as parameters: no=1&expectedDate=dd mm yyyy&returnedDate=dd mm yyyy"

		if len(returnedDate) < 1 || len(expectedDate) < 1 {
			http.Error(w, err_msg, http.StatusBadRequest)
			return
		}

		returnedDate_split := strings.Split(returnedDate, " ")
		returnedDateSplit_days, err := strconv.Atoi(returnedDate_split[0])
		if err != nil {
			http.Error(w, "Error not integer", http.StatusBadRequest)
			return
		}
		returnedDateSplit_month, err := strconv.Atoi(returnedDate_split[1])
		if err != nil {
			http.Error(w, "Error not integer", http.StatusBadRequest)
			return
		}
		returnedDateSplit_year, err := strconv.Atoi(returnedDate_split[2])
		if err != nil {
			http.Error(w, "Error not integer", http.StatusBadRequest)
			return
		}

		expectedDate_split := strings.Split(expectedDate, " ")
		expectedDateSplit_days, err := strconv.Atoi(expectedDate_split[0])
		if err != nil {
			http.Error(w, "Error not integer", http.StatusBadRequest)
			return
		}
		expectedDateSplit_month, err := strconv.Atoi(expectedDate_split[1])
		if err != nil {
			http.Error(w, "Error not integer", http.StatusBadRequest)
			return
		}
		expectedDateSplit_year, err := strconv.Atoi(expectedDate_split[2])
		if err != nil {
			http.Error(w, "Error not integer", http.StatusBadRequest)
			return
		}

		if returnedDateSplit_year > expectedDateSplit_year {
			var yearsFine = "12000"
			w.Write([]byte(yearsFine))
			fmt.Println(yearsFine)
			return
		} else if returnedDateSplit_month > expectedDateSplit_month {
			var monthsFine = (returnedDateSplit_month - expectedDateSplit_month) * 500

			fmt.Println(monthsFine)
			return
		} else if returnedDateSplit_days > expectedDateSplit_days {
			var daysFine = (returnedDateSplit_days - expectedDateSplit_days) * 15

			fmt.Println(daysFine)
			return
		}

		return

	case "2":
		student := r.URL.Query().Get("student")
		candies := r.URL.Query().Get("candies")
		firstStudent := r.URL.Query().Get("firstStudent")
		err_msg := "Please use this example format as parameters: no=2&student=1&candies=1&firstStudent=1"

		if len(student) < 1 || len(candies) < 1 || len(firstStudent) < 1 {
			http.Error(w, err_msg, http.StatusBadRequest)
			return
		}

		student_int, err := strconv.Atoi(student)
		if err != nil {
			http.Error(w, "Error not integer", http.StatusBadRequest)
			return
		}

		candies_int, err := strconv.Atoi(candies)
		if err != nil {
			http.Error(w, "Error not integer", http.StatusBadRequest)
			return
		}

		firstStudent_int, err := strconv.Atoi(firstStudent)
		if err != nil {
			http.Error(w, "Error not integer", http.StatusBadRequest)
			return
		}

		last_student := 0
		for i := 0; i < candies_int; i++ {
			for j := firstStudent_int; j <= student_int; j++ {
				if j < student_int {
					i++
				}

				if i == candies_int {
					last_student = j
					break
				}
			}
		}

		fmt.Println(last_student)
		return

	case "3":
		arrayLength := r.URL.Query().Get("arrayLength")
		array := r.URL.Query().Get("array")
		err_msg := "Please use this example format as parameters: no=3&arrayLength=1&array=[1,2,3]"

		if len(arrayLength) < 1 || len(array) < 1 {
			http.Error(w, err_msg, http.StatusBadRequest)
			return
		}

		arrayLength_int, err := strconv.Atoi(arrayLength)
		if err != nil {
			http.Error(w, "Error not integer", http.StatusBadRequest)
			return
		}
		temp_1 := strings.Replace(array, "[", "", 1)
		temp_2 := strings.Replace(temp_1, "]", "", 1)
		new_array := strings.Split(temp_2, " ")

		if arrayLength_int != len(new_array) {
			http.Error(w, "The array length value should be the same as in the array", http.StatusNotFound)
			return
		}

		// for i := 0; i < len(new_array); i++ {

		// }

		middleArrayIndex := 0
		count := 0
		leftValue := 0
		rightValue := 0
		for i := 0; i < len(new_array); i++ {
			if len(new_array) == middleArrayIndex {

			} else {
				leftValue = leftValue + len(new_array)
				count++
			}
		}

		if rightValue == leftValue {
			fmt.Println("YES")
			return
		} else {
			fmt.Println("NO")
			return
		}

	default:
		http.Error(w, "Question did not found! Available question only for: 1, 2, & 3.", http.StatusNotFound)
		return
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	err_msg := `Please provide parameters. E.g.:
Question 1: no=1&expectedDate=dd mm yyyy&returnedDate=dd mm yyyy 
Question 2: no=2&student=1&candies=1&firstStudent=1 
Question 3: no=3&arrayLength=1&array=[1,2,3] `

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("This is just a home page, please visit: /question. " + err_msg))
}
