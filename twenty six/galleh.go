package twentysix

type goyoon struct {
	next *goyoon
	name string
}

type galleh struct {
	goyoon
}

func addGoyoon(g *galleh, name string) {
	newGoyoon := &goyoon{
		name: name,
	}

	if g.next == nil {
		g.next = newGoyoon
	} else {
		node := g.next
		for node.next != nil {
			node = node.next
		}
		node.next = newGoyoon
	}
}

func twentysix() {
	g := &galleh{}
	addGoyoon(g, "babaii")
	addGoyoon(g, "candy")
	addGoyoon(g, "black")

}
