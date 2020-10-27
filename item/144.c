// 二叉树前序遍历，很简单

/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */


/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

void dfs(struct TreeNode* root,int* returnSize,int* res){
    if (root==NULL) return;
    res[*returnSize]=root->val;
    (*returnSize)++;
    dfs(root->left,returnSize,res);
    dfs(root->right,returnSize,res);
}

int* preorderTraversal(struct TreeNode* root, int* returnSize){
    int* res = (int*)malloc(sizeof(int)*100);
    *returnSize = 0;
    dfs(root,returnSize,res);
    return res;
}
