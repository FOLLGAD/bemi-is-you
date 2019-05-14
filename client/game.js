import * as PIXI from 'pixi.js';
import Level from './level';

const TILE_SIZE = 64 // 64*64 pixels represents a tile

// Events
const MOVE = 0,
	DEATH = 1,
	SPAWN = 2
// Transform: death+spawn

export default class Game {
	constructor() {
		PIXI.settings.SCALE_MODE = PIXI.SCALE_MODES.NEAREST
		this.app = new PIXI.Application({
			width: TILE_SIZE * 18,
			height: TILE_SIZE * 12,
			resolution: window.devicePixelRatio || 1,
			backgroundColor: 0xeeeeee,
		})
		this.playerNum = null
	}
	start() {
		document.body.appendChild(this.app.view)

		this.chars = new Map()

		const background = PIXI.Texture.from('../textures/background.png')
		let bgElement = new PIXI.Sprite.from(background)
		bgElement.width = this.app.view.width
		bgElement.height = this.app.view.height
		this.app.stage.addChild(bgElement)

		const container = new PIXI.Container()
		this.container = container
		container.x = 0
		container.y = 0

		this.app.stage.addChild(container)

		this.level.objects.forEach(ob => {
			this.setChar(ob)
		})
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
		if (object.item == "1" || object.item == "2") {
			if (this.playerNum == Number(object.item)) {
				return "1.png"
			} else {
				return "2.png"
			}
		} else if (object.kind == 0) return object.item + "_block.png"
		else return object.item + ".png"
	}
	removeChar(characterId) {
		let char = this.chars.get(characterId)
		if (!char) {
			console.warn("Tried to remove non-existing character", characterId)
			return
		}
		char.destroy() // Destroy the PIXI instance to remove it from the canvas
		this.chars.delete(characterId)
	}
	setLevel(data) {
		this.level = new Level(data)
	}
	listen(elem, func) {
		elem.addEventListener('keydown', func)
	}
	deltaTick(tick) {
		if (tick != null) {
			tick.forEach(change => {
				switch (change.event) {
					case MOVE: // Move
						let char = this.level.objects.find(ob => ob.id == change.id)
						char.pos.x += change.pos.x
						char.pos.y += change.pos.y
						this.updateChar(char)
						break
					case DEATH:
						this.level.objects = this.level.objects.filter(d => d.id != change.id)
						this.removeChar(change.id)
						break
					case SPAWN:
						let newChar = { pos: change.pos, id: change.id, item: change.item, kind: change.kind }
						this.level.objects.push(newChar)
						this.setChar(newChar)
						break
				}
			})
		} else {
			console.warn("Tick was null!")
		}
	}
}