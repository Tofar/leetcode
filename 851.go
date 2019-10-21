// 多叉树，有向
type MultyTree struct {
    Label int
    Quiet int
    AnswerTree *MultyTree
    Richer []*MultyTree
}

func loudAndRich(richer [][]int, quiet []int) []int {
    trees := make(map[int]*MultyTree, len(quiet))
    for l, q := range quiet {
        t := &MultyTree{
            Label: l,
            Quiet: q,
            Richer: make([]*MultyTree, 0),
        }
        trees[l] = t
    }
    
    for _, r := range richer {
        x, y := trees[r[0]], trees[r[1]]
        y.Richer = append(y.Richer, x)
    }
    
    answer := make([]int, len(quiet))
    for i := 0; i < len(answer); i++ {
        t := help(trees[i])
        answer[i] = t.Label
    }
    return answer
}

func help(root *MultyTree) *MultyTree {
    if root == nil {
        return nil
    }
    if root.AnswerTree != nil {
        return root.AnswerTree
    }
    tree := root
    for _, r := range root.Richer {
        temp := help(r)
        if temp.Quiet < tree.Quiet {
            tree = temp
        }
    }
    
    root.AnswerTree = tree
    return tree
}
