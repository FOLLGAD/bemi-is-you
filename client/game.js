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
		PIXI.settings.SCALE_MODE = PIXI.SCALE_MODES.NEAREST
		this.app = new PIXI.Application({
			width: TILE_SIZE * 16,
			height: TILE_SIZE * 12,
			resolution: window.devicePixelRatio || 1,
			backgroundColor: 0xeeeeee,
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
	setChar(char) {
		const texture = PIXI.Texture.from('../textures/' + this.getTexture(char))
		const pixiChar = new PIXI.Sprite(texture)
		const scaleY = 4
		const scaleX = scaleY
		pixiChar.scale.set(scaleX, scaleY)
		pixiChar.x = TILE_SIZE * char.pos.x
		pixiChar.y = TILE_SIZE * char.pos.y
		this.container.addChild(pixiChar)
		this.chars.set(char.id, pixiChar)
	}
	updateChar(charObject) {
		let char = this.chars.get(charObject.id)

		if (char) {
			char.x = TILE_SIZE * charObject.pos.x
			char.y = TILE_SIZE * charObject.pos.y
		}
	}
	getTexture(object) {
		if (object.kind == 0) return object.item + "_block.png"
		else return object.item + ".png"
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
			this.setChar(ob)
		})
	}
	listen(elem, func) {
		elem.addEventListener('keydown', func)
	}
	deltaTick(tick) {
		if (tick != null) {
			tick.forEach(change => {
				let { event, id, pos } = change

				switch (event) {
					case MOVE: // Move
						let char = this.level.objects.find(ob => ob.id == id)
						char.pos.x += pos.x
						char.pos.y += pos.y
						this.updateChar(char)
				}
			})
		}
	}
}