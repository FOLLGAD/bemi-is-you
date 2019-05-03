import * as PIXI from 'pixi.js';

export default class Game {
	// app = null;

	constructor() {
		this.app = new PIXI.Application({
			width: 1024, height: 800, resolution: window.devicePixelRatio || 1, forceCanvas: true, backgroundColor: 0xeeeeee,
		})

		document.body.appendChild(this.app.view)
	}
	start() {
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

		app.stage.addChild(container)

		const char = new PIXI.Graphics()
		char.beginFill(0x0011ee)
		char.drawRect(0, 0, 200, 200);
		// char.anchor.set(0.5)
		char.x = 20
		char.y = 20
		container.addChild(char)

		// Move container to the center
		container.x = app.screen.width / 2
		container.y = app.screen.height / 2

		// Center char sprite in local container coordinates
		container.pivot.x = container.width / 2
		container.pivot.y = container.height / 2
	}
	listen(elem) {
		elem.addEventListener('keydown', this.handleKeyDown)
	}
	handleKeyDown(event) {
		let { code, repeat } = event

		if (!repeat) {
			console.log(code)
		}
	}
}