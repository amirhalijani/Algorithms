<h2>📍 Quick Sort Algorithm</h2>

<p>🔹 Quick Sort is a widely used sorting algorithm that follows the "divide and conquer" strategy to efficiently sort a list or array of elements. It was developed by Tony Hoare in 1960 and is known for its speed and effectiveness in practice. Quick Sort works by selecting a "pivot" element from the array and partitioning the other elements into two sub-arrays, according to whether they are less than or greater than the pivot. The sub-arrays are then recursively sorted.</p>
<br />

<h3>📝 Here's a step-by-step explanation of the Quick Sort algorithm:</h3>

<p>1-<b>Choose a Pivot:</b> The first step is to select a pivot element from the array. The choice of the pivot can significantly affect the algorithm's performance. Common strategies include selecting the first element, the last element, the middle element, or even a random element. The choice of the pivot is often made to minimize the chances of encountering worst-case scenarios.</p>
<p>2-<b>Partitioning:</b> Rearrange the array so that all elements less than the pivot are on its left, and all elements greater than the pivot are on its right. This process is often referred to as partitioning. It can be done using two pointers: one starting from the left end and moving right, and the other starting from the right end and moving left. When they meet, swap the elements they point to if they are out of place (i.e., the element on the left is greater than the pivot, and the element on the right is less than the pivot). Continue this process until the pointers cross each other.</p>
<p>3-<b>Recursion:</b> After partitioning, you have two sub-arrays: one containing elements less than the pivot and one containing elements greater than the pivot. Recursively apply the Quick Sort algorithm to these sub-arrays. This means you will choose a pivot for each sub-array and partition it until the sub-arrays are sorted.</p>
<p>4-<b>Base Case:</b> The base case for recursion is when a sub-array has zero or one element, in which case it is already sorted by definition.</p>
<p>5-<b>Combine:</b> There's no explicit combine step in Quick Sort since the elements are sorted in place.</p>
<p>6-<b>Repeat:</b> Continue partitioning and sorting the sub-arrays until the entire array is sorted.</p>
<p>7-<b>Final Result:</b> Once the recursion is complete, and all sub-arrays are sorted, the original array will also be sorted.</p>
<br />
<br />
<p align="center">
  <image src="https://upload.wikimedia.org/wikipedia/commons/6/6a/Sorting_quicksort_anim.gif" width="500" />
</p>
<br />
<br />
<h2>📚 Conclusion</h2>
    
<p>🔸 Quick Sort is known for its average-case time complexity of O(n log n), where 'n' is the number of elements in the array. However, in the worst-case scenario (when the pivot selection consistently results in unbalanced partitions), it can have a time complexity of O(n^2). To mitigate this, various strategies for pivot selection, such as choosing the median of three random elements, are often employed to improve performance in practice.</p>
