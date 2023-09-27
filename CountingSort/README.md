<h2>📍 Counting Sort Algorithm</h2>

<p>🔹 Counting Sort is a sorting algorithm that works well when you have a known range of integer input values. It is a non-comparison sorting algorithm, which means it doesn't rely on comparing elements to each other to sort them. Instead, it counts the number of occurrences of each element in the input and uses that information to sort the elements.</p>
<br />

<h3>📝 Here's a step-by-step explanation of how the Counting Sort algorithm works:</h3>

<p>1-<b>Determine the Range:</b> First, you need to determine the range of input values. You should know the minimum and maximum values present in the input array. Let's call these values min and max.</p>
<p>2-<b>Create a Counting Array:</b> Create an auxiliary array called the "counting array" with a size equal to the range of values plus one (max - min + 1). This array is used to store the count of each distinct input value. Initialize all elements of the counting array to zero.</p>
<p>3-<b>Count Occurrences:</b> Traverse the input array and, for each element, increment the corresponding index in the counting array. For example, if the input array contains the value 5, you would increment countingArray[5 - min].</p>
<p>4-<b>Calculate Cumulative Counts:</b> Modify the counting array so that each element at index i contains the sum of the counts of elements less than or equal to i. This step helps in determining the correct position of each element in the sorted output.</p>
<p>5-<b>Create a Temporary Array:</b> Create a temporary array of the same size as the input array to hold the sorted elements.</p>
<p>6-<b>Fill the Temporary Array:</b> Traverse the input array from left to right. For each element arr[i], find its position in the sorted order by looking up its count in the counting array and placing it at the correct position in the temporary array. Then, decrement the count in the counting array for that element.</p>
<p>7-<b>Copy to the Original Array:</b> After sorting all elements, copy the elements from the temporary array back to the original input array to obtain the sorted result.</p>
<br />
<br />
<p align="center">
  <image src="https://www.boardinfinity.com/blog/content/images/2023/03/Counting-sort.png" width="500" />
</p>
<br />
<br />
<h2>📚 Conclusions</h2>
    
<p>🔸 Counting Sort is stable (it preserves the relative order of equal elements) and performs well for sorting a large number of integers within a specific range, but it is not suitable for sorting data with a wide range or floating-point numbers since it relies on integer counts.</p>
