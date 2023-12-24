import {Auth, CheckAuth} from "./auth.js"

window.onload = async function() {
    await CheckAuth()
    document.getElementById("profileButton").onclick = async function() {
        window.location.href = '/profile'
    }
    document.getElementById("playButton").onclick = async function() {
        window.location.href = 'game.html'
    }
}
async function auth(user) {
    await Auth(user)
}