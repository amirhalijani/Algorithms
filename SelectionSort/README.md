<h2>📍 Selection Sort Algorithm</h2>

<p>🔹 Selection sort is a simple comparison-based sorting algorithm used to arrange elements in a specific order, typically ascending or descending. It works by repeatedly selecting the minimum (or maximum) element from the unsorted part of the list and placing it at the beginning (or end) of the sorted part. The algorithm continues to do this until the entire list is sorted. Selection sort has an average and worst-case time complexity of O(n^2), making it inefficient for large lists but suitable for small datasets due to its simplicity.</p>
<br />

<h3>📝 Here's a step-by-step explanation of how the selection sort algorithm works:</h3>
<p>1-<b>Initialization:</b> The list is divided into two parts: the sorted part on the left and the unsorted part on the right. Initially, the sorted part is empty, and the unsorted part contains all the elements.</p>
<p>2-<b>Finding the Minimum / Maximum Element:</b> The algorithm searches for the minimum (or maximum) element in the unsorted part of the list. To do this, it iterates through the unsorted part, comparing each element with the current minimum (or maximum) found so far. If a smaller element (or larger element) is encountered, it becomes the new minimum (or maximum).</p>
<p>3-<b>Swapping:</b> Once the minimum element (or maximum) in the unsorted part is found, it is swapped with the leftmost (or rightmost) element in the unsorted part. This effectively moves the minimum (or maximum) element to the end of the sorted part and expands the sorted part by one element. The unsorted part becomes one element shorter.</p>
<p>4-<b>Repeat:</b> Repeat: Steps 2 and 3 are repeated until the entire list is sorted. In each iteration, the smallest (or largest) remaining element in the unsorted part is moved to the sorted part.</p>
<p>5-<b>Termination:</b> The algorithm terminates when the entire list is sorted, and the sorted part covers the entire list.</p>
<br />
<br />
<p align="center">
  <image src="https://www.simplilearn.com/ice9/free_resources_article_thumb/Selection-Sort-Soni/what-is-selection-sort.png" />
</p>
<br />
<br />
<h2>📚 Conclusions</h2>
    
<p>🔸 Selection sort has a time complexity of O(n^2) in the worst case, where 'n' is the number of elements in the list. This makes it inefficient for large lists. It also performs a fixed number of swaps (O(n)), which can be advantageous in situations where minimizing data movement is a priority, such as in embedded systems or when the cost of swaps is high. Despite its simplicity, selection sort is not commonly used in practice for sorting large datasets, as more efficient sorting algorithms like merge sort, quicksort, or heapsort typically offer better performance. However, selection sort can be a good choice for small lists or as a learning tool to understand sorting algorithms' basic principles.</p>
