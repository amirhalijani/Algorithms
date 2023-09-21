<h2>📍 Bubble Sort Algorithm</h2>

<p>🔹 Bubble sort is a simple sorting algorithm that repeatedly steps through a list of elements (such as an array) and compares adjacent elements. If they are in the wrong order, it swaps them. The pass through the list is repeated until no swaps are needed, which indicates that the list is sorted. Bubble sort is named because smaller elements "bubble" to the top of the list during each pass.</p>
<br />

<h3>📝 Here's a step-by-step explanation of how the bubble sort algorithm works:</h3>
<p>1- Start with the first element (index 0) in the list.</p>
<p>2- Compare the current element with the next element (index 1). If the current element is greater than the next element, swap them.</p>
<p>3- Move to the next pair of elements (index 1 and index 2) and compare them. Again, if the current element is greater than the next element, swap them.</p>
<p>4- Repeat this process for each adjacent pair of elements in the list, moving from left to right, until you reach the end of the list. This completes one pass through the list.</p>
<p>5- After the first pass, the largest element will have "bubbled up" to the end of the list. You can think of this as the largest element "sinking" to its correct position.</p>
<p>6- Repeat the process for the entire list, excluding the last element (which is already in its correct position after the first pass).</p>
<p>7- Continue these passes through the list until no more swaps are needed during a pass. This means the entire list is sorted.</p>
<br />
<br />
<p align="center">
  <image src="https://www.computersciencebytes.com/wp-content/uploads/2016/10/bubble_sort.png" />
</p>
<br />
<br />
<h2>📚 Conclusion</h2>
    
<p>🔸 Bubble sort has a time complexity of O(n^2) in the worst and average cases, where 'n' is the number of elements in the list. It is not considered efficient for large lists, but it is straightforward to understand and implement, making it a useful educational tool and sometimes practical for small lists or nearly sorted data. There are more efficient sorting algorithms, such as quicksort or mergesort, that are generally preferred for larger datasets.</p>
