package main

var tree = &BinaryNode{
	Value: 20,
	right: &BinaryNode{
		Value: 50,
		right: &BinaryNode{
			Value: 100,
		},
		left: &BinaryNode{
			Value: 30,
			right: &BinaryNode{
				Value: 45,
			},
			left: &BinaryNode{
				Value: 29,
			},
		},
	},
	left: &BinaryNode{
		Value: 10,
		right: &BinaryNode{
			Value: 15,
		},
		left: &BinaryNode{
			Value: 5,
			right: &BinaryNode{
				Value: 7,
			},
		},
	},
}

var tree2 = &BinaryNode{
	Value: 20,
	right: &BinaryNode{
		Value: 50,
		left: &BinaryNode{
			Value: 30,
			right: &BinaryNode{
				Value: 45,
				right: &BinaryNode{
					Value: 49,
				},
				left: &BinaryNode{
					Value: 29,
					left: &BinaryNode{
						Value: 21,
					},
				},
			},
		},
	},
	left: &BinaryNode{
		Value: 10,
		right: &BinaryNode{
			Value: 15,
		},
		left: &BinaryNode{
			Value: 5,
			right: &BinaryNode{
				Value: 7,
			},
		},
	},
}