import {config} from "./config/config.js";
let serverUrl = config.serverUrl

function sendRequest(content, url = '', method) {
	console.log("request:", JSON.stringify(content), method)
	// console.log(localStorage.getItem("accessToken"));
	let requestInit = {
		method: method,
		headers: {
			'Content-Type': 'application/json',
			"Authorization": 'Bearer ' + localStorage.getItem("accessToken")
		},
		body: JSON.stringify(content),
	}
	if (method === 'get') {
		requestInit.body = undefined
	}

	return fetch(url, requestInit).then(response => {return response.json()})
}

let loginContent = {
	email: 'meowts@gmail.com',
	username: 'not null',
	password: 'qwerty123'
}

document.querySelector('#sign_up').addEventListener('click', signUp)
document.querySelector('#log_in').addEventListener('click', logIn)
document.querySelector('#check').addEventListener('click', check)

function signUp() {
	sendRequest(loginContent, `${serverUrl}/auth/sign-up`, 'post')
		.then(res => { console.log("response:", res); localStorage.setItem('accessToken', res.accessToken.text()) })
}

function logIn() {
	sendRequest(loginContent, `${serverUrl}/auth/log-in`, 'post')
		.then(res => { console.log("response:", res); localStorage.setItem('accessToken', res.accessToken.text()) })
}
function check() {
	sendRequest(null, `${serverUrl}/auth/check`, 'get')
		.then(res => { console.log("response:", res); renderDiv(res) })
}
