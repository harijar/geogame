import {ProfileRequest} from "../api.js"

window.onload = async function() {
    document.getElementById("mainMenuButton").onclick = async function() {
        window.location.href = '/index.html'
    }
    document.getElementById("settingsButton").onclick = async function() {
        window.location.href = '/profile/settings'
    }
    await loadProfile()
}

async function loadProfile() {
    const data = await ProfileRequest()
    document.getElementById("title").textContent = data['nickname']
    document.getElementById("greetings").textContent = `Greetings, ${data['nickname']}!`
    document.getElementById("totalGames").textContent = data['total_games']
    document.getElementById("gamesWon").textContent = data['games_won']
    document.getElementById("averageGuesses").textContent = data['average_guesses']
}