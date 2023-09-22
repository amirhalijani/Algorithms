<h2>📍 Merge Sort Algorithm</h2>

<p>🔹 Merge Sort is a widely-used, efficient, and stable comparison-based sorting algorithm. It follows the divide-and-conquer paradigm to sort an array or list of elements. The basic idea behind Merge Sort is to divide the input into smaller, more manageable parts, sort those smaller parts, and then merge them back together to obtain a sorted array.</p>
<br />

<h3>📝 Here's a step-by-step explanation of the Merge Sort algorithm:</h3>
<p>1-<b>Divide:</b> The first step is to divide the input array into two roughly equal halves. This is done recursively until each subarray contains only one element, as a single element is always considered sorted.</p>
<p>2-<b>Conquer:</b> Once you have divided the array into smaller parts, you recursively sort each of these smaller subarrays. This is typically done by calling the Merge Sort algorithm recursively on each half.</p>
<p>3-<b>Merge:</b> After all the smaller subarrays are sorted, the merging step comes into play. In this step, you take two sorted subarrays and merge them to create a larger, sorted subarray. The merging process involves comparing elements from both subarrays and placing them in the correct order in a new, temporary array.</p>
<ul>
    <li>Start with two pointers, one for each subarray, initially pointing to the first element in each subarray.</li>
    <li>Compare the elements pointed to by the two pointers and select the smaller one.</li>
    <li>Place the smaller element in the temporary array.</li>
    <li>Move the pointer in the subarray from which the smaller element was selected to the next element.</li>
    <li>Repeat these steps until you've merged all elements from both subarrays into the temporary array.</li>
    <li>Finally, replace the original subarray with the merged, sorted subarray.</li>
</ul>
<p>4-<b>Repeat and Combine:</b>Continue this process of dividing, sorting, and merging until you have merged all the subarrays back together into a single sorted array. This final merged array is your sorted result.</p>
<br />
<br />
<p align="center">
  <p><a href="https://commons.wikimedia.org/wiki/File:Merge-sort-example-300px.gif#/media/File:Merge-sort-example-300px.gif"><img src="https://upload.wikimedia.org/wikipedia/commons/c/cc/Merge-sort-example-300px.gif" alt="Merge-sort-example-300px.gif" height="180" width="300"></a><br>By &lt;a href="//commons.wikimedia.org/w/index.php?title=User:Swfung8&amp;amp;action=edit&amp;amp;redlink=1" class="new" title="User:Swfung8 (page does not exist)"&gt;Swfung8&lt;/a&gt; - &lt;span class="int-own-work" lang="en"&gt;Own work&lt;/span&gt;, <a href="https://creativecommons.org/licenses/by-sa/3.0" title="Creative Commons Attribution-Share Alike 3.0">CC BY-SA 3.0</a></p>
</p>
<br />
<br />
<h2>📚 Conclusion</h2>
    
<p>🔸 The key characteristics of Merge Sort are its stability (the relative order of equal elements is preserved) and its consistent O(n log n) time complexity for the worst-case, average-case, and best-case scenarios. However, it does require additional space for the temporary arrays used in the merging process, which can make it less efficient in terms of memory usage compared to some other sorting algorithms. Nevertheless, it is a robust and widely-used sorting algorithm, especially in situations where stable sorting or predictable performance is crucial.</p>
