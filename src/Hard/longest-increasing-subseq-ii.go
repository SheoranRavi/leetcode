// https://leetcode.com/problems/longest-increasing-subsequence-ii/

// solving it with usual dp results in TLE

package hard

func lengthOfLIS(nums []int, k int) int {
	// dp[i] = max(dp[i], dp[j]+1)
	// let's try to solve it the same way we did the longest increasing subsequence
	n := len(nums)
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && (nums[i]-nums[j] <= k) {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	maxLen := 1
	for _, lis := range dp {
		maxLen = max(maxLen, lis)
	}
	return maxLen
}

func lengthOfLIS(nums []int, k int) int {
	// dp[i] = max(dp[i], dp[j]+1)
	// let's try to solve it the same way we did the longest increasing subsequence
	// failed
	// use segment tree for efficient range max queries
	segTree := NewSegTree(100000)
	maxLen := 0
	for _, num := range nums {
		L := max(1, num-k)
		R := num - 1
		best := 0
		if L <= R {
			// query for max LIS for the range L - R
			best = segTree.Query(L, R, segTree.root)
		}
		dp := best + 1
		// For this num, dp is the LIS length so update it
		segTree.Update(num, dp, segTree.root)
		if dp > maxLen {
			maxLen = dp
		}
	}
	return maxLen
}

func (this *SegmentTree) Query(start, end int, node *Node) int {
	if node == nil {
		return this.minVal
	}
	if end < node.start || start > node.end {
		return this.minVal
	}
	if start == node.start && end == node.end {
		return node.val
	}

	// lies completely in one of the child nodes
	// or crosses over b/w the two child nodes
	mid := node.start + (node.end-node.start)/2
	if end <= mid {
		return this.Query(start, end, node.leftNode)
	} else if start > mid {
		return this.Query(start, end, node.rightNode)
	}

	leftMax := this.Query(start, mid, node.leftNode)
	rightMax := this.Query(mid+1, end, node.rightNode)
	return max(leftMax, rightMax)
}

func (this *SegmentTree) Update(idx, val int, node *Node) {
	if idx < node.start || idx > node.end {
		return
	}
	if node.start == node.end {
		node.val = val
		return
	}
	mid := node.start + (node.end-node.start)/2
	if idx <= mid {
		if node.leftNode == nil {
			node.leftNode = &Node{
				start: node.start,
				end:   mid,
			}
		}
		this.Update(idx, val, node.leftNode)
	} else {
		if node.rightNode == nil {
			node.rightNode = &Node{
				start: mid + 1,
				end:   node.end,
			}
		}
		this.Update(idx, val, node.rightNode)
	}
	leftMax := 0
	rightMax := 0
	if node.leftNode != nil {
		leftMax = node.leftNode.val
	}
	if node.rightNode != nil {
		rightMax = node.rightNode.val
	}
	childMax := max(leftMax, rightMax)
	node.val = max(node.val, childMax)
}

func NewSegTree(n int) SegmentTree {
	return SegmentTree{
		root: &Node{
			start: 1,
			end:   n,
		},
		minVal: 0,
	}
}

type SegmentTree struct {
	root   *Node
	minVal int // minVal gets used for a node that doesn't exist
}

type Node struct {
	leftNode, rightNode *Node
	start, end          int
	val                 int
}
