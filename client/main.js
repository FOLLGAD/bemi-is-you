import Game from './game';

window.onload = main

function main() {
	let ws = new WebSocket('ws://localhost:8081/websocket')
	ws.addEventListener('error', console.log)

	let game = new Game()

	function wsHandler(e) {
		let s = e.data.slice(1, -2)
		let data = JSON.parse(atob(s))

		switch (data.msgType) {
			case 0: // Level info
				game.setLevel(data.data)
				game.start()
				break
			case 1: // Delta info
				game.deltaTick(data.data)
			case 2:
				game.playerNum = data.data
		}
	}
	ws.addEventListener('message', wsHandler)

	game.listen(window, e => {
		// if (e.repeat) return;

		let data;

		switch (e.code) {
			case "KeyW":
				data = "up"
				break
			case "KeyA":
				data = "left"
				break
			case "KeyS":
				data = "down"
				break
			case "KeyD":
				data = "right"
				break
			case "KeyZ":
				data = "undo"
				break
			case "KeyR":
				data = "restart"
				break
			default:
				return
		}

		ws.send(JSON.stringify({
			msgType: 0,
			data: data,
		}))
	}) // Initiate listeners on the 'window' object
}