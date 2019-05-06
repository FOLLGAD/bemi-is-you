import * as PIXI from 'pixi.js';
import Level from './level';

const TILE_SIZE = 64 // 64*64 pixels represents a tile

// Events
const MOVE = 0,
	DEATH = 1,
	SPAN = 2
// Transform: death+spawn

export default class Game {
	constructor() {
		this.app = new PIXI.Application({
			width: TILE_SIZE * 16, height: TILE_SIZE * 12, resolution: window.devicePixelRatio || 1, forceCanvas: true, backgroundColor: 0xeeeeee,
		})

		document.body.appendChild(this.app.view)

	}
	start() {
		this.chars = new Map()

		let app = this.app

		const background = PIXI.Texture.from('../textures/background.png')
		// background.width = app.view.width
		// background.height = app.view.height
		// Create the background

		let bgElement = new PIXI.Sprite.from(background)
		bgElement.width = app.view.width
		bgElement.height = app.view.height
		app.stage.addChild(bgElement)

		const container = new PIXI.Container()
		this.container = container
		container.x = 0
		container.y = 0

		app.stage.addChild(container)
	}
	updateChar(characterId, x, y) {
		let char = this.chars.get(characterId)

		if (char) {
			char.x = TILE_SIZE * x
			char.y = TILE_SIZE * y
		} else {
			const char = new PIXI.Graphics()
			char.beginFill(0x0011ee) // Fill colour
			char.drawRect(0, 0, TILE_SIZE, TILE_SIZE) // Size
			char.x = TILE_SIZE * x
			char.y = TILE_SIZE * y
			this.container.addChild(char)
			this.chars.set(characterId, char)
		}
	}
	removeChar(characterId) {
		let index = this.chars.findIndex(ch => ch.id == characterId)
		if (index !== -1) {
			this.chars.splice(index, 1)
		}
	}
	setLevel(data) {
		this.level = new Level(data)

		this.level.objects.forEach(ob => {
			this.updateChar(ob.id, ob.pos.x, ob.pos.y)
		})
	}
	listen(elem, func) {
		elem.addEventListener('keydown', func)
	}
	deltaTick(tick) {
		tick.forEach(change => {
			let { event, id, pos } = change

			switch (event) {
				case 0: // Move
					let char = this.level.objects.find(ob => ob.id == id)
					char.pos.x += pos.x
					char.pos.y += pos.y
					this.updateChar(char.id, char.pos.x, char.pos.y)
			}
		})
	}
}