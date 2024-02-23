package waterbottles

// ban đầu, số lượng bình ta có thể uống nhỏ nhất bằng số bình ta có, số bình trống cũng bằng số bình ta uống.
// ta sẽ tiến hành kiểm tra bằng cách rà xem bao giờ số bình rỗng nhỏ hơn số bình ta có thể đổi được. vậy sẽ là 1 vòng lặp for, có thể là O(n)

func NumWaterBottles(numsBottles int, numExchange int) int{
	// số bình ít nhất có thể uống
	var result int = numsBottles
	var emptyBottles int= numsBottles
	
	 for emptyBottles >= numExchange {
		var bottlesExchanged = emptyBottles/numExchange // 1 op
		result+=bottlesExchanged //1 op
		emptyBottles = bottlesExchanged + emptyBottles%numExchange // 1op
	 } 

	 return result;

}