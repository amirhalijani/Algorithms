<h2>📍 Heap Sort Algorithm</h2>

<p>🔹 Heap Sort is a comparison-based sorting algorithm that falls under the category of in-place and unstable sorting algorithms. It was first proposed by J. W. J. Williams in 1964 and later refined by R. W. Floyd in 1965. Heap Sort is known for its efficient average and worst-case time complexities, making it suitable for sorting large datasets.</p>
<br />

<h3>📝 Here's a step-by-step explanation of how Heap Sort works:</h3>

<p>1-<b>Heapify the Array:</b></p>
<ul>
    <li>The first step in Heap Sort is to create a binary heap from the given array. A binary heap is a complete binary tree that satisfies the heap property, which means that the parent node's value is either greater than or less than its child nodes' values, depending on whether it's a max heap or a min heap.</li>
    <li>Starting from the last non-leaf node (usually located at index (n/2) - 1, where n is the number of elements in the array), we perform a process called "heapify" on each node.
    </li>
    <li>Heapify involves comparing the node with its children and swapping it with the largest (for max heap) or smallest (for min heap) child if necessary, and then recursively heapifying the affected child subtree.</li>
    <li>We repeat this process for all non-leaf nodes, working our way up the tree. After this step, the array is transformed into a valid max heap.</li>
</ul>
<p>2-<b>Extract Elements from Heap:</b></p>
<ul>
    <li>Once the array has been heapified, the largest (for max heap) or smallest (for min heap) element is at the root of the heap (i.e., the first element in the array).</li>
    <li>Swap this element with the last element in the array. This moves the maximum (or minimum) element to the end of the array.</li>
    <li>Decrease the heap size by 1 to exclude the last element from further consideration.</li>
</ul>
<p>3-<b>Heapify the Remaining Heap:</b></p>
<ul>
    <li>After extracting the maximum (or minimum) element, the heap property may be violated.
    </li>
    <li>To restore the heap property, heapify the root of the heap (which is now the first element in the array).</li>
    <li>This involves comparing the root with its children and swapping it with the largest (for max heap) or smallest (for min heap) child if necessary, and then recursively heapifying the affected child subtree.</li>
</ul>
<p>4-<b>Repeat Steps 2 and 3:</b></p>
<ul>
    <li>Repeat the process of extracting elements and heapifying the remaining heap until the entire array is sorted.</li>
    <li>The extracted elements are placed at the end of the array in descending order (for max heap) or ascending order (for min heap).</li>
</ul>

<p>5-<b>Result:</b></p>
<ul>
    <li>After repeating steps 2 and 3 for all elements in the array, you'll have a sorted array in either ascending or descending order, depending on whether you used a max heap or min heap.</li>
</ul>
<br />
<br />
<p align="center">
  <image src="https://www.boardinfinity.com/blog/content/images/2023/03/Heap-sort-algo.png" width="500" />
</p>
<br />
<br />
<h2>📚 Conclusion</h2>
    
<p>🔸 Heap Sort has a time complexity of O(n log n) for both its average and worst cases, making it more efficient than some other sorting algorithms like bubble sort or insertion sort, especially for large datasets. However, it's not as simple to implement as some of the other sorting algorithms.</p>
