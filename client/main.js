import Game from "./game";

const serverUrl = window.location.host
const wsProtocol = window.location.protocol === "https:"
  ? "wss"
  : "ws"

document.addEventListener("DOMContentLoaded", main) // Call main func on dom ready

function main() {
	addDomListeners()

	if (window.location.hash) {
		let d = window.location.hash.match(/\d*$/)[0] // Get room id from hash
		let room = Number(d)
		if (d && !isNaN(room)) {
			joinGame(room)
			return
		} else {
			window.location.hash = ""
		}
	}

	fetchSessions()
}

function addDomListeners() {
	document.querySelector("#session-list").addEventListener("click", e => {
		let elem = e.target

		if (elem.className && elem.className.indexOf("join-game") != -1) {
			let room = elem.dataset.room
			joinGame(room)
		}
	})

	document.querySelector("#session-create").addEventListener("click", () => {
		createGame()
	})
}

function fetchSessions() {
	document.querySelector("#lobby").className = ""
	document.querySelector("#canv").className = "hide"

	fetch(`/sessions`)
		.then(res => res.json())
		.then(list => {
			let html = list.map(entry => {
				return `<li>Room ${entry.room} (${entry.players} playing) <a class="join-game" data-room="${entry.room}">Join</a></li>`
			})
			document.querySelector("#session-list").innerHTML = html
		})
		.catch(() => {
			document.querySelector("#session-list").innerHTML = `<p>Couldn't connect to server</p>`
		})

}

function createGame() {
	console.log("Creating")
	let ws = new WebSocket(`${wsProtocol}://${serverUrl}/create`)
	instantiateGame(ws)
}

function joinGame(roomNum) {
	console.log("Joining", roomNum)
	let ws = new WebSocket(`${wsProtocol}://${serverUrl}/join?room=${roomNum}`)
	instantiateGame(ws)
}

function setRoomHash(room) {
	window.location.hash = "#/room/" + room
}

function instantiateGame(ws) {
	document.querySelector("#canv").innerHTML = ""
	document.querySelector("#lobby").className = "hide"
	document.querySelector("#canv").className = ""
	ws.addEventListener("error", () => {
		window.location.hash = ""
		fetchSessions()
	})
	let game = new Game()

	function wsHandler(e) {
		let data = JSON.parse(e.data)

		switch (data.msgType) {
			case 0: // Level info
				if (!game.running) {
					game.start()
				}
				game.setLevel(data.data)
				if (data.room) {
					setRoomHash(data.room)
				}
				break
			case 1: // Delta info
				game.deltaTick(data.data)
				break
			case 2:
				game.playerNum = data.data
				break
		}
	}
	ws.addEventListener("message", wsHandler)

	// Initiate listeners on the "window" object
	game.listen(window, e => {
		// if (e.repeat) return;

		let data;

		switch (e.code) {
			case "KeyW":
			case "ArrowUp":
				data = "up"
				break
			case "KeyA":
			case "ArrowLeft":
				data = "left"
				break
			case "KeyS":
			case "ArrowDown":
				data = "down"
				break
			case "KeyD":
			case "ArrowRight":
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
	})
}
