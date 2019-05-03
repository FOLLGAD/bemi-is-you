class Pos {
	x;
	y;
	constructor(x,y) {
		this.x = x
		this.y = y
	}
}

export default class Object {
	pos;
	constructor(pos, id) {
		this.pos = pos
	}
}