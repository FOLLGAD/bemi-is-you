class Pos {
	constructor({ x, y }) {
		this.x = x
		this.y = y
	}
}

export default class Object {
	constructor({ pos, id, item, kind }) {
		this.pos = pos
		this.id = id
		this.item = item
		this.kind = kind
	}
}